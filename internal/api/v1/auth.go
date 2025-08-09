package v1

import (
	"context"
	"dione-backend/domain"
	"dione-backend/dto"
	"dione-backend/internal/util"
	"github.com/gofiber/fiber/v2"
	"time"
)

type authApi struct {
	authService domain.AuthService
}

func NewAuthApi(app *fiber.App, authService domain.AuthService) {
	api := &authApi{authService: authService}

	app.Post("/auth/login", api.Login)
	app.Post("/auth/register", api.Register)
}

func (api authApi) Login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(dto.CreateErrorResponseWithPayload("validation error", fails))
	}

	res, err := api.authService.Login(c, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(dto.CreateErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccessResponse(res))
}

func (api authApi) Register(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(dto.CreateErrorResponseWithPayload("validation error", fails))
	}

	res, err := api.authService.Register(c, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(dto.CreateErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccessResponse(res))
}
