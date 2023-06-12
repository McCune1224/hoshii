package store

import (
	"github.com/mccune1224/hoshii/types"
	"gorm.io/gorm"
)

// Interface for how DB operations should be handled for a wishlist
type UserStore interface {
	Get(id int) (*types.User, error)
	GetMany(offset int, limit int) ([]*types.User, error)
	Create(u *types.User) error
	Update(u *types.User) error
	Delete(id int) error
}

type PostgreUserStore struct {
	db *gorm.DB
}

func NewPostgreUserStore(db *gorm.DB) *PostgreUserStore {
	return &PostgreUserStore{db}
}

func (s *PostgreUserStore) Get(id int) (*types.User, error) {
	var user types.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *PostgreUserStore) GetMany(offset int, limit int) ([]*types.User, error) {
	var users []*types.User
	if err := s.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *PostgreUserStore) Create(u *types.User) error {
	if err := s.db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (s *PostgreUserStore) Update(u *types.User) error {
	if err := s.db.Save(&u).Error; err != nil {
		return err
	}
	return nil
}

func (s *PostgreUserStore) Delete(id int) error {
	if err := s.db.Delete(&types.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
