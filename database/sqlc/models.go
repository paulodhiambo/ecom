// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package database

import (
	"time"
)

type Category struct {
	ID      int64  `json:"id"`
	CatName string `json:"cat_name"`
}

type Country struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	ContinentName string `json:"continent_name"`
}

type Merchant struct {
	ID           int64     `json:"id"`
	MerchantName string    `json:"merchant_name"`
	CountryCode  string    `json:"country_code"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Order struct {
	ID        int64     `json:"id"`
	UserID    int32     `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItem struct {
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type Product struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	MerchantID int32     `json:"merchant_id"`
	Price      int32     `json:"price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CategoryID int32     `json:"category_id"`
}

type User struct {
	ID          int64     `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"date_of_birth"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CountryCode string    `json:"country_code"`
}
