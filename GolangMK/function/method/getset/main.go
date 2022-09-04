package main

import "fmt"

type SomeStruct struct {
	Public  string
	private string
}

type result struct {
	Public  string
	private int
}

func NewStruct(Public string, private int) result {
	return result{Public, private}
}

func main() {

	fmt.Println(NewStruct("Big", 3))

}
