package denominators_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/common-denominators"
)

var _ = Describe("Denominators", func() {
	It("Basic tests", func() {
		dotest([][]int{{1, 2}, {1, 3}, {1, 4}}, "(6,12)(4,12)(3,12)")

		dotest([][]int{{69, 130}, {87, 1310}, {30, 40}}, "(18078,34060)(2262,34060)(25545,34060)")
	})
})

func dotest(a [][]int, exp string) {
	Expect(ConvertFracts(a)).To(Equal(exp))
}
