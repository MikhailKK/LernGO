package main

import "fmt"

// func zero(x *int) {
// 	*x = 1
// }

// func main() {
// 	x := 5
// 	zero(&x)
// 	fmt.Println(x)
// }

func square(x *float64) {
	*x = *x * *x
}
func main() {
	x := 1.5
	square(&x)
	fmt.Println(x)
}
