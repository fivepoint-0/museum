package main

import "fmt"

// instead of 'a int, b int', we can say 'a, b int'
// because 'a' and 'b' are the same type.
func Add(a, b int) int {
	return a + b
}

func main() {
	var num1 int = 4
	var num2 int = 5

	fmt.Printf("The sum is: %v\n", Add(num1, num2))
}
