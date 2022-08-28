package main

import (
	"fmt"
	"GolangMK/maths/subtraction"
	"GolangMK/maths/addition"
)

func main() {
	var Minus uint32
	var Summa int32
	Minus = subtraction.Maths2()
	Summa = addition.Maths1()
	fmt.Println("Result of addition is ", Summa)
	fmt.Println("Result of subtruction is ", Minus)

}
