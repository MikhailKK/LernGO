package main

import "fmt"

func main() {
	fmt.Println("xyz")
	print("The program output range of numbers from 1 till 100 , \nwhen number multiplay of tree types 'fizz'\nwhen multiplay of five types 'Bizz'\n if number multiplay of bouth types 'FizzBizz")

	for Numbers := 1; Numbers <= 100; Numbers++ {
		switch {
		case Numbers%3 == 0 && Numbers%5 == 0:
			fmt.Println(Numbers, " FizzBizz")
		case Numbers%3 == 0:
			fmt.Println(Numbers, " Fizz")
		case Numbers%5 == 0:
			fmt.Println(Numbers, " Bizz")
			// Можно раскомментить и выводить все числа с кейсом Another
			// default:
			// 	fmt.Println(Numbers, " Another")
		}
	}
	fmt.Println("the end")
}
