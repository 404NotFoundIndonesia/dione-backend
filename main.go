package main

import (
	api "dione-backend/internal/api/v1"
	"dione-backend/internal/config"
	"dione-backend/internal/connection"
	"dione-backend/internal/repository"
	"dione-backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := config.Get()
	db := connection.GetDatabase(conf.Database)

	userRepository := repository.NewUserRepository(db)

	authService := service.NewAuthService(conf, userRepository)
	userService := service.NewUserService(userRepository)

	app := fiber.New()

	authMiddleware := api.AuthMiddleware(conf)

	api.NewAuthApi(app, authService)
	api.NewProfileApi(app, userService, authMiddleware)

	_ = app.Listen(conf.Server.Host + ":" + conf.Server.Port)
}
