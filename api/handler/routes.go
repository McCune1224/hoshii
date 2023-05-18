package handler

import "github.com/gin-gonic/gin"

func (h *Handler) NotImplemented(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Not Implemented",
	})
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello ROOT",
		})
	})
	v1 := router.Group("/api/v1")

	users := v1.Group("/users")
	{
		users.POST("/signup", h.NotImplemented)
		users.POST("/login", h.NotImplemented)

		users.GET("/:id", h.NotImplemented)

		users.PUT("/:id", h.NotImplemented)

		users.DELETE("/:id", h.NotImplemented)
	}

	wishlist := v1.Group("/wishlist")
	{
		wishlist.POST("/", h.NotImplemented)

		wishlist.GET("/:id", h.NotImplemented)
		wishlist.GET("/", h.NotImplemented)
	}
}
