package main

import (
	"fmt"

	"math/rand"
)

func main() {
	var h int = rand.Intn(10)
	// the switch evaluation order goes from top to bottom,
	// stopping when a case matches the switch condition
	switch h {
	case 0:
		fmt.Println("h < 5 !")
	case 1:
		fmt.Println("h < 5 !")
	case 2:
		fmt.Println("h < 5 !")
	case 3:
		fmt.Println("h < 5 !")
	case 4:
		fmt.Println("h < 5 !")
	case 5:
		fmt.Println("h == 5 !")
	case 6:
		fmt.Println("h > 5 !")
	case 7:
		fmt.Println("h > 5 !")
	case 8:
		fmt.Println("h > 5 !")
	case 9:
		fmt.Println("h > 5 !")
	default:
		fmt.Println("I didn't account for this case")
	}
}
