package main

import (
	"fmt"
)

func addString(strToAppend string) func(string) string {
	return func(s string) string {
		return s + strToAppend
	}
}

func main() {
	helloer, goodbyer := addString(", hello!"), addString(", goodbye!")
	names := []string{
		"Luke",
		"Jane",
		"Marissa",
	}

	for _, v := range names {
		fmt.Printf(helloer(v) + "\n")
		fmt.Printf(goodbyer(v) + "\n")
	}
}
