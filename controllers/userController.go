package controllers

import (
	"goudemy/database"
	"goudemy/models"
	"strconv"

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

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //we get id from params convert back again strting to unsigned avoid error when assign user := models.User{}
	user := models.User{
		Id: uint(id),
	}
	database.DB.Find(&user) //find user ex user/4 , user/3
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //we get id from params convert back again strting to unsigned avoid error when assign user := models.User{}
	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	} //to regist input new for update
	database.DB.Model(&user).Updates(user)
	return c.JSON(user)
}
func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //we get id from params convert back again strting to unsigned avoid error when assign user := models.User{}
	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)
	return c.JSON(user)
}
