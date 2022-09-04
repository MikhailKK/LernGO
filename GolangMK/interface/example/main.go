package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %s article written by %s", a.Title, a.Author)
}

// func Print(s fmt.Stringer) {
// 	fmt.Println(s)
// }

func main() {
	art := Article{
		Title:  "Success",
		Author: "Mike",
	}

	// Print(art)
	fmt.Println(art.String())

}
