package model_test

import (
	. "shapes/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func generateRectangles() (Rectangle, Rectangle, Rectangle) {
	rect1 := MakeRectangle(4, 5)
	rect2 := MakeRectangle(5, 6)
	rect3 := MakeRectangle(6, 7)
	return rect1, rect2, rect3
}

var _ = Describe("Rectangle", func() {
	var (
		rect1, rect2, rect3 Rectangle
	)

	BeforeEach(func() {
		rect1, rect2, rect3 = generateRectangles()
	})

	Describe("Perimeter tests", func() {
		It("Should be 18.0", func() {
			Expect(rect1.GetPerimeter()).To(Equal(18.0))
		})

		It("Should be 22.0", func() {
			Expect(rect2.GetPerimeter()).To(Equal(22.0))
		})

		It("Should be 26.0", func() {
			Expect(rect3.GetPerimeter()).To(Equal(26.0))
		})
	})
})
