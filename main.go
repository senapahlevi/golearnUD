package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Addrr     Address
}

type Address struct {
	Country string
	City    string
}

func main() {
	hasilAdd := Address{
		Country: "germany",
		City:    "Berlin",
	}
	user := User{
		ID:        1,
		FirstName: "First",
		LastName:  "lassst",
		Addrr:     hasilAdd,
	}
	fmt.Println(user)
}
