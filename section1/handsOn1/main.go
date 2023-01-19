package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func info(s shape) {
	fmt.Printf("Area is : %v\n", s.area())
}

func main() {
	s := square{3}
	c := circle{3}

	info(s)
	info(c)
}
