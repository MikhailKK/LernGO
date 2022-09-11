package main

import (
	"fmt"
	"lerngo"
)

func main() {
	fmt.Println("Phone list")

	var AP lerngo.AndroidPhone
	AP.Model = "Xiamomi"

	fmt.Println(AP.GetModel())

}
