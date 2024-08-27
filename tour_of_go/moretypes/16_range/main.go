package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	var users []User = []User{
		{Name: "Luke", Age: 28},
		{Name: "Richard", Age: 33},
		{Name: "Kent", Age: 50},
	}
	// you can iterate over values and indexes with the 'range' keyword
	// good for arrays and slices
	for i, v := range users {
		fmt.Printf("User (%v/%v) | Name = '%s'\tAge = '%v'\n", i+1, len(users), v.Name, v.Age)
	}
}
