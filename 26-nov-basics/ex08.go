package main

import "fmt"

// **Problem 8: The Admin User (Struct Embedding)** **Concept:** Composition (Inheritance-like behavior)

// * Define a struct `User` with fields `Name` and `Email`.
// * Define a struct `Admin`. Embed the `User` struct inside it (anonymous field).
// * Add a specific field `Level` (string) to `Admin`.
// * Create an instance of `Admin`. Print the name by accessing `admin.Name` directly (demonstrating field promotion).

type User struct {
	Name  string
	Email string
}

type Admin struct {
	User
	Level string
}

func main() {
	u := User{"rishabh", "rishabh@mail.com"}
	a := Admin{u, "beginner"}

	fmt.Println(a.Name)
}
