package controllers

import (
	"goudemy/models"

	"github.com/gofiber/fiber/v2"

	"golang.org/x/crypto/bcrypt"
)

// func Hello(c *fiber.Ctx) {
// 	c.SendString("hello from auth")
// }

func Register(c *fiber.Ctx) error {

	//buat masukkin inputan dari user di postman localhost:3000/api/register
	var data map[string]string //[key]value 2 2 nya string buat input nya
	if err := c.BodyParser(&data); err != nil {
		//ini juga if err!=nil tapi disatuin biar pendek
		return err
	}
	//if password != confirm then give status bad request and error message

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "password did not match"})

	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	//ada 2 cara masukkin nama/value ke struct nya
	user := models.User{
		// FirstName: "sena",
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		// Password:  data["password"],
		Password: password,
	}
	//no 2 simple
	// user.LastName = "pahlevi"

	// return c.JSON(user)
	return c.JSON(user)
}
