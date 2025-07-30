package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	app2 "main/app"
	"main/database"
	"main/utils"
	"os"
)

const DUFAULT_PORT = "3000"

func main() {

	var app *fiber.App = app2.NewFiberApp()
	database.InitDatabase(utils.GetValue("DB_NAME"))
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = DUFAULT_PORT
	}
	app.Listen(fmt.Sprintf(":%s", PORT))
}
