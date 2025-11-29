package main

// **Problem 10: The Payment Gateway (Interfaces)** **Concept:** Polymorphism

// * Define an interface `PaymentMethod` with a single method `Pay(amount int)`.
// * Define two structs: `CreditCard` and `Bitcoin`.
// * Implement the `Pay` method for both (e.g., CreditCard prints "Paid $X via Card", Bitcoin prints "Paid $X via BTC").
// * Write a function `Process(p PaymentMethod, amount int)` that takes the interface as an argument.
// * In `main`, pass both a `CreditCard` and a `Bitcoin` instance to `Process`.
