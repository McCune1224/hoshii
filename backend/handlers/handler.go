package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/hoshii/store"
	"gorm.io/gorm"
)

func notImplemented(c *fiber.Ctx) error {
	message := fmt.Sprintf("%s Not implemented yet", c.Path())
	return c.SendString(message)
}

// Universal handler struct (ideally handled by dependency injection besides the passed db)
type Handler struct {
	// Really just here incase I need to do a manual query, ideally stores should handle all db ops though...
	db            *gorm.DB
	UserStore     store.UserStore
	WishlistStore store.WishlistStore
}

func NewHandler(db *gorm.DB, us store.UserStore, ws store.WishlistStore) *Handler {
	return &Handler{
		db:            db,
		UserStore:     us,
		WishlistStore: ws,
	}
}

// Tie all routes from fiber to this handler
func (h *Handler) AddRoutes(app *fiber.App) {
	api := app.Group("/api")
	// User routes
	users := api.Group("/users")
	users.Get("/", notImplemented)

	// Wishlist routes
	wishlists := api.Group("/wishlists")
	wishlists.Get("/", notImplemented)
}
