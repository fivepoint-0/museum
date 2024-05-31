Interfaces in Golang are a bit weird in relation to other programming languages. In a way, in Golang, you initialize a variable that is of the type of an interface defined, and then set the value of the variable to a struct who has an implementation for that interface.

A quick example can be shown in [Routines vs Methods](<../Functions/Routines vs Methods.md>), but here is a different example:

```go
package main

import (
	"fmt"
)

type Person struct {
    FirstName   string
    LastName    string
    Age         uint32
}

type Greeter interface {
	Greet()
}

func (p *Person) Greet() {
	fmt.Println("Hello", p.FirstName, p.LastName, "!")
}

func main() {
	// "greeter" is of type "Greeter", which has a "Greet()" method defined
	var greeter Greeter
	
	// A Person has an implementation of the Greeter by means of the Greet() method
	var luke Person = Person {
        FirstName: "Luke",
        LastName:  "Parker",
        Age:       28,
    }
	
	// Set the greeter to the struct instance defined above
	greeter = &luke
	
	// Call the Greet() function.
	greeter.Greet()
}
```

Output:
```
Hello Luke Parker!
```
