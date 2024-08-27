package main

import "fmt"

func main() {
	array_literal := [5]int{1, 2, 3, 4, 5}
	slice_literal := []int{1, 2, 3, 4, 5} // this slice literal is the same as
	// the array literal above, except without length

	fmt.Printf("Array literal: %v\n", array_literal)
	fmt.Printf("Slice literal: %v\n", slice_literal)
}
