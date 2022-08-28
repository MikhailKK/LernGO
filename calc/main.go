package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Объявлеяем переменные
	var textInPut string
	var number1 float64
	var operator string
	var number2 float64

	// Создаем цикл для переменной ввода по которой будем проходиться три раза (первый операд, оператор, второй операнд)
	for i := 0; i < 3; i++ {
		if i != 1 {
			fmt.Println("entry number")
		} else {
			fmt.Println("entry operanor +, -, *, /")
		}
		_, err := fmt.Scanln(&textInPut)

		// тут распарсиваем текст в вещесмтвенные значения, оператор оставляем текстовым, проставляем панику для ошибок (для оператора пока не работает так как он текст и паника не приминяется но калькудятор не рассчитает)
		switch i {
		// операнд один
		case 0:
			number1, err = strconv.ParseFloat(textInPut, 64)
			if err != nil {
				panic(err)

			}
			// оператор
		case 1:
			operator = textInPut
			if err != nil {
				panic(err)
			}
			// операнд два
		case 2:
			number2, err = strconv.ParseFloat(textInPut, 64)
			if err != nil {
				panic(err)
			}
			// тут проверяем оператора и выполняем действие +, -, *, /
			switch operator {
			case "+":
				fmt.Println("answer is ", number1+number2)
			case "-":
				fmt.Println("answer is ", number1-number2)
			case "*":
				fmt.Println("answer is ", number1*number2)
			case "/":
				fmt.Println("answer is ", number1/number2)
			}

		}
	}

}

// Ворой вариант на будущее спередачей функции пока не работает
/*
var A float64 = first()
var B string = operator()
var C float64 = second()

switch operator() {
case "+":
	fmt.Println("result", first()+second())
case "-":
	fmt.Println(first() + second())
case "*":
	fmt.Println(first() + second())
case "/":
	fmt.Println(first() + second())
default:
	fmt.Println("herny")
}
fmt.Println("All", A, B, C)
}

func first() float64 {
var A = 0.0
fmt.Print("entry")
_, err := fmt.Scanf("%f", &A)
if err != nil {
	fmt.Print("wrong")
}
return A
}

func operator() string {
var B = "str"
fmt.Print("operator")
_, err := fmt.Scanf("%s", &B)
if err != nil {
	fmt.Print("wrong")
}
return B
}

func second() float64 {
var C = 0.0
fmt.Print("second")
_, err := fmt.Scanf("%f", &C)
if err != nil {
	fmt.Print("wrong")
}
return C*/
