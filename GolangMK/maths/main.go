package maths

import (
	"fmt"

	"golangmk/skilfactory/maths/math/skilfactory/maths/addition"
	"golangmk/skilfactory/maths/math/skilfactory/maths/subtraction"
)

func main() {
	var Minus uint32
	var Summa int32
	Summa = addition.Maths1()
	Minus = subtraction.Maths2()

	fmt.Println("Result of addition is ", Summa)
	fmt.Println("Result of subtruction is ", Minus)

}
