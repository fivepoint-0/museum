package main

/*
	Constants are declared like other variables, except with the `const` keyword.
	Constants can be character, string, boolean, or numeric values.
	Constants cannot be declared using the `:=` syntax.
*/

func main() {

	// Constants can be declared using the `const` keyword
	const Pi = 3.14

	// Constants can be declared in a group
	const (
		Username = "admin"
		Password = "password"
	)

	println(Pi)
	println(Username)
	println(Password)

	// Username += "hey" // This line would not compile, as constants cannot be modified
}
