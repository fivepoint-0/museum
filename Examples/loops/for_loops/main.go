package main

import (
	"fmt"
)

func main() {
	sum := 0

	// Parameters = init; condition; post
	for i := 0; i < 100; i++ {
		sum += i
	}

	fmt.Println(sum)

	second_sum := 0

	// Init and post paramaters are optional
	for second_sum < 1000 {
		second_sum += 10
	}

	fmt.Println(second_sum)
}

/* Forever loops: (infinite loops)

for {

}

*/
