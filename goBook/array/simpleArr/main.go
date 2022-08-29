package main

import "fmt"

func main() {
	var arr [5]float64
	arr[0] = 5
	arr[1] = 1
	arr[2] = 3
	arr[3] = 2
	arr[4] = 5

	var SumArr float64 = 0

	fmt.Println("Входные данные", arr)

	for _, value := range arr {
		SumArr += value

	}
	fmt.Println("Общая сумма ", SumArr)
	fmt.Println("Итог ", SumArr/float64(len(arr)))
}
