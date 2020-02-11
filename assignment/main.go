package main

import "fmt"

type triangle struct{
	height float64
	base float64
}

type square struct{
	sideLength float64
}

func main() {
	t := triangle{base: 10, height: 10}
	s := square{sideLength: 10}
	printArea(t)
	printArea(s)
}

// receivers, access properties with t., s.
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}