package main

/*
Разработать программу нахождения расстояния между двумя точками, которые представлены
в виде структуры Point инкапсулированными параметрами x,y и конструктором.
*/

import (
	"fmt"
	"math"
)

//структура для точки
type Point struct {
	x, y float64
}

//конструктор новой точки х и у
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// считает дистанцию от точки х до точки у по формуле sqrt(x2-x1)^2+(y2-y1)^2
func Distance(p1, p2 *Point) float64 {

	dp1 := math.Pow(p2.x-p1.x, 2)
	dp2 := math.Pow(p2.y-p1.y, 2)

	return math.Sqrt(dp1 + dp2)
}

func main() {

	p1 := NewPoint(1, 5)
	p2 := NewPoint(2, 7)
	fmt.Println("Дистанция Х и У", Distance(p1, p2))
}
