package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/JoelTinx/ws-matrix/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.GetWelcome)
	app.Post("/rotate", controllers.PostRotateMatrix)
}