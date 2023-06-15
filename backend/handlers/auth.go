package handlers

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mccune1224/hoshii/types"
	"golang.org/x/crypto/bcrypt"
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

	// Validate body struct
	if err := h.Validator.Struct(userRegisterReq); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to validate request body\n%s", err.Error()),
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

	// bcrypt hash password

	hash, err := bcrypt.GenerateFromPassword([]byte(userRegisterReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to hash password\n%s", err.Error()),
		})
	}
	// Create user in db
	newUser := types.User{
		Username: userRegisterReq.Username,
		Email:    userRegisterReq.Email,
		Password: string(hash),
	}
	if err := h.UserStore.Create(&newUser); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to create user\n%s", err.Error()),
		})
	}

	return c.JSON(fiber.Map{
		"message": userRegisterReq,
	})
}

func (h *Handler) LoginUser(c *fiber.Ctx) error {
	userLoginReq := struct {
		Username string `json:"username" validate:"required,min=3,max=32"`
		Password string `json:"password" validate:"required,min=8,max=32"`
	}{}

	if err := c.BodyParser(&userLoginReq); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to parse request body\n%s", err.Error()),
		})
	}

	if err := h.Validator.Struct(userLoginReq); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to validate request body\n%s", err.Error()),
		})
	}

	dbUser, err := h.UserStore.GetByUsername(userLoginReq.Username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to query database\n%s", err.Error()),
		})
	}

	if dbUser == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("username %s does not exist", userLoginReq.Username),
		})
	}

	compareErr := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(userLoginReq.Password))

	if compareErr != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("incorrect password"),
		})
	}

	// Make claims
	claims := jwt.MapClaims{
		"username": dbUser.Username,
		"id":       dbUser.ID,
		// Just doing a week for now, can modify later
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to sign token\n%s", err.Error()),
		})
	}

	return c.JSON(fiber.Map{
		"token": signedToken,
	})
}
