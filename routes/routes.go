package routes

import (
	"goudemy/controllers"
	"goudemy/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// app.Get("/", controllers.Hello)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)

	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	//roles
	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)

	app.Get("/api/role/:id", controllers.GetRole)
	app.Put("/api/role/:id", controllers.UpdateRole)
	app.Delete("/api/role/:id", controllers.DeleteRole)

	// app.Get("/other", controllers.Other)
} //biar bisa kepanggil dan no error antar package make huruf capital

// func Setup(){}
