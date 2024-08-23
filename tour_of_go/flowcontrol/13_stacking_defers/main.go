package main

import (
	"fmt"
)

func WaitToExecute(i int) {
	fmt.Printf("The current int is: '%v'\n", i)
}

// the defers will execute in last-in-first-out order (LIFO)
// so for this example, the output will be:
// The current int is: '9'
// The current int is: '8'
// The current int is: '7'
// The current int is: '6'
// The current int is: '5'
// The current int is: '4'
// The current int is: '3'
// The current int is: '2'
// The current int is: '1'
// The current int is: '0'

func main() {
	for i := 0; i < 10; i++ {
		defer WaitToExecute(i)
	}
}
