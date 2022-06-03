package model

type Rectangle struct {
	height int
	width int
}

func MakeRectangle(height int, width int) Rectangle {
	return Rectangle{height: height, width: width}
}

func (r Rectangle) GetPerimeter() float64 {
	return float64((r.height * 2) + (r.width * 2))
}

func (r Rectangle) GetArea() float64 {
	return float64(r.height * r.width)
}