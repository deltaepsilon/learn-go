package squares_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSquares(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SquaresInRectangle Suite")
}
