package main

import "fmt"

func main() {
	// To create a dynamically sized array, we can create a slice with `make`:
	dynSized := make([]int, 10)

	fmt.Printf("dynSized slice array: %v\n", dynSized)
}
