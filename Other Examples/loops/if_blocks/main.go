package main

import (
	"fmt"
	"math"
)

func sqrt(i float64) string {
	if i < 0 {
		return sqrt(-i) + "i"
	}

	return fmt.Sprint(math.Sqrt(i))
}

/*
Go's if statements are like its for loops;
the expression need not be surrounded by parentheses ( ) but the braces { } are required.

You can make a short statement before the condition that will be executed before the condition.
*/
func pow(x, n, lim float64) float64 {
	// v is a variable that is only available in the if scope
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(4), sqrt(-2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
