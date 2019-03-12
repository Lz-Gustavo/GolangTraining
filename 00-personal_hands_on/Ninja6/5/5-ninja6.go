package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float32
}

type square struct {
	side int
}

func (s square) area() float32 {
	return float32(s.side * s.side)
}

type circle struct {
	radius int
}

func (c circle) area() float32 {
	return math.Pi * float32(c.radius) * 2
}

func info(s shape) {

	area := s.area()
	fmt.Println("the shape area is:", area)
}

func main() {

	sq := square{
		side: 10,
	}

	cr := circle{
		radius: 2,
	}

	info(sq)
	info(cr)
}
