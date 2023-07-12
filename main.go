package main

import (
	"go-vercel/controllers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", controllers.AppName)
	app.Get("/app", controllers.AppName)

	err := app.Listen(":" + "9001")
	if err != nil {
		log.Fatalf("Error starting the server: %s", err.Error())
	}
}
