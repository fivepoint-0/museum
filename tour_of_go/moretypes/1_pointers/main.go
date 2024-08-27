package main

import (
	"fmt"
)

func main() {
	var pointer *int

	number := 42

	pointer = &number

	fmt.Printf("Pointer address: %v\n", pointer)
	fmt.Printf("Pointer value (*pointer): %v\n", *pointer)
	*pointer = 21

	fmt.Printf("Pointer address: %v\n", pointer)
	fmt.Printf("Pointer value (*pointer): %v\n", *pointer)
	fmt.Printf("\nSame memory address, different value!\n")
}
