package routes

import (
	"github.com/Javohir1649/login/controllers"
	"github.com/gofiber/fiber/v2"
)

// Setup yes
func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}
