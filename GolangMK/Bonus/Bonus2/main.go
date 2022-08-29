package main

import "fmt"

func main() {
	var arr1 []uint
	// var arr2 []int

	// var eOA int = 0
	for i := 1; i < 21; i++ {
		arr1 = append(arr1, uint(i))

		if i == 10 {
			// break
		}
	}
	fmt.Println(arr1)
}
