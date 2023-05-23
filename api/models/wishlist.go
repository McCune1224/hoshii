package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wishlist struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title       string    `gorm:"not null"`
	Description string
	UserId      uint `gorm:"not null"`
	User        User
	Items       []Item
	Favorites   []User `gorm:"many2many:favorites;"`
}

func (wl *Wishlist) BeforeCreate(tx *gorm.DB) (err error) {
	wl.ID = uuid.New()
	return
}

type Item struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title       string    `gorm:"not null"`
	Description string
	Link        string
	Image       string
	Price       float64
	WishlistId  uint `gorm:"not null"`
	Wishlist    Wishlist
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New()
	return
}
