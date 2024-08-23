package main

import (
	"fmt"
)

/*
 *   `0` for numeric types,
 *   `false` for the boolean type, and
 *   `""` (the empty string) for strings.
 */
func main() {
	var (
		a int
		b bool
		c string
	)
	fmt.Printf("Value of a = '%v'\n", a)
	fmt.Printf("Value of b = '%v'\n", b)
	fmt.Printf("Value of c = '%s'\n", c)
}
