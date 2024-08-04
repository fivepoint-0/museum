# Common Parameters

In Go, you can omit the type of preceding arguments if the following argument(s) are of the same type. In the below example, you can see that we can change `x int, y int` with `x, y int`:

```go
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

