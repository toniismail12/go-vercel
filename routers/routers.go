package routers

import (
	"go-vercel/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.AppName)
	app.Get("/app", controllers.AppName)
}
