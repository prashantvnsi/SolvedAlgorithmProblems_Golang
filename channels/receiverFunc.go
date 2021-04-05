package main

import (
	"fmt"
)

type square struct {
	side float32
}

type circle struct {
	radius float32
}

type shape interface {
	area() float32
}

func main() {
	s1 := square{
		side: 5,
	}

	c1 := circle{
		radius: 5,
	}

	fmt.Println("The area of Square is", s1.area(), "and the area of circle is", c1.area())
	info(s1)
	info(c1)
}

func (s square) area() float32 {
	a := s.side * s.side
	return a
}

func (c circle) area() float32 {
	a := 3.14 * c.radius * c.radius
	return a
}

func info(s shape) {
	fmt.Println("The area is", s.area())
}
