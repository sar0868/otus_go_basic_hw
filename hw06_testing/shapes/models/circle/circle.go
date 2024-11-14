package circle

import "math"

type Circle struct {
	Radius int
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(float64(c.Radius), 2)
}

func NewCircle(radius int) *Circle {
	circle := Circle{}
	circle.Radius = radius
	return &circle
}
