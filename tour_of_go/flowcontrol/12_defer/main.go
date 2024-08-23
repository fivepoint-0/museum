package main

import (
	"fmt"
)

func WaitToExecute() {
	fmt.Println("The calling block is done executing!")
}

func main() {
	defer WaitToExecute()

	fmt.Println("This is the calling block, will execute before WaitToExecute()")
}
