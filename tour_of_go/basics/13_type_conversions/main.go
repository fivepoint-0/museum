package main

import "fmt"

func main() {
	// You can use T(X) to convert X to type T
	// assignment between items of different type requires an explicit conversion
	var a int = 12345
	var b float64 = float64(a)
	var c uint = uint(b)

	fmt.Printf("The value of c = '%v'\n", c)
}
