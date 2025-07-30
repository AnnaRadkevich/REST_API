package app

import (
	"github.com/gofiber/fiber/v2"
	"main/routes"
	"main/utils"
)

func NewFiberApp() *fiber.App {
	utils.LoadEnv()
	var app *fiber.App = fiber.New()
	routes.SetupRoutes(app)
	return app
}
