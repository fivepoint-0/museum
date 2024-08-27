package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	userMap := map[string]User{
		"admin": {
			Name: "Luke",
			Age:  28,
		},
		"guest": {
			Name: "Jane",
			Age:  30,
		},
	}

	fmt.Printf("User map: %v\n", userMap)
}
