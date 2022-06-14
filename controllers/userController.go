package controllers

import (
	"goudemy/database"
	"goudemy/models"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//these we will create user but a bit different not like func Register
func AllUsers(c *fiber.Ctx) error {
	//add pagination
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit //where to start page from limit = ex: 5 and when we want to next page page-1
	var users []models.User
	var total int64
	// database.DB.Find(&users)
	// database.DB.Preload("Role").Find(&users)

	// return c.JSON(users) //only return users

	database.DB.Preload("Role").Offset(offset).Limit(limit).Find(&users) //we want to showing roles , get query for counts
	database.DB.Model(&models.User{}).Count(&total)                      //
	return c.JSON(fiber.Map{                                             //we want to return page showing page=3 example
		"data": users,
		"meta": fiber.Map{
			"page":      page,
			"total":     total,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
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
	// database.DB.Find(&user) //find user ex user/4 , user/3
	database.DB.Preload("Role").Find(&user) //find user ex user/4 , user/3 oreload to find roles query before user
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
