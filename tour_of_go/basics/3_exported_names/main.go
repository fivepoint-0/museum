package main

import (
	"fmt"

	"github.com/fivepoint-0/learning-golang/tour_of_go/basics/3_exported_names/utils"
)

func main() {
	// Learned Printf should get the \n if you want a newline before the next
	// bash prompt comes back up.
	fmt.Printf("Get random integer: %v\n", utils.GetRandomInt())
}
