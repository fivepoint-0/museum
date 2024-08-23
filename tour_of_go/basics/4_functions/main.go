package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func main() {
	var num1 int = 4
	var num2 int = 5

	fmt.Printf("The sum is: %v\n", Add(num1, num2))
}
