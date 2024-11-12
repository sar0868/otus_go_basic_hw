package main

import (
	"fmt"

	"github.com/sar0868/sar0868/otus_go_basic_hw/hw05_shapes/models/circle"
	"github.com/sar0868/sar0868/otus_go_basic_hw/hw05_shapes/models/rectangle"
	"github.com/sar0868/sar0868/otus_go_basic_hw/hw05_shapes/models/triangle"
	"github.com/sar0868/sar0868/otus_go_basic_hw/hw05_shapes/shape"
)

func calculateArea(s shape.Shape) float64 {
	return s.Area()
}

type Square struct {
	Side int
}

func NewSquare(side int) *Square {
	s := Square{}
	s.Side = side
	return &s
}

func main() {
	circle := circle.NewCircle(5)
	rectangle := rectangle.NewRectangle(10, 5)
	triangle := triangle.NewTriangle(8, 6)
	square := NewSquare(20)
	fmt.Printf("square side = %d\n", square.Side)

	fmt.Println(calculateArea(circle))
	fmt.Println(calculateArea(rectangle))
	fmt.Println(calculateArea(triangle))
	// fmt.Println(calculateArea(square))
}
