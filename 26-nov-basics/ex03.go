package main

import "fmt"

//todo:
// 1. implement pointers
// 2. change map to {"isbn":{"price":"22","author":"dale"}}
// 3. implement interface and methods
// 4. transform into cli application

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

func DeleteBook(lib map[string]string, isbn []string) {
	for _, val := range isbn {
		if _, ok := lib[val]; ok {
			delete(lib, val)
			fmt.Printf("%s removed from library\n", val)

		} else {
			fmt.Println("book does not exist\n")
		}
	}

}

func GetBook(lib map[string]string, isbn string) string {
	if val, ok := lib[isbn]; ok {
		return val
	} else {
		return "book does not exists"
	}
}
func main() {

	res := GetBook(lib, "12a34")
	fmt.Println(res)

}
