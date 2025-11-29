package main

import "fmt"

// **Problem 6: The Safe Divider (Error Handling)** **Concept:** Multiple Returns, Error Interface

// * Write a function `Divide(a, b float64) (float64, error)`.
// * **Logic:** If `b` is 0, return `0` and an error using `fmt.Errorf("cannot divide by zero")`. Otherwise, return the result and `nil`.
// * In `main`, call the function twice: once with valid numbers, and once with a divisor of 0\.
// * Use `if err != nil` to handle the error gracefully in the console.

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil

}

func main() {
	val := []float64{20, 40, 20, 0}

	for i := 0; i < len(val); i += 2 {
		res, err := Divide(val[i], val[i+1])
		if err != nil {
			fmt.Println("error : ", err)
		} else {
			fmt.Println(res)
		}
	}

}
