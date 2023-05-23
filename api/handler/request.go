package handler

import (
	"github.com/mccune1224/listr/models"
)

type UserLoginRequest struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

type UserSignUpRequest struct {
	User struct {
		Username string `json:"username" validate:"required,min=3,max=20"`
		Email    string `json:"email" validate:"required,email,max=50"`
		Password string `json:"password" validate:"required,min=6,max=32"`
	} `json:"user"`
}

func (r *UserSignUpRequest) bindToModel(u *models.User, v *Validator) error {
	if err := v.Validate(r); err != nil {
		return err
	}

	hashedPassword, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Username = r.User.Username
	u.Email = r.User.Email
	u.Password = hashedPassword
	return nil
}

type UserResponse struct {
	User struct {
		Username string  `json:"username"`
		Email    string  `json:"email"`
		Bio      *string `json:"bio"`
	} `json:"user"`
}

func NewUserResponse(dbUser *models.User) *UserResponse {
	resp := UserResponse{}
	resp.User.Username = dbUser.Username
	resp.User.Email = dbUser.Email
	resp.User.Bio = dbUser.Bio
	return &resp
}
