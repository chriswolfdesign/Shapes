package main

import "shapes/model"

func main() {
	shapes_list := []model.Shape{}

	shapes_list = append(shapes_list, model.MakeRectangle(4, 5))

	for _, shape := range(shapes_list) {
		shape.DisplayStats()
	}
}