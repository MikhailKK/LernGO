package main

import (
	"fmt"
)

func main() {
	i := 4
	defer func(n int) {
		fmt.Println("Everything is fine:", i)
	}(i)
	i = 3
	fmt.Println(i)
}

// Вывод:
// 3
// Everything is fine: 4
