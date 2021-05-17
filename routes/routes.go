package routes

import "github.com/gofiber/fiber"
import "../controllers"

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
}
