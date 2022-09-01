package main

func squareOfSums(a, b int, sumFunc func(a, b int) int) int {
	return sumFunc(a, b) * sumFunc(a, b)
}

func main() {
	calc := func(a, b int) int {
		c := a + b
		return c
	}

	println(squareOfSums(3, 5, calc))

}
