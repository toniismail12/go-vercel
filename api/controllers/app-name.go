package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func AppName(c *fiber.Ctx) error {

	c.Status(200)
	return c.JSON(fiber.Map{
		"app_name": "go vercel",
		"desc":     "go vercel cuyy",
	})

}
