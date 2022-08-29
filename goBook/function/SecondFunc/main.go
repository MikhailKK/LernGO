package main

import "fmt"

func main() {
	x := 0
	func1 := func() int {
		x += 1
		return x
	}

	fmt.Println(func1())
	fmt.Println(func1())

}
