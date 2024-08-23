package main

import (
	"fmt"
)

/*
  Constants are declared like variables, but with the `const` keyword.

  Constants can be character, string, boolean, or numeric values.

  Constants cannot be declared using the `:=` syntax.
*/

const variable string = "hey there"

func main() {
	// variable = "hey" // error: cannot assign to variable (neither addressable nor a map index expression)
	fmt.Printf("%s, person\n", variable)
}
