# Multiple Return Values

In Go, you can return more than one value from a function, like the `swap` function:

```go
package main

func swap[T any](x, y *T) (T, T) {
	return *y, *x
}
```