package main

import (
	"fmt"
	"log"
)

func averageArr(anyArr ...int) float64 {
	total := 0.0
	for _, value := range anyArr {
		total += float64(value)
	}
	log.Print("count is\t", total)
	return total / float64(len(anyArr))

}

func main() {
	arr := []int{2, 3, 5, 5, 3, 6, 7}
	fmt.Println(averageArr(arr...))
}
