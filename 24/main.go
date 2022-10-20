package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x, y float64
}

func newPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (a *Point) distance(b *Point) float64 {
	return math.Sqrt(math.Pow(b.x-a.x, 2) + math.Pow(b.y-a.y, 2))
}

func main() {
	a := newPoint(1, 2)
	b := newPoint(3, 4)
	dist := a.distance(b)
	fmt.Println(dist)
}
