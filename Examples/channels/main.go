package main

import (
	"sync"

	cast "github.com/spf13/cast"
)

func output_number(number uint32) string {
	return "Your number is " + cast.ToString(number) + "!"
}

func channel_test_one(ch *chan bool) {
	*ch <- true
}

func main() {

	// Wait groups are simple to understand
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// For each 'i', we 'Add' a Workgroup item.
		// Conceptually, this is a piece of work to be done asynchronously.
		wg.Add(1)

		// declare anonymous function and call it immediately
		go func(a uint32) {
			// on the inside of the anonymous function,
			// we defer the 'Done' method for the Workgroup.
			// This allows us to tell the Workgroup this item
			// is done when the anonymous function finishes.
			defer wg.Done()

			// Call other functions we rely on
			output_number(a)
		}(cast.ToUint32(i))
	}

	// We wait for the entire Workgroup to finish.
	// Almost like a Thread.join() in other langs
	wg.Wait()

	var testChannel = make(chan bool)

	channel_test_one(&testChannel)

}
