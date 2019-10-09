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
		dotest("733049910872815764", 5, "330479108928157")
		dotest("123456987654", 6, "234561876549")
		dotest("123456987653", 6, "234561356789")
		dotest("66443875", 4, "44668753")
		dotest("66443875", 8, "64438756")
		dotest("664438769", 8, "67834466")
		dotest("123456779", 8, "23456771")
		dotest("123456779", 0, "")
		dotest("", 8, "")
		dotest("563000655734469485", 4, "0365065073456944")
	})
})

func dotest(s string, n int, exp string) {
	var ans = Revrot(s, n)

	Expect(ans).To(Equal(exp))
}
