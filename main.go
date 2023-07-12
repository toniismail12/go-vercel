package main

import (
	"go-vercel/api/routers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routers.Setup(app)

	err := app.Listen(":" + "9001")
	if err != nil {
		log.Fatalf("Error starting the server: %s", err.Error())
	}
}
