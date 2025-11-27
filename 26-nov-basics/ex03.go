package main

import "fmt"

var v = make(map[string]int)

var n = "rishabh"
var r *string = &n

var book = map[string]map[string]int{
	"python": {
		"first": 1,
	},
	"c": {
		"first": 0,
	},
}

func main() {
	v = map[string]int{
		"python": 1,
		"c":      2,
	}
	lan, ok := v["java"]
	if ok {
		fmt.Println(lan)
	} else {
		fmt.Println("key does not exists", ok)
	}
	fmt.Println(book["python"]["first"])

}
