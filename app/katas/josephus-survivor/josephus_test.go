package josephus_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/josephus-survivor"
)

var _ = Describe("Josephus", func() {
	It("should handle basic tests", func() {
		dotest(7, 3, 4)
		// dotest(11, 19, 10)
		// dotest(40, 3, 28)
		// dotest(14, 2, 13)
		// dotest(100, 1, 100)
	})
})

func dotest(n, k, e int) {
	Expect(JosephusSurvivor(n, k)).To(Equal(e))
}
