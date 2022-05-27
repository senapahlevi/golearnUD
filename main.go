package main

import (
	"database/sql"

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
	DB.Migrator().DropTable(&User{})
	DB.Migrator().CreateTable(&User{})
	// gorm.Model
	//test buat nullstring(sql.nullable) bukan nullable
	user := User{
		Email: sql.NullString{
			String: "example@gmail.com",
			Valid:  true, //true itu biar ga error
		},
	}
	DB.Create(&user)
}

//nullable string but when using these db will continue migrate/save/update data with empty
// type User struct {
// 	gorm.Model        //so these gorm model will automatically add Id, delete/update At
// 	FirstName  string `gorm:"type:VARCHAR(30); null;"`
// 	LastName   string `gorm:"size:100; default:'Smith weberjenkinson'"`
// 	Email      string `gorm:"unique; not null;"`
// }

//using sql.NullString when condition null/not null will gives warning or like handling errors so the data will not save with empty data
type User struct {
	gorm.Model                //so these gorm model will automatically add Id, delete/update At
	FirstName  sql.NullString `gorm:"type:VARCHAR(30); null;"`
	LastName   sql.NullString `gorm:"size:100; default:'Smith weberjenkinson'"`
	Email      sql.NullString `gorm:"unique; not null;"`
}
