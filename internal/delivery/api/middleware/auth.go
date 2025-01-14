package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (middleware *AuthMiddleware) IsUser(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeaders := c.GetReqHeaders()["Authorization"]
		if len(authHeaders) == 0 {
			next(c)
			return nil
		}

		authHeader := authHeaders[0]
		authToken := authHeader[len("Bearer "):]

		token, err := jwt.Parse(authToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if !token.Valid || err != nil {
			c.SendStatus(http.StatusUnauthorized)
			return nil
		}

		var email string
		token, _, err = new(jwt.Parser).ParseUnverified(authToken, jwt.MapClaims{})
		if err != nil {
			c.SendStatus(http.StatusUnauthorized)
			return nil
		} else if claims, ok := token.Claims.(jwt.MapClaims); ok {
			email = fmt.Sprint(claims["email"])
		}

		c.Locals("auth_user_email", email)

		next(c)
		return nil
	}
}
