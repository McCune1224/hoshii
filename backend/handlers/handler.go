package handlers

import (
	"fmt"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/hoshii/store"
	"gorm.io/gorm"
)

func notImplemented(c *fiber.Ctx) error {
	message := fmt.Sprintf("%s Not implemented yet", c.Path())
	return c.JSON(fiber.Map{
		"message": message,
	})
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

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	}))

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/login", h.LoginUser)
	auth.Post("/register", h.RegisterUser)

	// Me routes
	me := api.Group("/me")
	me.Get("/users", notImplemented)
	me.Get("/wishlists", notImplemented)

	// User routes
	users := api.Group("/users")
	users.Get("/:id", h.GetUser)
	users.Get("/", h.GetUsers)
	users.Post("/", h.CreateUser)
	users.Put("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)
	// Wishlist routes
	wishlists := api.Group("/wishlists")
	wishlists.Get("/", notImplemented)
}
