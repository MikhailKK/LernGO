package main

import (
	"fmt"
	"math"
)

func main() {
	//  уровень масштабирования
	// возведение в степень
	z := 14.0
	p1 := math.Pow(2, z+8) / 2

	// данные для "широта"
	// l1 := 37.660557
	fmt.Println("Введите долготу и широту через запятую")
	var l1 float64
	var l2 float64
	fmt.Scanf("%f, %f", &l1, &l2)

	// данные для "долгота"
	// l2 := 55.691734

	// функция расчета широты
	long := func(p1, l1 float64) float64 {

		return p1 * (1.0 + l1/180.0)

	}

	// данные для злаебучей математики
	p2 := math.Pow(2, z+8) / 2

	// e := 0.0
	e := 0.0818191908426
	B := (l2 * math.Pi) / 180
	F := (1 - e*math.Sin(B)) / (1 + e*math.Sin(B))
	T := math.Tan((math.Pi/4)+(B/2)) * math.Pow(F, e/2)

	// расчет долготы
	lut := func(p2, T float64) float64 {
		return p2 * (1 - math.Log(T)/math.Pi)
	}

	x := long(p1, l1)
	y := lut(p2, T)
	fmt.Println("x =", x)
	fmt.Println("y = ", y)

	// округляем (переводим float --> int)
	var x2 int = int(math.Floor(x))
	var y2 int = int(math.Floor(y))

	fmt.Println("x округленный", x2)
	fmt.Println("у округленный", y2)

	// Функция для расчета номера тайла на основе глобальных пиксельных координат.
	final := func(x2, y2 int) (int, int) {

		return x2 / 256, y2 / 256
	}
	fmt.Println("номера тайлов \nсначала х потом у")
	fmt.Println(final(x2, y2))

}
