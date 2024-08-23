package main

import (
	"fmt"
	"math"
)

func main() {
	// Switch without a condition is the same as `switch true`.
	switch {
	case 1 < 0:
		fmt.Printf("1 IS less than 0!\n")
	case math.Pow(3, 3) == 27:
		fmt.Printf("3 to the 3 is 27!\n")
	case false == true:
		fmt.Printf("OOF")
	default:
		fmt.Printf("Nah fam")
	}
}
