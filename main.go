package main

import (
	"errors"
	"fmt"
)

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
	//function
	// result, err := Divide(4, 0)
	result, zzz := Divide(4, 0)
	if zzz != nil {
		fmt.Println(zzz)
	}
	fmt.Println(result)
}

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide with zero !!!!!")
	}
	return a / b, nil
}
