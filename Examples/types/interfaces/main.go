package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Animal struct {
	Name string
	Age  int
}

type Speaker interface {
	Speak()
}

func (p *Person) Speak() {
	fmt.Println(p.Name, "says hello!")
}

func (a *Animal) Speak() {
	fmt.Println(a.Name, "makes a noise!")
}

func main() {
	p := &Person{
		Name: "John Doe",
		Age:  25,
	}

	a := &Animal{
		Name: "Dog",
		Age:  3,
	}

	// Create a slice of speakers (remember a Speaker is an interface, but used here more like a trait)
	speakers := []Speaker{p, a}

	// Loop through the slice and call the Speak method
	for _, s := range speakers {
		s.Speak()
	}
}
