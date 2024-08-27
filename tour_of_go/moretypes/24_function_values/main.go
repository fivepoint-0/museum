package main

import (
	"fmt"
)

// Functions are values too. They can be passed around just like other values.

func val(fn func(int, int) int, a, b int) int {
	return fn(a, b)
}

func main() {
	func1 := func(a, b int) int { return a + b }
	func2 := func(a, b int) int { return a * b }
	func3 := func(a, b int) int { return a - b }

	val1 := val(func1, 3, 4)
	val2 := val(func2, 3, 4)
	val3 := val(func3, 3, 4)

	fmt.Printf("val1 = '%v'\n", val1)
	fmt.Printf("val2 = '%v'\n", val2)
	fmt.Printf("val3 = '%v'\n", val3)
}
