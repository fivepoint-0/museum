package main

import "fmt"

type User struct {
	Age  uint8
	Name string
}

func main() {
	user := User{Age: 1, Name: "John"}
	pointer := &user
	pointer.Age = 33 // we can set the values through the pointer like `pointer.Age`, without having to dereference like `(*pointer).Age`
	fmt.Println(pointer)
	fmt.Println(user)
}
