package main

import "fmt"

type pet struct {
	Name        string
	Owner       string
	CountInject int
}

type cat struct {
	pet
	Age int
}

func newCat(Name string, Owner string, CountInject int, Age int) cat {
	return cat{pet{Name, Owner, 2}, Age}
}

func main() {
	fmt.Println(newCat("cot", "Masha", 4, 3))
}
