package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) NotImplemented(c *fiber.Ctx) error {
	return c.Status(http.StatusNotImplemented).JSON(
		fiber.Map{
			"message": fmt.Sprintf("MY Not implemented %s", c.Path()),
		},
	)
}

func (h *Handler) RegisterRoutes(router *fiber.App) {
	v1 := router.Group("/api/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Greetings from API Root"})
	})

	// User routes
	users := v1.Group("/users") // api/v1/users
	{
		users.Post("/signup", h.UserSignup)
		users.Post("/login", h.NotImplemented)

		authedUsers := users.Group("/") // api/v1/users
		{
			authedUsers.Get("/:id", h.NotImplemented)
			authedUsers.Put("/:id", h.NotImplemented)
			authedUsers.Delete("/:id", h.NotImplemented)
		}
	}

	// Wishlist routes
	wishlists := v1.Group("/wishlists") // api/v1/wishlists
	{
		wishlists.Get("/:slug", h.NotImplemented)
		wishlists.Get("/", h.NotImplemented)

		authedWishlists := wishlists.Group("/") // api/v1/wishlists
		{

			authedWishlists.Post("/", h.NotImplemented)
			authedWishlists.Put("/:slug", h.NotImplemented)
			authedWishlists.Delete("/:slug", h.NotImplemented)

			// Favorite/unfavorite routes for wishlists
			authedWishlists.Post("/:slug/favorites", h.NotImplemented)
			authedWishlists.Delete("/:slug/favorites", h.NotImplemented)

		}

	}
}
