// Just a collection of common functions that are used in multiple places in the store package
package store

import (
	"errors"

	"gorm.io/gorm"
)

// Generic Function that will take in a struct that contains a gorm.model, and a search term of any type, and return the first item that matches the search term
func getDBItemByIdentifier[Model any, Property int | string | uint](db *gorm.DB, model *Model, searchTerm Property) (*Model, error) {
	err := db.Where("id = ?", searchTerm).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return model, nil
}
