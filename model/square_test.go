package model_test

import (
	. "shapes/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func generateSquares() (Square, Square, Square) {
	sq1 := MakeSquare(4)
	sq2 := MakeSquare(8)
	sq3 := MakeSquare(12)
	return sq1, sq2, sq3
}

var _ = Describe("Square", func() {
	var (
		sq1, sq2, sq3 Square
	)

	BeforeEach(func() {
		sq1, sq2, sq3 = generateSquares()
	})

	Describe("Perimeter tests", func() {
		It("Should be 16", func() {
			Expect(sq1.GetPerimeter()).To(Equal(16.0))
		})

		It("Should be 32", func() {
			Expect(sq2.GetPerimeter()).To(Equal(32.0))
		})

		It("Should be 48", func() {
			Expect(sq3.GetPerimeter()).To(Equal(48.0))
		})
	})

	Describe("Area tests", func() {
		It("Should be 16", func() {
			Expect(sq1.GetArea()).To(Equal(16.0))
		})

		It("Should be 64", func() {
			Expect(sq2.GetArea()).To(Equal(64.0))
		})

		It("Should be 144", func() {
			Expect(sq3.GetArea()).To(Equal(144.0))
		})
	})
})
