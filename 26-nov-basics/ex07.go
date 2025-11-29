package main

import "fmt"

// **Problem 7: Health Potion (Pointers)** **Concept:** Pointer Receivers vs Value Receivers

// * Define a struct `Player` with fields `Name` (string) and `Health` (int).
// * Create a method `DrinkPotion()` associated with the `Player` struct.
// * **Requirement:** The method must receive a **pointer** to the player (`*Player`) so it can modify the actual health value. Add 20 to health (cap it at 100).
// * Prove it works by printing the health *before* and *after* calling the method in `main`.

type Player struct {
	Name   string
	Health int
}

// Add 20 to health (cap at 100).
func (p *Player) DrinkPotion() {
	p.Health += 20
	if p.Health > 100 {
		p.Health = 100
	}
}

func main() {
	//health before
	p := Player{Name: "virat", Health: 40}
	fmt.Println("health before potion : ", p.Health)
	p.DrinkPotion()
	fmt.Println("health after potion : ", p.Health)
}
