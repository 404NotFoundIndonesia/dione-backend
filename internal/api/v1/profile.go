package v1

import (
	"context"
	"dione-backend/domain"
	"dione-backend/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type profileApi struct {
	userService domain.UserService
}

func NewProfileApi(app *fiber.App, userService domain.UserService, middleware fiber.Handler) {
	api := &profileApi{userService: userService}

	app.Get("/profile", middleware, api.Show)
}

func (api profileApi) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	user := ctx.Locals("user")

	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	id, ok := claims["id"].(string)
	if !ok {
		return ctx.Status(http.StatusForbidden).JSON(dto.CreateErrorResponse("invalid token"))
	}

	res, err := api.userService.Show(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateErrorResponse(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateSuccessResponse(res))
}
