package main

import "fmt"

// **Problem 9: RPG Character Stats (Methods)** **Concept:** Value Receivers

// * Using the `Player` struct from Problem 7\.
// * Create a method `PrintStats()` that uses a **value receiver** (not a pointer).
// * The method should print a formatted string: "Player \[Name\] has \[Health\]% health remaining."

// Player struct from Problem 7
type Player struct {
	Name   string
	Health int
}

func (p Player) PrintStats() {
	fmt.Printf("Player %s has %d health remaining", p.Name, p.Health)
}

func main() {
	p := Player{"rishabh", 44}
	p.PrintStats()
}
