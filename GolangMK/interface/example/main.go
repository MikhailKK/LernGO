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

// Для того чтобы привести интерфейс к конкретному типу,
// внутри typeCast пишем переменную, точку и в скобках указываем тот тип,
// к которому мы хотим привести article := s.(Article). Таким образом,
// переменная article уже тип Article, и мы работаем с конкретным типом
func typeCast(s fmt.Stringer) {
	article := s.(Article)
	fmt.Printf("%+v\n", article.Title)
	// fmt.Printf("%+v\n", s.Title)
}

func main() {
	art := Article{
		Title:  "Success",
		Author: "Mike",
	}

	// Print(art)
	fmt.Println(art.String())
	typeCast(art)
}
