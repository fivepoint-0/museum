package main

import "fmt"

// generic impl
func printSlice[T any](slice []T) {
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
}

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	printSlice(primes)

	primes = append(primes, 23)
	printSlice(primes)

	primes = primes[3:]
	printSlice(primes)

	primes = primes[0:4]
	printSlice(primes)
}
