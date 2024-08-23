package main

import (
	"fmt"
)

// did a generic implementation instead. just looked up the generics syntax
func swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	string1 := "Hello"
	string2 := "World"

	fmt.Printf("Before Swap | string1 = '%s', string2 = '%s'\n", string1, string2)

	string1, string2 = swap(string1, string2)

	fmt.Printf("After Swap | string1 = '%s', string2 = '%s'\n", string1, string2)
}
