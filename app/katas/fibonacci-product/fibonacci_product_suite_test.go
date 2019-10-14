package fibonacci_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFibonacciProduct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FibonacciProduct Suite")
}
