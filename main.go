package main

import (
	"dione-backend/internal/config"
	"dione-backend/internal/connection"
	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := config.Get()
	_ = connection.GetDatabase(conf.Database)

	app := fiber.New()

	_ = app.Listen(conf.Server.Host + ":" + conf.Server.Port)
}
