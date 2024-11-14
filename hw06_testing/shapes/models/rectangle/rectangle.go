package rectangle

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() float64 {
	return float64(r.Width) * float64(r.Height)
}

func NewRectangle(width int, height int) *Rectangle {
	rectangle := Rectangle{}
	rectangle.Width = width
	rectangle.Height = height
	return &rectangle
}
