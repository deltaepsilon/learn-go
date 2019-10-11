package digits_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/playing-with-digits"
)

var _ = Describe("Digits", func() {
	It("should handle basic cases", func() {
		dotest(89, 1, 1)
		dotest(92, 1, -1)
		dotest(695, 2, 2)
		dotest(46288, 3, 51)
	})
})

func dotest(n, p int, exp int) {
	var ans = DigPow(n, p)
	Expect(ans).To(Equal(exp))
}
