package twice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/twice-linear"
)

var _ = Describe("Twice", func() {
	It("should handle basic cases", func() {
		dotest(10, 22)
		dotest(50, 175)
		dotest(100, 447)
		dotest(500, 3355)
		dotest(1000, 8488)
	})
})

func dotest(n int, exp int) {
	var ans = DblLinear(n)
	Expect(ans).To(Equal(exp))
}
