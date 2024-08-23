package main

import (
	"fmt"

	"math/rand"
)

func main() {
	var sum int = 0

	// the init and post conditions are optional
	for sum < 50 {
		fmt.Printf("Adding %v to the sum\n", sum)
		sum += rand.Intn(10)
		fmt.Printf("The sum is %v\n", sum)
	}

	fmt.Printf("The final sum is %v\n", sum)
}
