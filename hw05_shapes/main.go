package main

import (
	"errors"
	"fmt"

	"github.com/sar0868/sar0868/otus_go_basic_hw/hw05_shapes/models/circle"
	"github.com/sar0868/sar0868/otus_go_basic_hw/hw05_shapes/models/rectangle"
	"github.com/sar0868/sar0868/otus_go_basic_hw/hw05_shapes/models/triangle"
	"github.com/sar0868/sar0868/otus_go_basic_hw/hw05_shapes/shape"
)

func calculateArea(s any) (float64, error) {
	if shape, ok := s.(shape.Shape); ok {
		return shape.Area(), nil
	}
	err := errors.New("don't Shape")
	return 0, err
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

	areaCircle, err := calculateArea(circle)
	if err != nil {
		fmt.Errorf("circle calculateArea: %w", err)
	}
	fmt.Println(areaCircle)
	areaRect, err := calculateArea(rectangle)
	if err != nil {
		fmt.Errorf("rectangle calculateArea: %w", err)
	}
	fmt.Println(areaRect)
	areaTriangle, err := calculateArea(triangle)
	if err != nil {
		fmt.Errorf("triangle calculateArea: %w", err)
	}
	fmt.Println(areaTriangle)
	areaSquare, err := calculateArea(square)
	if err != nil {
		fmt.Errorf("square calculateArea: %w", err)
	}
	fmt.Println(areaSquare) 
}
