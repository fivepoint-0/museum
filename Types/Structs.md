Coming from an Object-oriented perspective, `struct` types in Golang can be looked at as properties and values associated with a `Class` object.

The **functions** that apply to **structs** are called **Methods** (see [Routines vs Methods](<../Functions/Routines vs Methods.md>))

Here is an example of a usage of a `struct`:
```go
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age uint32
}

func main() {
	var luke Person = Person { Name: "Luke Parker", Age: 28 }
	fmt.Println("The person's name is :", luke.Name)
	fmt.Println("The person's age is  :", luke.Age)
}
```

Output:
```
The person's name is : Luke Parker

The person's age is  : 28
```
