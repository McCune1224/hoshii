package models

import "gorm.io/gorm"

type Wishlist struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"isPrivate"`
	// Wishlists belong to user models with a foreign key of UserId
	UserId int  `json:"userId"`
	User   User `json:"user"`
}

type Item struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Link        string  `json:"link"`
	Price       float64 `json:"price"`
	// Items belong to wishlists with a foreign key of WishlistId
	Wishlist   Wishlist `json:"wishlist"`
	WishlistId int      `json:"wishlistId"`
}

type Favorite struct {
    gorm.Model
    User User `json:"user"`
    UserId int `json:"userId"`
    Wishlist Wishlist `json:"wishlist"`
    WishlistId int `json:"wishlistId"`
}
