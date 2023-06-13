package store

import (
	"github.com/mccune1224/hoshii/types"
	"gorm.io/gorm"
)

// Interface for how DB operations should be handled for a wishlist
type UserStore interface {
	GetById(id int) (*types.User, error)
	GetByUsername(username string) (*types.User, error)
	GetByEmail(email string) (*types.User, error)
	GetMany(offset int, limit int) ([]*types.User, error)
	Create(u *types.User) error
	Update(u *types.User) error
	Delete(id int) error
}

type PostgreUserStore struct {
	db *gorm.DB
}

// GetByEmail implements UserStore.
func (s *PostgreUserStore) GetByEmail(email string) (*types.User, error) {
	var user types.User
	if err := s.db.First(&user, email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewPostgreUserStore(db *gorm.DB) *PostgreUserStore {
	return &PostgreUserStore{db}
}

// GetById implements UserStore.
func (s *PostgreUserStore) GetById(id int) (*types.User, error) {
	var user types.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername implements UserStore.
func (s *PostgreUserStore) GetByUsername(username string) (*types.User, error) {
	var user types.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetMany implements UserStore.
func (s *PostgreUserStore) GetMany(offset int, limit int) ([]*types.User, error) {
	var users []*types.User
	if err := s.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Create implements UserStore.
func (s *PostgreUserStore) Create(u *types.User) error {
	if err := s.db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

// Update implements UserStore.
func (s *PostgreUserStore) Update(u *types.User) error {
	if err := s.db.Save(&u).Error; err != nil {
		return err
	}
	return nil
}

// Delete implements UserStore.
func (s *PostgreUserStore) Delete(id int) error {
	if err := s.db.Delete(&types.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
