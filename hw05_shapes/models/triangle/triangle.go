package triangle

type Triangle struct {
	Base   int
	Height int
}

func (t Triangle) Area() float64 {
	return float64(t.Base) * float64(t.Height) / 2
}

func NewTriangle(base int, height int) *Triangle {
	triangle := Triangle{}
	triangle.Base = base
	triangle.Height = height
	return &triangle
}
