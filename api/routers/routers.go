package routers

import (
	"go-vercel/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.AppName)
	app.Get("/app", controllers.AppName)
}
