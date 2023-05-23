package handler

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/listr/models"
)

func (h *Handler) UserSignup(c *fiber.Ctx) error {
	incomingSignupData := UserSignUpRequest{}
	err := c.BodyParser(&incomingSignupData.User)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	log.Println(incomingSignupData)
	newUserModel := models.User{}
	err = incomingSignupData.bindToModel(&newUserModel, h.StructValidator)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err = h.UserStore.CreateUser(&newUserModel)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(NewUserResponse(&newUserModel))
}

func (h *Handler) UserLogin(c *fiber.Ctx) error {
	foo := struct {
		Username string `json:"username"`
	}{}
	c.BodyParser(&foo)
	return c.Status(http.StatusNotImplemented).JSON(foo)
}
