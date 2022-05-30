package controllers

import "github.com/gofiber/fiber"

func Hello(c *fiber.Ctx) {
	c.SendString("hello from auth")
}
