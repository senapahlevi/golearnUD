package routes

import (
	"goudemy/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// app.Get("/", controllers.Hello)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	// app.Get("/other", controllers.Other)
} //biar bisa kepanggil dan no error antar package make huruf capital

// func Setup(){}
