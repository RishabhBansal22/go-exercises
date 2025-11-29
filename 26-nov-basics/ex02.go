package main

// **Problem 2: Dynamic Inventory (Slices)** **Concept:** Slices, Append, Slicing

// * Create an empty slice of strings named `cart`.
// * Append three items: "Laptop", "Mouse", "Keyboard".
// * Print the cart contents, its length, and its capacity.
// * **Action:** Remove the first item ("Laptop") using slice syntax (`cart[low:high]`) so only the last two items remain.
// * Print the updated cart.

import "fmt"

var cart []string

func addToCart(items []string) {
	p := &cart
	*p = append(cart, items...)
}

func remFirst() {
	p := &cart
	*p = cart[1:]

}

func main() {
	items := []string{"mouse", "keyboard", "cpu"}

	addToCart(items)
	fmt.Println(cart, "length :", len(cart), "capacity :", cap(cart))

	remFirst()
	fmt.Println("updated cart :", cart)

}
