package reverse_rotate_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/reverse-rotate"
)

var _ = Describe("Kata", func() {
	It("should handle basic cases", func() {
		dotest("1234", 0, "")
		dotest("", 0, "")
		dotest("1234", 5, "")
		var s = "733049910872815764"
		dotest(s, 5, "330479108928157")	
	})
})


func dotest(s string, n int, exp string) {
	var ans = Revrot(s, n)
	Expect(ans).To(Equal(exp))
}
