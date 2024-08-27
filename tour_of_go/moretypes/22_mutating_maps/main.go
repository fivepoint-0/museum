package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	userMap := map[string]User{
		"admin":   {Name: "Luke", Age: 28},
		"guest":   {Name: "John", Age: 30},
		"partner": {Name: "Jin", Age: 33},
	}

	fmt.Printf("User Map: %v\n", userMap)
	userMap["admin"] = User{Name: "NewLuke", Age: 99}
	fmt.Printf("User Map after admin change: %v\n", userMap)

	guest := userMap["guest"]
	guest.Age = 65
	fmt.Printf("Guest: %v\n", guest)
	fmt.Printf("User map after guest change: %v\n", userMap)

	delete(userMap, "partner")
	fmt.Printf("User map after delete partner: %v\n", userMap)

	maybeUser, ok := userMap["maybeUser"]
	if ok {
		fmt.Printf("Found user --> %v\n", maybeUser)
	} else {
		fmt.Printf("Could not find user --> %v\n", maybeUser)
	}
}
