package main

import (
	"fmt"
	"time"

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

	// gorm.Model
	user := User{
		Model: gorm.Model{
			CreatedAt: time.Now(),
		},
	}
	fmt.Println(user)
}

type User struct {
	gorm.Model //so these gorm model will automatically add Id, delete/update At
	FirstName  string
	LastName   string
	Email      string
}
