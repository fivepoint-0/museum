package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Create a new instance of the Person struct
	p := Person{
		Name: "John Doe",
		Age:  25,
	}

	// Print the person's name and age
	fmt.Println(p.Name, "is", p.Age, "years old.")
}
