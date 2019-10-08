package digital_root_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/digital-root"
)

var _ = Describe("Kata", func() {
	It("fixed tests", func() {
		Expect(DigitalRoot(16)).To(Equal(7))
		Expect(DigitalRoot(195)).To(Equal(6))
		Expect(DigitalRoot(992)).To(Equal(2))
		Expect(DigitalRoot(167346)).To(Equal(9))
		Expect(DigitalRoot(0)).To(Equal(0))
	})
})
