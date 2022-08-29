package main

import (
	"fmt"
	"math"
)

func main() {

	a := 0.33
	b := 0.2

	multiChance := func(a float64, b float64) float64 {
		return a * b

	}

	amount := func(a float64, b float64) float64 {
		return a + b
	}

	amountOf := func(args float64) float64 {
		return args + args + args
	}

	am := amountOf(amount(a, b))

	fmt.Println("сумма ", amount(a, b))
	fmt.Println("одновременно", multiChance(a, b))
	fmt.Println("сумма вероятностей for 3:", am)

	

}
