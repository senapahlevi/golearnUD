package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	const DNS = "host=127.0.0.1 user=postgres password=123456789 dbname=goudemy port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	DB, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		// fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{})
	// //ini create
	// user := User{
	// 	FirstName: "senaa",
	// 	LastName:  "pahlevii",
	// 	Email:     "senapahlevi1@gmail.com",
	// }
	// DB.Create(&user)

	// //update first comment above we not gonna create user again just updates and deletes
	// user := User{
	// 	Id:        1,
	// 	FirstName: "xenna",
	// 	LastName:  "levi",
	// 	Email:     "xenna@gmail.com",
	// }
	// DB.Updates(&user) //pass user to params (&user)

	//deletes only find the id so the data will deletes immediately comment above update and create
	user := User{
		Id: 1,
	}
	DB.Delete(&user)
}

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}
