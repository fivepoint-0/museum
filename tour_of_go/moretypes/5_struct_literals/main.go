package main

import "fmt"

type Rectangle struct {
	Width  uint32
	Height uint32
}

func main() {
	rectangle1 := Rectangle{Width: 1}  // Height defaults to '0'
	rectangle2 := Rectangle{Height: 1} // Width defaults to '0'
	rectangle3 := Rectangle{}          // Both Height and Width default to '0'
	// rectangle4 := Rectangle{1} // This is an error, because the struct construction doesn't know whether you're setting Width or Height.
	rectangle4 := &Rectangle{Width: 12, Height: 20} // This becomes a pointer to a new Rectangle

	fmt.Printf("Rectangle 1: %v\n", rectangle1)
	fmt.Printf("Rectangle 2: %v\n", rectangle2)
	fmt.Printf("Rectangle 3: %v\n", rectangle3)
	fmt.Printf("Rectangle 4: %v\n", rectangle4)
}
