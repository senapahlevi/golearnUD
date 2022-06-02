package controllers

import (
	"goudemy/database"
	"goudemy/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

//these we will create user but a bit different not like func Register
func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte("1234"), 14)
	user.Password = password
	database.DB.Create(&user)
	return c.JSON(user)
}
