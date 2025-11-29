package main

import "fmt"

// **Problem 10: The Payment Gateway (Interfaces)** **Concept:** Polymorphism

// * Define an interface `PaymentMethod` with a single method `Pay(amount int)`.
// * Define two structs: `CreditCard` and `Bitcoin`.
// * Implement the `Pay` method for both (e.g., CreditCard prints "Paid $X via Card", Bitcoin prints "Paid $X via BTC").
// * Write a function `Process(p PaymentMethod, amount int)` that takes the interface as an argument.
// * In `main`, pass both a `CreditCard` and a `Bitcoin` instance to `Process`.

type PaymentMethod interface {
	pay(amount int)
}

type CreditCard struct{}

type Bitcoin struct{}

func (c CreditCard) pay(amount int) {
	fmt.Printf("Paid %d via Card\n", amount)
}

func (b Bitcoin) pay(amount int) {
	fmt.Printf("Paid %d via bitcoin\n", amount)
}

func Process(p PaymentMethod, amount int) {
	p.pay(amount)
}

func main() {
	c := CreditCard{}
	b := Bitcoin{}
	Process(c, 100)
	Process(b, 150)
}
