package main

import (
	"fmt"
	"log"
)

func directRec(a, b int) int {
	var c int
	c = a - b
	if c > 0 {
		log.Println("Это лог", c)
		directRec(c, b)
	}
	if c <= 0 {
		log.Println("Это лог №2", c)
	}
	return c
}

func main() {
	fmt.Println(directRec(12, 3))
}
