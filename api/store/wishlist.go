package store

import (
	"github.com/mccune1224/listr/models"

	"gorm.io/gorm"
)

// All expected operations for Wishlists database operations
type WishlistOps interface {
	CreateWishlist(wishlist *models.Wishlist) error
	GetWishlistById(id uint) (*models.Wishlist, error)
	GetWishlistsByUserId(id uint) ([]*models.Wishlist, error)
	UpdateWishlist(wishlist *models.Wishlist) error
	DeleteWishlist(id uint) error
}

// Handles Database operations for Wishlists
type WishlistStore struct {
	db *gorm.DB
}

func NewWishlistStore(db *gorm.DB) *WishlistStore {
	return &WishlistStore{
		db: db,
	}
}

func (us *UserStore) CreateWishlist(wishlist *models.Wishlist) error {
	return us.db.Create(wishlist).Error
}

func (us *UserStore) GetWishlistById(id uint) (*models.Wishlist, error) {
	wishlist := &models.Wishlist{}
	return getDBItemByIdentifier(us.db, wishlist, id)
}

func (us *UserStore) GetWishlistsByUserId(id uint) ([]*models.Wishlist, error) {
	wishlists := []*models.Wishlist{}
	err := us.db.Where("user_id = ?", id).Find(&wishlists).Error
	if err != nil {
		return nil, err
	}
	return wishlists, nil
}

func (us *UserStore) UpdateWishlist(wishlist *models.Wishlist) error {
	return us.db.Save(wishlist).Error
}

func (us *UserStore) DeleteWishlist(id uint) error {
	return us.db.Delete(&models.Wishlist{}, id).Error
}
