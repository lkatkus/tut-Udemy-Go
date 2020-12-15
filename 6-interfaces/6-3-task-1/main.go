package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("Your shape area:", s.getArea())
}

func main() {
	myTriangle := triangle{height: 3, base: 10}
	mySquare := square{sideLength: 6}

	printArea(myTriangle)
	printArea(mySquare)
}
