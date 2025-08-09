package v1

import (
	"dione-backend/dto"
	"dione-backend/internal/config"
	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(conf *config.Config) fiber.Handler {
	return jwtMiddleware.New(jwtMiddleware.Config{
		SigningKey: jwtMiddleware.SigningKey{Key: []byte(conf.Jwt.Key)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).
				JSON(dto.CreateErrorResponse("invalid token"))
		},
	})
}

func RequireRoles(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user")
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.CreateErrorResponse("unauthorized"))
		}

		token := user.(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		role, ok := claims["role"].(string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(dto.CreateErrorResponse("invalid role"))
		}

		for _, allowed := range allowedRoles {
			if role == allowed {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(dto.CreateErrorResponse("access denied"))
	}
}

func GetUserRole(c *fiber.Ctx) string {
	user := c.Locals("user")

	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	return claims["role"].(string)
}

func GetUserID(c *fiber.Ctx) string {
	user := c.Locals("user")

	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	return claims["id"].(string)
}
