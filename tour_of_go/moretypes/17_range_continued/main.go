package main

import (
	"fmt"
)

func main() {
	pow := make([]int, 10)
	for i := range pow { // You can omit getting the value
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow { // You can ignore the index or value by assigning it to '_'
		fmt.Printf("%d\n", value)
	}
}
