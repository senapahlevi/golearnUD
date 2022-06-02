package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	// Password  string
	Password []byte `json:"-"` //- is not showing output after user regis/login in json using these coz password on Register func authController avoid error if using string coz using hash
}

//so in here set and compare func are slicing to reusable just call these func not type over and over again(painfuls)
func (user *User) SetPassword(password string) { //not return type
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword

}
func (user *User) ComparePassword(password string) error { // return type error
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))

}
