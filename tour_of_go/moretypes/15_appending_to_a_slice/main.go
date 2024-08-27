package main

import (
	"fmt"
)

func main() {
	slice := []int{}

	slice = append(slice, 4)

	fmt.Printf("Slice: %v\n", slice)
}
