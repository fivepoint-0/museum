package main

import (
	"fmt"
)

func main() {
	names := [5]string{
		"DrWho",
		"Robert",
		"Catdog",
		"TonytheTiger",
		"Superwoman",
	}

	slice1 := names[1:3]

	fmt.Printf("Names: %v\n", names)
	fmt.Printf("Slice 1: %v\n", slice1)

	slice1[0] = "NewName" // this changes the value in the []names array, not the slice.
	// This means slices are like references to arrays, not copies of them.

	fmt.Printf("Names after slice value change: %v\n", names)
}
