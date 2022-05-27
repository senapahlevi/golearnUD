package main

import "fmt"

func main() {
	m := map[string]int{"key": 5}
	// fmt.Println(m)
	m["key"] = 10 //ini add/nambah stack
	// fmt.Println(m)
	delete(m, "key") //ini delete dan delete dari awal array
	m["newsss"] = 2  //ini juga add key
	fmt.Println(m)
}
