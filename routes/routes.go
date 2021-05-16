package routes

import "github.com/gofiber/fiber"
import "../controllers"

func Setup(app *fiber.App){
	app.Get("/",controllers.Hello )
}
