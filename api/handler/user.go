package handler

import "github.com/gin-gonic/gin"

type UserRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`

}

func (h *Handler) SignUp(c *gin.Context) {
}
