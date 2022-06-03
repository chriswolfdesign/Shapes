package model

import "fmt"

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

func (r Rectangle) DisplayStats() {
	fmt.Println("Perimeter:", r.GetPerimeter())
	fmt.Println("Area:", r.GetArea())
	fmt.Println()  // for spacing
}