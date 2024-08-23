package main

import (
	"fmt"
)

var (
	GLOBAL_A string
	GLOBAL_B string = "initialized"
)

func main() {
	fmt.Printf("GLOBAL_A = '%s'\n", GLOBAL_A)
	fmt.Printf("GLOBAL_B = '%s'\n", GLOBAL_B)
}
