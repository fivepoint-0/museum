package main

/*
	"comparable" is an interface that is implemented by all comparable
	types (booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types).
	The comparable interface may only be used as a type parameter constraint, not as the type of a variable.
*/

type Location struct {
	Lat, Long float64
}

type Person struct {
	Name     string
	Age      int
	Location *Location
}

func swap[T any](x, y *T) (T, T) {
	return *y, *x
}

func main() {
	p1 := Person{"John", 25, &Location{10, 20}}
	p2 := Person{"Jane", 26, &Location{30, 40}}

	println(p1.Name, p2.Name)
	p1, p2 = swap(&p1, &p2)
	println(p1.Name, p2.Name)
}
