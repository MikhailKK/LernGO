package main

import "fmt"

func main() {
	arr1 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	arr2 := make([]int, 3)
	copy(arr2, arr1[1:3])

	fmt.Println(arr2)
	fmt.Println(cap(arr2), len(arr2))

	arr2 = append(arr2, 11)
	fmt.Println(arr2)
	fmt.Println(cap(arr2), len(arr2))
}
