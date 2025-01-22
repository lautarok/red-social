package middleware

import (
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
		var authToken string

		authHeaders := c.GetReqHeaders()["Authorization"]
		if len(authHeaders) > 0 {
			authHeader := authHeaders[0]
			authToken = authHeader[len("Bearer "):]
		} else {
			authQuery := c.Query("auth_token")
			if len(authQuery) > 0 {
				authToken = authQuery
			} else {
				next(c)
				return nil
			}
		}

		token, err := jwt.Parse(authToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if !token.Valid || err != nil {
			c.SendStatus(fiber.StatusUnauthorized)
			return nil
		}

		var id int64
		var email string
		token, _, err = new(jwt.Parser).ParseUnverified(authToken, jwt.MapClaims{})
		if err != nil {
			c.SendStatus(fiber.StatusUnauthorized)
			return nil
		} else if claims, ok := token.Claims.(jwt.MapClaims); ok {
			id = int64(claims["id"].(float64))
			email = claims["email"].(string)
		}

		c.Locals("auth_user_id", id)
		c.Locals("auth_user_email", email)

		next(c)
		return nil
	}
}
