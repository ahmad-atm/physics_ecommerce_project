package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

func allProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product []Product
	db.Find(&product)
	json.NewEncoder(w).Encode(product)
}

func count(w http.ResponseWriter, r *http.Request) {
	type Count struct {
		UserCount    int64 `json:"userCount"`
		ProductCount int64 `json:"productCount"`
		Order        int64 `json:"orderCount"`
	}
	var count Count

	db.Find(&User{}).Count(&count.UserCount)
	db.Find(&Product{}).Count(&count.ProductCount)
	db.Find(&Cart{}, "status == ?", "Order_Confirmed").Count(&count.Order)

	json.NewEncoder(w).Encode(count)
}

func latestProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product []Product
	db.Find(&product).Limit(3)
	json.NewEncoder(w).Encode(product)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/get_product/")
	var product Product
	db.Find(&product, "id == ?", id)
	json.NewEncoder(w).Encode(product)
}
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "multipart/form-data")
	r.ParseMultipartForm(10 * 1024 * 1024)

	var product Product
	var user User

	file, _, err := r.FormFile("image1")
	product.Name = r.FormValue("name")
	product.Description = r.FormValue("description")
	product.Category = r.FormValue("category")
	product.Price = r.FormValue("price")

	db.First(&user, "id = ?", currentUser.ID)
	if currentUser.Role != "Admin" {
		http.Error(w, "Forbidden Access", http.StatusForbidden)
		return
	}

	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	temp, tempErr := os.CreateTemp("static", "image-*.jpg")

	if tempErr != nil {
		log.Fatalln(tempErr)
		return
	}
	defer temp.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
		return
	}
	temp.Write(fileBytes)
	fmt.Println("Done")
	db.Create(&Product{
		ID:          uuid.NewString(),
		Name:        product.Name,
		Image1:      temp.Name(),
		Description: product.Description,
		Price:       product.Price,
		UserId:      currentUser.ID,
		Category:    product.Category,
	})
	json.NewEncoder(w).Encode(&product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "multipart/form-data")
	r.ParseMultipartForm(10 * 1024 * 1024)

	var product Product
	var updatedProduct Product
	var user User

	product.Name = r.FormValue("name")
	product.Description = r.FormValue("description")
	product.Category = r.FormValue("category")
	product.Price = r.FormValue("price")
	product.ID = r.FormValue("productId")

	db.First(&user, "id == ?", currentUser.ID)
	db.First(&updatedProduct, "id == ?", product.ID)
	if currentUser.Role != "Admin" {
		http.Error(w, "Forbidden Access", http.StatusForbidden)
		return
	}

	db.Model(&updatedProduct).Select("Name", "Description", "Category", "Price").Updates(Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
	})
	json.NewEncoder(w).Encode(product)
}

func updateProductImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "multipart/form-data")
	r.ParseMultipartForm(10 * 1024 * 1024)

	var product Product
	var updatedProduct Product
	var user User

	file, _, err := r.FormFile("image1")
	product.ID = r.FormValue("productId")

	db.First(&user, "id == ?", currentUser.ID)
	db.First(&updatedProduct, "id == ?", product.ID)
	if currentUser.Role != "Admin" {
		http.Error(w, "Forbidden Access", http.StatusForbidden)
		return
	}

	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	temp, tempErr := os.CreateTemp("static", "image-*.jpg")

	if tempErr != nil {
		log.Fatalln(tempErr)
		return
	}
	defer temp.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
		return
	}
	temp.Write(fileBytes)
	fmt.Println("Done")

	db.Model(&updatedProduct).Update("image1", temp.Name())
	json.NewEncoder(w).Encode(product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusNotImplemented)
		return
	}
	if product.UserId == currentUser.ID {
		db.Where("id == ?", product.ID).Unscoped().Delete(Product{})
	}
	json.NewEncoder(w).Encode(product)
}
