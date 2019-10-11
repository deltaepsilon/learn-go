package functional_addition_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFunctionalAddition(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FunctionalAddition Suite")
}
