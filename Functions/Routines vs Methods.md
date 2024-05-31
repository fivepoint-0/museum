Golang semantically splits functions into particular use-cases.

## Routines
Routines are functions that take a set of values as inputs, that conform to a routine's signature. Here is an example:
```go
package main

import (
	"fmt"
)

func Add(a int, b int) {
	return a + b
}

func main() {
	fmt.Println("The result of 3 + 9 = ", Add(3, 9))
}
```

Output:
```sh
The result of 3 + 9 = 12
```
## Methods
Methods are an extension of Routines. Methods are specifically applied to `struct` objects. Basically, these are methods meant to work on a structure. These methods, paired with interfaces, can make a structure seem like a class-like object. Here is an example:

```go
package main

import (
	"fmt"
)

type Square struct {
	SideLength float64
}

type Rectangle struct {
	Width, Height float64
}

type AreaCalculator interface {
	Area() float64
}

func (s *Square) Area() float64 {
	return s.SideLength * s.SideLength
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// The interface is a standalone object we create, different from other langs
	var areaCalculator AreaCalculator
	
	// We initialize the structs like we do with classes in other langs
	var s Square = Square{ SideLength: 3 }
	var r Rectangle = Rectangle{ Width: 5, Height: 10 }
	
	// The calculator interface is referencing the square first
	areaCalculator = &s
	fmt.Println("Area of the square is    :", areaCalculator.Area())
	
	// Then the rectangle second. 
	areaCalculator = &r
	fmt.Println("Area of the rectangle is :", areaCalculator.Area())
}
```

The output of the program is:
```sh
Area of the square is    : 9
Area of the rectangle is : 50
```

We do not need to use pointers for the Area calculation. To see more about values vs pointers, see [Pointers vs Copies](<Pointers vs Copies.md>)