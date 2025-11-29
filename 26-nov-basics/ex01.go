package main

// **Problem 1: The "Smart" Countdown** **Concept:** Loops, Conditionals Write a program that simulates a rocket launch sequence.

// * Initialize a loop counting down from **10 to 1**.
// * For every number, print the number.
// * **Condition A:** If the number is even, print "System Check OK" next to the number.
// * **Condition B:** If the number is **3**, print "Ignition Sequence Start".
// * After the loop finishes (at 0), print "Liftoff\!".

import "fmt"

func main() {
	//countdown loop
	for i := 10; i > 0; i-- {
		switch {
		case i%2 == 0:
			fmt.Println(i, " system check ok")
		case i == 3:
			fmt.Println(i, " Ignition Sequence Start")
		default:
			fmt.Println(i)
		}
	}
	fmt.Println("Liftoff!")

}
