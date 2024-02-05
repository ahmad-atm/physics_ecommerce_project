package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	ProfileImage string `json:"profileImage"`
	Password     string `json:"password"`
	Role         string `json:"role"`
}

type Product struct {
	gorm.Model
	ID     string `gorm:"primaryKey"`
	Name   string `json:"name"`
	Image1 string `json:"image1"`
	// Image2      string `json:"image2"`
	// Image3      string `json:"image3"`
	Description string `json:"description"`
	Price       string `json:"price"`
	UserId      string `json:"userId"`
	Category    string `json:"category"`
	// Inventory   ProductInventory
}

// type Cart struct {
// 	gorm.Model
// 	ID        string `gorm:"primaryKey"`
// 	ProductID string `json:"productID"`
// 	Quantity  string `json:"quantity"`
// }

type LoggedUser struct {
	ID    string
	Name  string
	Email string
	Role  string
}

type UserAddress struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	Contact      string `json:"contact"`
}

type Cart struct {
	gorm.Model
	ID             string `gorm:"primaryKey"`
	UserId         string `json:"userId"`
	ProductID      string `json:"productID"`
	ProductName    string `json:"productName"`
	ProductImage   string `json:"productImage"`
	ProductPrice   string `json:"productPrice"`
	Quantity       string `json:"quantity"`
	AdditionalInfo string `json:"additionalInfo"`
	TotalPrice     string `json:"totalPrice"`
	Status         string `json:"status"`
}

// type Order struct {
// 	gorm.Model
// 	ID string `gorm:"primaryKey"`
// 	CartID string `gorm:"cartId"`
// 	Cart Cart
// }
