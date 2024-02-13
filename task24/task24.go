package main

import (
	"fmt"
	"math"
)

/*
	=== Задача №24 ===

	Разработать программу нахождения расстояния между двумя точками,
	которые представлены в виде структуры Point с инкапсулированными
	параметрами x,y и конструктором.

*/

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	p := &Point{x: x, y: y}
	return p
}

func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(1,2)
	p2 := NewPoint(4,6)
	fmt.Println(p1.Distance(p2))
}
