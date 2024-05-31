In Golang, you can either accept a Pointer to a structure or pass a Copy of the structure into the function. 

Here is a note on Mutability from https://medium.com/@meeusdylan/when-to-use-pointers-in-go-44c15fe04eac:
# Mutability

The only way to mutate a variable that you pass to a function is by passing a pointer. By default, the pass-by-value means that changes you make are on the copy you’re working on. Thus, they are not reflected in the calling function.

Let's test this out.

## Copies
```go
package main

import (
	"fmt"
)

func UseCopyValue(i int) int {
	i++
	return i
}

func main() {
	a := 5
	UseCopyValue(a)
	fmt.Println("The value of 'a' is:", a)
}
```

Output:
```sh
The value of 'a' is: 5
```

Note, in the `UseCopyValue` function above, we are not taking in a value as a pointer. We are taking in the **value** passed to the function. 

In English, the rough explanation of why the output is still `5` is because we store the value of `5` to `a`, and then pass the value of `5` to `UseCopyValue` function instead of the variable `a`. To run the point home, `5` is what we are passing here, not `a`. If we passed `a` as a pointer, the function would act on the variable `a` instead of the value contained in `a` which is `5`.

TL;DR - Pointers are references to objects, non-pointers are just values.

## Pointers
In a pointers example similar to the above `Copies` example, here is the new code:
```go
package main

import (
	"fmt"
)

func UsePointerValue(i *int) {
	*i++
}

func main() {
	a := 5
	UsePointerValue(&a)
	fmt.Println("The value of 'a' is:", a)
}
```

Output:
```sh
The value of 'a' is: 6
```

The `UsePointerValue` function doesn't have a return type. The reason is because we edited the value of `a` inside the function. There is no need to return `a` again. We could return a value if we wanted, but that can get a little messy with doing more than one type of task within a function. 
