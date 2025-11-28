package main

import "fmt"

//map of books; key = isbn, val = book title
var lib = map[string]string{
	"12a34": "verity",
	"23b45": "funk",
	"44c89": "how to win friends",
}

//function to add book in lib map
func AddBook(lib, newBooks map[string]string) {
	count := 0
	for i, val := range newBooks {
		if _, ok := lib[i]; ok {
			fmt.Println("book already exists")
		} else {
			lib[i] = val
			fmt.Printf("book %s added to library\n", i)
			count++
		}

	}
	fmt.Printf("%d books added to library\n", count)
}

func main() {
	book := map[string]string{
		"33a44": "metamorphosis",
		"23d44": "golang-class",
		"44255": "remmeber me",
	}

	AddBook(lib, book)
	// AddBook(lib, book)
	fmt.Println(lib)

}
