package main

import (
	"fmt"
	"log"
)

// N-ый элемент арифметической прогрессии. Прямая рекурсия.
func ArithmeticProgressionElement(a1, d, n int) int {
	if n == 1 {
		log.Println("Лог ", n)
		return a1
	}
	return d + ArithmeticProgressionElement(a1, d, n-1)
}

func main() {
	fmt.Println(ArithmeticProgressionElement(2, 3, 4))

}

// N-ый элемент геометрической прогрессии. Косвенная рекурсия.
/*func GeometricProgressionElement(b1, q, n int) int{
	if n == 1 {
		return b1
	}
	return q * ProgressionElement(b1, q, n)
}

func ProgressionElement(b1, q, n int) int {
	return GeometricProgressionElement(b1, q, n-1)
}
*/
