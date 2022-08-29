package main

import ("fmt"

// "sort"
)

func main() {
	array := []int{12,44,0,22,27,31,2,17}
	// вариант - подгружаем "sort", сортируем по возрастанию и выводим нулевой элемент массива 
// sort.Ints(array)
// fmt.Println(array[0])


lowestNum := array[0]
	for _, findNum := range array {
		if findNum < lowestNum {lowestNum = findNum}
		}
		
		fmt.Println(lowestNum)
}