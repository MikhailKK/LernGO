package main

import "fmt"

func oddGenerator() func() int {
	m := int(1)
	return func() (odd int) {
		odd = m
		m += 2
		return
	}

}

func main() {

	nextOdd := oddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}
