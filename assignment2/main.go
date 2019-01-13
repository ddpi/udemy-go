package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}

func main() {
	s := square{
		sideLength: 4.5,
	}

	t := triangle{
		height: 5,
		base:   9,
	}

	printArea(s)
	printArea(t)
}
