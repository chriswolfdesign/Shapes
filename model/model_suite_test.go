package model_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestShapes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shapes Suite")
}
