package main

import "fmt"
func main (){
	arr1 := []int{0,1,2,3,4}
	arr2 := make([]int, 2)
	copy (arr2, arr1[1:3])

	fmt.Println (arr2)


}