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
	var roleDto fiber.Map

	// var role models.Role

	// if err := c.BodyParser(&role); err != nil {
	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}
	list := roleDto["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))
	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}
	role := models.Role{
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}
	database.DB.Create(&role)
	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //we get id from params convert back again strting to unsigned avoid error when assign user := models.User{}
	role := models.Role{
		Id: uint(id),
	}
	// database.DB.Find(&role) //find user ex user/4 , user/3
	database.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //we get id from params convert back again strting to unsigned avoid error when assign user := models.User{}
	// role := models.Role{
	// 	Id: uint(id),
	// }
	var roleDto fiber.Map
	// if err := c.BodyParser(&role); err != nil {
	if err := c.BodyParser(&roleDto); err != nil {
		return err
	} //to regist input new for update
	list := roleDto["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))
	for i, permissionId := range list {
		id, _ := permissionId.(float64)
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}
	var result interface{}
	//deletes the old one to avoid redundant/still saving in db after update assign value new on roles_permission
	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)
	//after delete the old one now create new
	role := models.Role{
		Id:          uint(id),
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}
	database.DB.Model(&role).Updates(role)
	return c.JSON(role)
}
func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //we get id from params convert back again strting to unsigned avoid error when assign user := models.User{}
	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)
	// return c.JSON(role)
	return nil
}
