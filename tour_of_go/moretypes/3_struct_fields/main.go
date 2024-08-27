package main

import "fmt"

type User struct {
	Age  uint8
	Name string
}

func main() {
	var user1 User = User{Age: 20, Name: "Bob"}
	var user2 User = User{Age: 21, Name: "Alice"}
	var user3 User = User{Age: 43, Name: "Gerald"}

	var ageSum uint8 = 0

	ageSum += user1.Age
	ageSum += user2.Age
	ageSum += user3.Age

	fmt.Printf("Sum of the users' ages: %v\n", ageSum)
}
