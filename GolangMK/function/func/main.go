package main

import (
	"fmt"
)

var logic bool

func main() {

	fmt.Println(divide(8, 2))
}

func divide(a, b int) (float64, bool) {
	if a%b == 1 {
		logic = false
		fmt.Println("нечетное число ", float64(a), float64(b))
		return float64((a / b)), logic
	}
	logic = true
	println("четное")
	return float64((a / b)), logic

}
