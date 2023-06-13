package handlers

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h *Handler) RegisterUser(c *fiber.Ctx) error {
	userRegisterReq := struct {
		Username string `json:"username" validate:"required,min=3,max=32"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=32"`
	}{}
	if err := c.BodyParser(&userRegisterReq); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to parse request body\n%s", err.Error()),
		})
	}
	if err := h.Validator.Struct(userRegisterReq); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return c.Status(500).JSON(fiber.Map{
				"error": fmt.Sprintf("failed to validate request body\n%s", err.Error()),
			})
		}
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("invalid request body\n%s", err.Error()),
		})
	}
	// Check if username is taken
	existingUsername, err := h.UserStore.GetByUsername(userRegisterReq.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to query database\n%s", err.Error()),
		})
	}
	if existingUsername != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("username %s is already taken", userRegisterReq.Username),
		})
	}
	// Check if email is taken
	existingEmail, err := h.UserStore.GetByEmail(userRegisterReq.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to query database\n%s", err.Error()),
		})
	}
	if existingEmail != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("email %s is already taken", userRegisterReq.Email),
		})
	}

	// Hash password

	// Create user in db

	return c.JSON(fiber.Map{
		"message": userRegisterReq,
	})
}

func (h *Handler) LoginUser(c *fiber.Ctx) error {
	userLoginReq := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := c.BodyParser(&userLoginReq); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to parse request body\n%s", err.Error()),
		})
	}
	return c.JSON(fiber.Map{
		"message": userLoginReq,
	})
}
