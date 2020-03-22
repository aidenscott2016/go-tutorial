package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return math.Pow(s.side, 2)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printShape(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{10, 5}
	s := square{10}
	printShape(t)
	printShape(s)
}
