package main

import (
	"fmt"
)

func main() {
	var a int = 67
	b := a

	fmt.Printf("The type of b is implicit: type(b) = '%T'\n", b)
}
