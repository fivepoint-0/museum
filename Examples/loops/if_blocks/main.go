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

func main() {
	fmt.Println(sqrt(2), sqrt(4), sqrt(-2), sqrt(-4))
}
