package main

import (
	"fmt"
)

func main() {
	var names [10]string

	names[0] = "Ronald"
	names[1] = "Leticia"
	names[5] = "Koolaid Guy"

	slice1 := names[1:3]
	slice2 := names[5:]
	slice3 := names[:4]

	fmt.Printf("Slice 1: %v\n", slice1)
	fmt.Printf("Slice 2: %v\n", slice2)
	fmt.Printf("Slice 3: %v\n", slice3)
}
