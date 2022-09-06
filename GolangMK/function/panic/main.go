// деление на 0, разыменование nil указателя,
// выход за границы массива
// аварийно завершает программу и вызывает panic

package main

import (
	"fmt"
)

func main() {
	CheckCardSuit("Hearts")
	CheckCardSuit("Joker")
	CheckCardSuit("Spades")
}

func CheckCardSuit(s string) {
	switch s {
	case "Spades": // ...
	case "Hearts": // ...
	case "Diamonds": // ...
	case "Clubs": // ...
	default:
		panic("Нет такой масти: " + s)
	}
	fmt.Println(s)
}

// package main

// import "fmt"

// func main() {
// 	defer func() {
// 		str := recover()
// 		fmt.Println(str)
// 	}()
// 	panic("PANIC")
// }
