package handler

import (
	"listr/store"

	"github.com/go-playground/validator/v10"
)

// Wrapper struct for validator for custom functionality ontop of base validator
type Validator struct {
	validator *validator.Validate
}

func NewStructValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Takes in a struct and validates it using go-playground validator library
func (v Validator) Validate(s any) error {
	return v.validator.Struct(s)
}

type Handler struct {
	UserStore       store.UserOps
	WishlistStore   store.WishlistOps
	StructValidator *Validator
}

func NewHandler(userStore store.UserOps, WishlistStore store.WishlistOps) *Handler {
	return &Handler{
		UserStore:       userStore,
		WishlistStore:   WishlistStore,
		StructValidator: NewStructValidator(),
	}
}
