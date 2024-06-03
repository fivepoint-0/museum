package main

import (
	"sync"

	cast "github.com/spf13/cast"
)

func output_number(number uint32) string {
	return "Your number is " + cast.ToString(number) + "!"
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(a uint32) {
			defer wg.Done()
			output_number(a)
		}(cast.ToUint32(i))
	}

	wg.Wait()

}
