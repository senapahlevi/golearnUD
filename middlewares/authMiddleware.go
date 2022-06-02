package middlewares

import (
	"goudemy/util"

	"github.com/gofiber/fiber/v2"
)

//these func run before get response on backend side / to make reusable these function lie's in middle between in client request and backend get respons
func IsAuthenticated(c *fiber.Ctx) error { //will check is authenticated or not?
	cookie := c.Cookies("jwt")
	//if user not authenticated/not login/after logout and in postman Get user will throw these
	if _, err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated!!!!",
		})
	}
	return c.Next() //if user are login will pass into func User at authController
}
