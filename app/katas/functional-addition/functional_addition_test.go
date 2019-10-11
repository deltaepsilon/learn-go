package functional_addition_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/functional-addition"
)

var _ = Describe("FunctionalAddition", func() {
	It("add", func() {
		Expect(Add(1)(3)).To(Equal(4))
	})
})
