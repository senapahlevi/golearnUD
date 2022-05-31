package controllers

import (
	"goudemy/database"
	"goudemy/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

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
	//store to db after input post api/register
	database.DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	//[key]value 2 2 nya string buat input nya
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	//conditional coz login only input email and password
	//conditionla email
	database.DB.Where("email = ?", data["email"]).First(&user)
	//when user not found

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "not found",
		})
	}
	//compare password input with stored password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	//jwt jadi ketika di postman response nya itu ya berupa token bukan nama,email,pass, bahaya uy
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{ //coz jwt.standardclaims deprecated HS256
		// 	//buat login and issuer allow user found on db and strconv itoa convert int to string default
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(5 * time.Hour).Unix(), //for 1 day expires
	})
	token, err := claims.SignedString([]byte("secret")) //ini wajib make secret (bebeas supaya hacker tidak gampang tau isi token nya)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
		// return c.SendString("jancuk")

	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	// return c.JSON(user) //please dont use these coz return what user type email and pass

	// return c.JSON(token)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
