package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomFloat32 := rand.Float32()
	fmt.Println("The random float32 is", randomFloat32)

	randomFloat64 := rand.Float64()
	fmt.Println("The random float64 is", randomFloat64)

	randomInt := rand.Int()
	fmt.Println("The random int is", randomInt)

	randomIntn := rand.Intn(10)
	fmt.Println("The random intn is", randomIntn)
}
