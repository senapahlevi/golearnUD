package controllers

import (
	"goudemy/database"
	"goudemy/models"

	"github.com/gofiber/fiber/v2"
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
	// password, _ := bcrypt.GenerateFromPassword([]byte("1234"), 14) //not reusable
	user.SetPassword("1234") //reusable broo
	// user.Password = password not used coz reusable
	database.DB.Create(&user)
	return c.JSON(user)
}
