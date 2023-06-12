package store

import (
	"github.com/mccune1224/hoshii/models"
	"gorm.io/gorm"
)

// Interface for how DB operations should be handled for a wishlist 
type WishlistStore interface {
	Get(id int) (*models.Wishlist, error)
	GetMany(offset int, limit int) ([]*models.Wishlist, error)
	Create(w *models.Wishlist) error
	Update(w *models.Wishlist) error
	Delete(id int) error
}

type PostgreWishlistStore struct {
	db *gorm.DB
}

func NewPostgreWishlistStore(db *gorm.DB) *PostgreWishlistStore {
	return &PostgreWishlistStore{db}
}

func (s *PostgreWishlistStore) Get(id int) (*models.Wishlist, error) {
    var wishlist models.Wishlist
    if err := s.db.First(&wishlist, id).Error; err != nil {
        return nil, err
    }
    return &wishlist, nil
}

func (s *PostgreWishlistStore) GetMany(offset int, limit int) ([]*models.Wishlist, error) {
    var wishlists []*models.Wishlist
    if err := s.db.Offset(offset).Limit(limit).Find(&wishlists).Error; err != nil {
        return nil, err
    }
    return wishlists, nil
}

func (s *PostgreWishlistStore) Create(w *models.Wishlist) error {
    if err := s.db.Create(&w).Error; err != nil {
        return err
    }
    return nil
}

func (s *PostgreWishlistStore) Update(w *models.Wishlist) error {
    if err := s.db.Save(&w).Error; err != nil {
        return err
    }
    return nil
}

func (s *PostgreWishlistStore) Delete(id int) error {
    if err := s.db.Delete(&models.Wishlist{}, id).Error; err != nil {
        return err
    }
    return nil
}
