package main

import (
	"errors"
	"fmt"
)

func main() {
	abc, err := max(2, 5, 8, 2, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(abc)
}
func max(args ...int) (int, error) {
	var maximum int
	maximum = args[0]
	if maximum == 0 {
		return 0, errors.New("это ноль")
	}
	for _, a := range args {
		if maximum < a {
			maximum = a
		}
	}
	return maximum, nil
}
