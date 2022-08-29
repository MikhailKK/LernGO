package main

import "fmt"

func main() {
	fmt.Println(max(7, 2, 3, 4, 5))
}
func max(args ...int) int {
	var maximum int
	for _, a := range args {
		if maximum < a {
			maximum = a
		}
	}
	return maximum
}
