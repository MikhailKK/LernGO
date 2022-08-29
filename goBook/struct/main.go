package main

import (
	"fmt"
	"math"
)

type Shape interface {
	AREA() float64
}

type Perimiter struct {
	p1, p2, p3 float64
}
type Circle struct {
	x, y, r float64
}

type RectAngel struct {
	x1, y1, x2, y2 float64
}

func TOTALarea(shapes ...Shape) float64 {
	var area float64
	for _, S := range shapes {
		area += S.AREA()
	}
	return area
}
func DISTANCE(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func (r RectAngel) AREA() float64 {
	l := DISTANCE(r.x1, r.y1, r.x1, r.y2)
	w := DISTANCE(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func (c Circle) AREA() float64 {
	return math.Pi * c.r * c.r
}

func (p Perimiter) AREA() float64 {

	return p.p1 + p.p2 + p.p3
}
func main() {
	r := RectAngel{0, 0, 10, 10}
	c := Circle{0, 0, 5}
	p := Perimiter{5, 7, 8}
	
	fmt.Println(c.AREA())

	fmt.Println(TOTALarea(r))

	fmt.Println(TOTALarea(p))
}
