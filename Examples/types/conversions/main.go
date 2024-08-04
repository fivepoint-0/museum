package main

/*
	Type conversions
	The expression T(v) converts the value v to the type T.
*/

func main() {
	// Convert a float to an int
	f := 3.14
	i := int(f)
	println(i)

	// Convert an int to a float
	i = 42
	f = float64(i)
	println(f)

	// Convert a string to a byte slice
	s := "hello there"
	b := []byte(s)
	println(b)

	// Convert a byte slice to a string
	b = []byte{104, 101, 108, 108, 111}
	s = string(b)
	println(s)
}
