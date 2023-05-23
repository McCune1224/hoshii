package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserCustomClaims struct {
	// The standard claims.
	jwt.StandardClaims
	Email  string `json:"email"`
	UserID uint   `json:"userID"`
}

func (c *UserCustomClaims) Validate(ctx context.Context) error {
	if c.UserID == 0 {
		return errors.New("userID is required")
	}
	return nil
}

// We want this struct to be filled in with
// our custom claims from the token.
var customClaims = func() validator.CustomClaims {
	return &UserCustomClaims{}
}

// ValidateJWT is a gin.HandlerFunc middleware
// that will check the validity of our JWT.
func ValidateJWT() gin.HandlerFunc {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}
	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
	// Set up the validator.
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.HS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(customClaims),
		validator.WithAllowedClockSkew(30*time.Second),
	)
	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	// Actual Handler function to run after JWT validation
	return func(ctx *gin.Context) {
		encounteredError := true
		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			encounteredError = false
			ctx.Request = r
			ctx.Next()
		}

		middleware.CheckJWT(handler).ServeHTTP(ctx.Writer, ctx.Request)

		if encounteredError {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"message": "JWT is invalid."},
			)
		}
	}
}
