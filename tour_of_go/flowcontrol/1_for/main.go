package main

import (
	"fmt"
)

func main() {
	// i := 0 <-- the init statement
	// i < 10 <-- the condition expression
	// i++    <-- the post statement
	// braces required for the scope of the for-loop.
	// parenth. not needed for init;condition;post block
	for i := 0; i < 10; i++ {
		fmt.Printf("The index is %v\n", i)
	}
}
