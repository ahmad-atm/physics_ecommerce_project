package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func allCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var cart []Cart

	db.Find(&cart, "user_id == ? AND status == ?", currentUser.ID, "Pending")
	json.NewEncoder(w).Encode(cart)
}

func addCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var cart Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		log.Fatalln(err)
		return
	}
	db.Create(&Cart{
		ID:             uuid.NewString(),
		UserId:         currentUser.ID,
		ProductID:      cart.ProductID,
		ProductName:    cart.ProductName,
		ProductImage:   cart.ProductImage,
		ProductPrice:   cart.ProductPrice,
		Quantity:       cart.Quantity,
		AdditionalInfo: cart.AdditionalInfo,
		Status:         "Pending",
	})

	json.NewEncoder(w).Encode(cart)
}

func getCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var cart Cart
	id := strings.TrimPrefix(r.URL.Path, "/get_cart/")
	db.First(&cart, "id == ?", id)
	json.NewEncoder(w).Encode(cart)
}

func confirmCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var cart Cart
	var updatedCart Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(cart)
	db.First(&updatedCart, "id == ?", cart.ID)
	fmt.Println(updatedCart)
	db.Model(&updatedCart).Select("Quantity", "AdditionalInfo", "Status", "TotalPrice").Updates(&Cart{
		Quantity:       cart.Quantity,
		AdditionalInfo: cart.AdditionalInfo,
		Status:         cart.Status,
		TotalPrice:     cart.TotalPrice,
	})
	json.NewEncoder(w).Encode(updatedCart)
}

func deleteCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var cart Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(cart.ID, "cart")
	db.First(&cart, "id == ?", cart.ID)
	fmt.Println("1", cart.UserId, "1", "2", currentUser.ID, "2")
	if cart.UserId != currentUser.ID {
		http.Error(w, "Unauthorized Operation", http.StatusUnauthorized)
		return
	}
	db.Where("id == ?", cart.ID).Unscoped().Delete(&Cart{})
	json.NewEncoder(w).Encode(cart)
}

func allConfirmedCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var cart []Cart

	db.Find(&cart, "status == ?", "Order_Confirmed")
	json.NewEncoder(w).Encode(cart)
}
