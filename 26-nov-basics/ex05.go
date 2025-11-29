package main

// **Problem 5: The ID Generator (Closures)** **Concept:** Functions as Values, State
// * Write a function named `idGenerator`. It should take no arguments but return a `func() int`.
// * The returned function must increment a counter every time it is called, preserving its state between calls.
// * In `main`, create a generator instance and call it 5 times, printing the returned ID each time.

import "fmt"

func idGenerator() func() int {
	count := 0

	CountIncr := func() int {
		c := &count
		(*c)++
		return count
	}
	return CountIncr
}

func main() {
	generator := idGenerator()

	for i := 0; i < 5; i++ {
		count := generator()
		fmt.Println(count)
	}
}
