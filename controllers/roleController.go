package controllers

import (
	"goudemy/database"
	"goudemy/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//these we will create user but a bit different not like func Register
func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)
	return c.JSON(roles)
}

func CreateRole(c *fiber.Ctx) error {
	var role models.Role

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	database.DB.Create(&role)
	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //we get id from params convert back again strting to unsigned avoid error when assign user := models.User{}
	role := models.Role{
		Id: uint(id),
	}
	database.DB.Find(&role) //find user ex user/4 , user/3
	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //we get id from params convert back again strting to unsigned avoid error when assign user := models.User{}
	role := models.Role{
		Id: uint(id),
	}

	if err := c.BodyParser(&role); err != nil {
		return err
	} //to regist input new for update
	database.DB.Model(&role).Updates(role)
	return c.JSON(role)
}
func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //we get id from params convert back again strting to unsigned avoid error when assign user := models.User{}
	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)
	return c.JSON(role)
}
