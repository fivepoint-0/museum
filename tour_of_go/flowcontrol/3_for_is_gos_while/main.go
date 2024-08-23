package main

import (
	"fmt"

	"math/rand"
)

func main() {
	var sum int

	// a blank for condition statement in a for-loop will
	// just run forever unless broken (with a 'break')
	for {
		sum += rand.Intn(10)
		if sum > 90 {
			break
		}
	}

	fmt.Printf("The final sum is %v\n", sum)
}
