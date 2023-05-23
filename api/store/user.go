package store

import (
	"github.com/mccune1224/listr/models"

	"gorm.io/gorm"
)

type UserOps interface {
	CreateUser(user *models.User) error
	GetUserById(id uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	// Not sure if neccesary so leaving out for now
	// GetUsers() ([]*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

// Handles Database operations for Users
type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) CreateUser(user *models.User) error {
	return us.db.Create(user).Error
}

func (us *UserStore) GetUserById(id uint) (*models.User, error) {
	user := &models.User{}
	return getDBItemByIdentifier(us.db, user, id)
}

func (us *UserStore) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	return getDBItemByIdentifier(us.db, user, username)
}

func (us *UserStore) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	return getDBItemByIdentifier(us.db, user, email)
}

// func (us *UserStore) GetUsers() ([]*models.User, error)  {}

func (us *UserStore) UpdateUser(user *models.User) error {
	return us.db.Save(user).Error
}

func (us *UserStore) DeleteUser(id uint) error {
	return us.db.Delete(&models.User{}, id).Error
}
