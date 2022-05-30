package routes

import (
	"goudemy/controllers"

	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
	app.Get("/other", controllers.Other)
} //biar bisa kepanggil dan no error antar package make huruf capital
