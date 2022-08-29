package main

import "fmt"

func main() {

	var num uint = 5
	fmt.Println(factorial(num))

}
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
