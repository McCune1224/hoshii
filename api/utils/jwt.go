package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JWTSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}
