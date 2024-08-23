package main

import (
	"fmt"
	"math"
)

func main() {
	if h := math.Pow(1.414, 2); h < 2 {
		h = h * 2
		fmt.Printf("The value of h = '%v'\n", h)
	} else {
		fmt.Println("'h'**2 is less than 2!")
	}
}
