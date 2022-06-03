package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Foo", func() {
	Describe("Tests", func() {
		It("Should break", func() {
			Expect(true).To(Equal(true))
		})
	})
})
