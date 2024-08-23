package main

// Golang server wont let me save individual imports lol. I get the idea though.
import (
	"fmt"
	"math"
)

func main() {
	sqrt4 := math.Sqrt(4)
	fmt.Printf("The square root of '4' is %f", sqrt4)
}
