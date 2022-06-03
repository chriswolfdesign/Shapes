package model

type Square struct {
	rect Rectangle
}

func MakeSquare(side int) Square {
	rect := MakeRectangle(side, side)
	return Square{rect: rect}
}

func (s Square) GetPerimeter() float64 {
	return s.rect.GetPerimeter()
}

func (s Square) GetArea() float64 {
	return s.rect.GetArea()
}

func (s Square) DisplayStats() {
	s.rect.DisplayStats()
}