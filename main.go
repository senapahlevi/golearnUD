package main

import (
	"goudemy/database"
	"goudemy/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// const DNS = "host=127.0.0.1 user=postgres password=123456789 dbname=goudemy port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// DB, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	// if err != nil {
	// 	// fmt.Println(err.Error())
	// 	panic("Cannot connect to DB")
	// }
	// fmt.Println(DB)
	// // DB.Migrator().DropTable(&User{}, &Book{})
	// // DB.Migrator().CreateTable(&User{}, &Book{})

	// // var name *string = new(string) //wajib make new inisialisasi new biar g error

	// // *name = "hello"
	// // fmt.Println(name, "ini siapa", &name)
	// app := fiber.New()

	// // Routes
	// app.Get("/", hello)

	// // start server
	// app.Listen(":3000")

	database.Connect()
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":3000")
}

// Handler
