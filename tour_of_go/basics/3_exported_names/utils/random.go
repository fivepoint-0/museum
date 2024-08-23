package utils

import (
	"math/rand"
)

// available in the 'main' package.
func GetRandomInt() int {
	return getRandomInt()
}

// not available in the 'main' package
func getRandomInt() int {
	return rand.Int()
}
