// package main

// import "fmt"

// func main() {
//   doSomething(func(number int) { fmt.Printf("number: %d", number) })
// }

// func doSomething(anonymousFunc(func(number int)) ) {
// 	anonymousFunc(100)
//   }

// package main

// func squareOfSums(a, b int, sumFunc func(a, b int) int) int {
// 	return sumFunc(a, b) * sumFunc(a, b)
// }

// func main() {

// 	var a, b int = 3, 5
// 	sum := func(a, b int) int {
// 		return a + b
// 	}

// 	println(squareOfSums(a, b, sum)) // результат: 64
// }

package main

func squareOfSums(a, b int, sumFunc func(a, b int) int) int {
	return sumFunc(a, b) * sumFunc(a, b)
}

func main() {
	sum := func(a, b int) int {
		return a + b
	}
	println(squareOfSums(3, 5, sum)) // результат: 64
}
