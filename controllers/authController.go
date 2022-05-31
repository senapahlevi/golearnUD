package controllers

import (
	"goudemy/models"

	"github.com/gofiber/fiber"
)

// func Hello(c *fiber.Ctx) {
// 	c.SendString("hello from auth")
// }

func Register(c *fiber.Ctx) {
	//ada 2 cara masukkin nama/value ke struct nya
	user := models.User{
		FirstName: "sena",
	}

	//no 2 simple
	user.LastName = "pahlevi"

	c.JSON(user)
}
