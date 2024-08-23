package main

import "fmt"

// Had to look up the '(a, b string)' portion.
// That is the 'top' of the function apparently, according to the README.
func returnSomeStuff() (a, b string) {
	a, b = "", ""

	a = a + "Hello"
	b = b + "World"

	return
}

func main() {
	var a, b string = returnSomeStuff()

	fmt.Println(a)
	fmt.Println(b)

}
