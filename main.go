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

	// gorm.Model

}

//constraints like laravels length text,unique not same like others, etc
type User struct {
	gorm.Model        //so these gorm model will automatically add Id, delete/update At
	FirstName  string `gorm:"type:VARCHAR(30)"`
	LastName   string `gorm:"size:100"`
	Email      string `gorm:"unique"`
}
