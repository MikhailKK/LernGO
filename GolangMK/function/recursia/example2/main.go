package main

import (
	"fmt"
	"log"
)

var a int

func main() {
	fmt.Println(recursive(5))
}

func recursive(number int) int {
	a = 0
	if number > 1 {

		a = number * recursive(number-1)
		log.Println("Это лог", a)

		return a
	}

	return 1

}
