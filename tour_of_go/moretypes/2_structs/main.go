package main

import "fmt"

type User struct {
	Age  uint8
	Name string
}

func main() {
	var user User = User{Age: 1, Name: "Baby"}

	fmt.Printf("User age: %v\n", user.Age)
	fmt.Printf("User name: %s\n", user.Name)
}
