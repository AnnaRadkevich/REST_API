package routes

import (
	"github.com/gofiber/fiber/v2"
	"main/handlers"
	"main/middlewares"
)

func SetupRoutes(app *fiber.App) {
	//public routes
	var publicRoutes fiber.Router = app.Group("/api/v1")
	publicRoutes.Post("/signup", handlers.SighUp)
	publicRoutes.Post("/login", handlers.Login)
	publicRoutes.Get("/items", handlers.GetAllItems)
	publicRoutes.Get("/items/:id", handlers.GetItemByID)

	//private routes, authentication is required
	var privateRoutes fiber.Router = app.Group("/api/v1")
	privateRoutes.Post("/items", middlewares.CreateMiddleWare(), handlers.CreateItem)
	privateRoutes.Put("/items/:id", middlewares.CreateMiddleWare(), handlers.UpdateItem)
	privateRoutes.Delete("/items/:id", middlewares.CreateMiddleWare(), handlers.DeleteItem)
}
