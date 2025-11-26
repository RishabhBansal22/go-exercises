package main

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
