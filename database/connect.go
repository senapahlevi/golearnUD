package database

import (
	"goudemy/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	const DNS = "host=127.0.0.1 user=postgres password=123456789 dbname=goudemy port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to DB")
	}
	db.AutoMigrate(&models.User{})
}
