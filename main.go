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
	//ini create
	user := User{
		FirstName: "senaa",
		LastName:  "pahlevii",
		Email:     "senapahlevi1@gmail.com",
	}
	DB.Create(&user)
}

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}
