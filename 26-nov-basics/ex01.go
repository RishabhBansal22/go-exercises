package main

import "fmt"

func main() {
	//countdown loop
	for i := 10; i > 0; i-- {
		switch {
		case i%2 == 0:
			fmt.Println(i, " system check ok")
		case i == 3:
			fmt.Println(i, " Ignition Sequence Start")
		default:
			fmt.Println(i)
		}
	}
	fmt.Println("Liftoff!")

}
