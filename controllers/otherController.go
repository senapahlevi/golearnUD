package controllers

import "github.com/gofiber/fiber"

func Other(c *fiber.Ctx) {
	c.SendString("hello from otherss")
}
