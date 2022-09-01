package main

import "fmt"

func main() {

	type Circle struct {
		r      float64 // radius
		x0, y0 float64 // center
	}

	var c1 Circle  // создаст экземпляр со полями по умолчанию
	c2 := Circle{} // аналогично первому варианту
	c3 := Circle{
		r:  1,
		x0: 2,
		y0: 3,
	} // создаст экземпляр и инициализирует соотвествующие поля

	c4 := Circle{1, 2, 3}             // аналогично варианту с3 , поля в назначеном порядке структуры
	c5 := new(Circle)                 // создаст экземпляр и указатель на него
	c6 := &Circle{r: 1, x0: 2, y0: 3} // аналогично предыдущему, но с инициализацией полей
	// с7 := &Circle{2, 3, 4}
	// var c8 = &Circle{}

	// доступ к полям структуры
	// 	var c Circle
	// c.r = 10
	// println(c.x0)

	fmt.Println(c1, c2, c3, c4, c5, c6)

}
