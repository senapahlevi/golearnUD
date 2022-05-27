package main

import (
	"fmt"

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
	//quering these comment immediately after succes input 3 data into postgresql
	// users := []User{
	// 	{
	// 		FirstName: "Sena",
	// 		LastName:  "Pahlevi",
	// 		Email:     "senapahlevi@gmail.com",
	// 	},
	// 	{
	// 		FirstName: "Xenna",
	// 		LastName:  "Smith",
	// 		Email:     "xennasmith@gmail.com",
	// 	},
	// 	{
	// 		FirstName: "William",
	// 		LastName:  "Blake",
	// 		Email:     "williamblake@gmail.com",
	// 	},
	// }
	// //these assign above 3 datas
	// for _, user := range users { //these using underscore(_) if didnt want to use variable because on docs
	// 	//for range always using variables to avoid error coz golang is type programming language (Strict not like .js)
	// 	DB.Create(&user)
	// }

	//query look like filter (first,last,where(similiarity like filters))
	user := User{} //these var user will assign and to printout like .push() create new array zero to got output
	//first() showing first datas
	// DB.First(&user) //always passing &user coz to asssign into var user
	// fmt.Println(user)

	//last showing last data
	// DB.Last(&user) //always passing &user coz to asssign into var user
	// fmt.Println(user)

	//query filter using where
	DB.Where("last_name", "Smith").First(&user) //why last_name not Lastname like these struct? coz it point out the column postgress last_name
	fmt.Println(user)

}

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}
