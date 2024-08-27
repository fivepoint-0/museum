package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	userMap := make(map[string]User)

	userMap["admin"] = User{Name: "Luke", Age: 28}

	fmt.Printf("User map: %v\n", userMap)
	fmt.Printf("User admin: %v\n", userMap["admin"])  // actual value
	fmt.Printf("User admin: %v\n", userMap["adminn"]) // nil key
}
