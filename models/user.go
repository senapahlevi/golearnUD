package models

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	// Password  string
	Password []byte `json:"-"` //- is not showing output after user regis/login in jsonusing these coz password on Register func authController avoid error if using string coz using hash
}
