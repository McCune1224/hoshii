package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h *Handler) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	dbUser, err := h.UserStore.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	resp := dbUser.ToResponse()

	return c.JSON(fiber.Map{
		"data": resp,
	})
}

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "GET USERS HIT",
	})
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	newUserReq := &struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := c.BodyParser(newUserReq); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to parse request body\n%s", err.Error()),
		})
	}

	return c.JSON(fiber.Map{
		"message": newUserReq,
	})
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "UPDATE USER HIT",
	})
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DELETE USER HIT",
	})
}
