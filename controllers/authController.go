package controllers

import (
	"goudemy/database"
	"goudemy/models"
	"goudemy/util"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
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

	// password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14) //we are just call ond userController funcSETPASSWORD
	// here

	//ada 2 cara masukkin nama/value ke struct nya
	user := models.User{
		// FirstName: "sena",
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleId:    1, //default nya example ceritanya when user regist default admin
		// Password:  data["password"],
		// Password: password, //opps weare not used anymore coz userController setPassword func reusable broo
	}
	//here reusable set password
	user.SetPassword(data["password"])
	//no 2 simpleda
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
	// if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil { //oops weare used reusable from func ComparePassword()
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	//jwt jadi ketika di postman response nya itu ya berupa token bukan nama,email,pass, bahaya uy
	///here below is generate jwt so these example if not using middleware were use these over and over to auth user generate jwt
	// claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{ //coz jwt.standardclaims deprecated HS256
	// 	// 	//buat login and issuer allow user found on db and strconv itoa convert int to string default
	// 	Issuer:    strconv.Itoa(int(user.Id)),
	// 	ExpiresAt: time.Now().Add(5 * time.Hour).Unix(), //for 1 day expires
	// })
	// token, err := claims.SignedString([]byte("secret")) //ini wajib make secret (bebeas supaya hacker tidak gampang tau isi token nya)

	///until up here

	//using middleware

	//ini wajib make secret (bebeas supaya hacker tidak gampang tau isi token nya)
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
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

// without middlewares
// type Claims struct {
// 	jwt.StandardClaims //ini klik isi ya  ada audience,expires,dll
// }
//with middleware not using claims struct just jwt.StandardClaims

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt") //get cookies dulu
	////below  theese for parse without middleware and slicing util
	// token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte("secret"), nil
	// })
	//up here
	//with middleware and util slicing
	id, _ := util.ParseJwt(cookie)

	//without middleware below here
	// if err != nil || !token.Valid { //and or err and token is invalid
	// 	//not authenticated
	// 	c.Status(fiber.StatusUnauthorized)
	// 	return c.JSON(fiber.Map{
	// 		"message": "unauthenticated !!",
	// 	})
	// }
	//up here

	// claims := token.Claims //ini isi nya iss (issuer) misal 3 = ini id data user id-3 bukan yg lain ya
	// return c.JSON(claims) //cukup sampe sini udah sesuai  dibawah ini opsional

	/////these opsional ga wajib
	//if mau nampilin iss aja dan nampilin isi si issuer (iss) komen dulu yak claims dan return nya
	var user models.User
	//without middleware
	// claims := token.Claims.(*Claims)
	//without middleware
	// database.DB.Where("id=?", claims.Issuer).First(&user)

	//with using middleware
	database.DB.Where("id=?", id).First(&user)
	return c.JSON(user)

}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",                         //no value for logout
		Expires:  time.Now().Add(-time.Hour), //these using minus because using past/ instant expires to  logout
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success logout",
	})
}
