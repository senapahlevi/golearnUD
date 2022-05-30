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
	// DB.AutoMigrate(&User{})
	//these used to drop and migrate like migrate:refresh laravel
	// DB.Migrator().DropTable(&User{}, &Address{})
	// DB.Migrator().CreateTable(&User{}, &Address{})
	DB.Migrator().DropTable(&User{}, &Book{})
	DB.Migrator().CreateTable(&User{}, &Book{})
	//kalo ga bisa migrator make auto migrate kadang suka aneh sih antara itu biar bisa migrate gitu
	// DB.AutoMigrate(&User{}, &Book{})
	// gorm.Model
	//test buat nullstring(sql.nullable) bukan nullable
	// user := User{
	// 	FirstName: "sena",
	// 	LastName:  "pahlevvvvv",
	// 	Email:     "senapahlevi2@gmail.com",
	// 	Address: Address{
	// 		Name: "Main str.",
	// 	},
	// }
	user := User{
		FirstName: "sena",
		LastName:  "pahlevvvvv",
		Email:     "senapahlevi2@gmail.com",
		Books: []Book{
			{
				Title: "Bookeeee", //masuk ke user_books
			},
			{
				Title: "Buku harry potter 2", //masuk ke user_books
			},
		},
	}
	DB.Create(&user)
	u := User{}
	// DB.Preload("Address").First(&u)
	DB.Preload("Books").First(&u)
	fmt.Println(u)
}

//nullable string but when using these db will continue migrate/save/update data with empty
// type User struct {
// 	gorm.Model        //so these gorm model will automatically add Id, delete/update At
// 	FirstName  string `gorm:"type:VARCHAR(30); null;"`
// 	LastName   string `gorm:"size:100; default:'Smith weberjenkinson'"`
// 	Email      string `gorm:"unique; not null;"`
// }

//using sql.NullString when condition null/not null will gives warning or like handling errors so the data will not save with empty data
//many to many
type User struct {
	gorm.Model //so these gorm model will automatically add Id, delete/update At
	// ID        uint
	FirstName string `gorm:"type:VARCHAR(30); null;"`
	LastName  string `gorm:"size:100; default:'Smith weberjenkinson'"`
	Email     string `gorm:"unique; not null;"`
	// Address   Address `gorm:"foreignKey:UserId"`
	Books []Book `gorm:"many2many:user_books"` //many to many
}
type Book struct {
	gorm.Model //so these gorm model will automatically add Id, delete/update At
	// ID     uint
	ID    uint
	Title string
}
