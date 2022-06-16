package model

type Square struct {
	Rectangle
}

func MakeSquare(side int) Square {
	return Square{MakeRectangle(side, side)}
}
