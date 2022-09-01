package main

import (
	"fmt"
)

func readInt() {
	var number int
	fmt.Println("Enter a number: ")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("Cannot read a number: %s", err)
	}
	fmt.Println("You wrote a: ", number)
}

func main() {
	readInt()
}
