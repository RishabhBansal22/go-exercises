# **GO PROGRAMMING ASSIGNMENT**

**Topic:** Fundamentals, Data Structures, and Interfaces **Difficulty:** Beginner to Intermediate **Target Audience:** Future Systems Architects

---

### **INSTRUCTIONS**

1. **Imports:** You are allowed to use **only** the `fmt` package. Do not import `math`, `strings`, or external libraries.  
2. **Goal:** Focus on syntax, memory management, and control flow.  
3. **Output:** Ensure every program prints the expected result to the console to prove logic validity.

---

### **PART 1: CONTROL FLOW & COLLECTIONS**

**Problem 1: The "Smart" Countdown** **Concept:** Loops, Conditionals Write a program that simulates a rocket launch sequence.

* Initialize a loop counting down from **10 to 1**.  
* For every number, print the number.  
* **Condition A:** If the number is even, print "System Check OK" next to the number.  
* **Condition B:** If the number is **3**, print "Ignition Sequence Start".  
* After the loop finishes (at 0), print "Liftoff\!".

**Problem 2: Dynamic Inventory (Slices)** **Concept:** Slices, Append, Slicing

* Create an empty slice of strings named `cart`.  
* Append three items: "Laptop", "Mouse", "Keyboard".  
* Print the cart contents, its length, and its capacity.  
* **Action:** Remove the first item ("Laptop") using slice syntax (`cart[low:high]`) so only the last two items remain.  
* Print the updated cart.

**Problem 3: The Library Index (Maps)** **Concept:** Maps, "Comma ok" idiom

* Create a map where the **Key** is an ISBN (string) and the **Value** is a Book Title (string).  
* Add 3 entries to the map.  
* **Query 1:** Print a book title using a key that exists.  
* **Query 2:** Try to access a key that *does not* exist. Use the `val, ok :=` syntax to check if the key is missing. If missing, print "Book not found" (do not let the program print an empty string).

**Problem 4: The Election Counter (Algorithmic Logic)** **Concept:** Iteration, Map Logic

* Given the slice: `votes := []string{"Go", "Python", "Go", "Java", "Go", "Python"}`.  
* Create an empty map to store the results.  
* Iterate through the slice and count the frequency of each language.  
* Print the final tally (e.g., "Go: 3, Python: 2...").

---

### **PART 2: FUNCTIONS & POINTERS**

**Problem 5: The ID Generator (Closures)** **Concept:** Functions as Values, State

* Write a function named `idGenerator`. It should take no arguments but return a `func() int`.  
* The returned function must increment a counter every time it is called, preserving its state between calls.  
* In `main`, create a generator instance and call it 5 times, printing the returned ID each time.

**Problem 6: The Safe Divider (Error Handling)** **Concept:** Multiple Returns, Error Interface

* Write a function `Divide(a, b float64) (float64, error)`.  
* **Logic:** If `b` is 0, return `0` and an error using `fmt.Errorf("cannot divide by zero")`. Otherwise, return the result and `nil`.  
* In `main`, call the function twice: once with valid numbers, and once with a divisor of 0\.  
* Use `if err != nil` to handle the error gracefully in the console.

**Problem 7: Health Potion (Pointers)** **Concept:** Pointer Receivers vs Value Receivers

* Define a struct `Player` with fields `Name` (string) and `Health` (int).  
* Create a method `DrinkPotion()` associated with the `Player` struct.  
* **Requirement:** The method must receive a **pointer** to the player (`*Player`) so it can modify the actual health value. Add 20 to health (cap it at 100).  
* Prove it works by printing the health *before* and *after* calling the method in `main`.

---

### **PART 3: OBJECT ORIENTED PATTERNS**

**Problem 8: The Admin User (Struct Embedding)** **Concept:** Composition (Inheritance-like behavior)

* Define a struct `User` with fields `Name` and `Email`.  
* Define a struct `Admin`. Embed the `User` struct inside it (anonymous field).  
* Add a specific field `Level` (string) to `Admin`.  
* Create an instance of `Admin`. Print the name by accessing `admin.Name` directly (demonstrating field promotion).

**Problem 9: RPG Character Stats (Methods)** **Concept:** Value Receivers

* Using the `Player` struct from Problem 7\.  
* Create a method `PrintStats()` that uses a **value receiver** (not a pointer).  
* The method should print a formatted string: "Player \[Name\] has \[Health\]% health remaining."

**Problem 10: The Payment Gateway (Interfaces)** **Concept:** Polymorphism

* Define an interface `PaymentMethod` with a single method `Pay(amount int)`.  
* Define two structs: `CreditCard` and `Bitcoin`.  
* Implement the `Pay` method for both (e.g., CreditCard prints "Paid $X via Card", Bitcoin prints "Paid $X via BTC").  
* Write a function `Process(p PaymentMethod, amount int)` that takes the interface as an argument.  
* In `main`, pass both a `CreditCard` and a `Bitcoin` instance to `Process`.

---

**End of Assignment**

