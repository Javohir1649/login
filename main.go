package main

import (
	"github.com/Javohir1649/login/database"
	"github.com/Javohir1649/login/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
