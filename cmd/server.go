package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/JoelTinx/ws-matrix/api"
)

func main() {
	app := fiber.New()
	api.SetupRoutes(app)
	app.Listen(":3000")
}