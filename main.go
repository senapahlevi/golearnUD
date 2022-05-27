package main

import "fmt"

func main() {
	m := map[string]int{"key": 5}
	fmt.Println(m)
	m["key"] = 10
	fmt.Println(m)

}
