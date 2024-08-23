package main

import (
	"fmt"
)

func main() {
	// if-blocks can have a short statement before the condition check.
	// h := 0 could be the value returned from another function, like:
	// h := somepackage.GetSomeValue()
	if h := 0; h < 10 {
		fmt.Println("Math still works.")
	}
}
