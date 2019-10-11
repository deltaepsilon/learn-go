package secure_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/not-very-secure"
)

var _ = Describe("Secure", func() {
	It("Should return the correct values for the sample test cases!", func() {
		Expect(Alphanumeric(".*?")).To(Equal(false))
		Expect(Alphanumeric("a")).To(Equal(true))
		Expect(Alphanumeric("Mazinkaiser")).To(Equal(true))
		Expect(Alphanumeric("hello world_")).To(Equal(false))
		Expect(Alphanumeric("PassW0rd")).To(Equal(true))
		Expect(Alphanumeric("     ")).To(Equal(false))
		Expect(Alphanumeric("")).To(Equal(false))
		Expect(Alphanumeric("\n\t\n")).To(Equal(false))
		Expect(Alphanumeric("ciao\n$$_")).To(Equal(false))
		Expect(Alphanumeric("__ * __")).To(Equal(false))
		Expect(Alphanumeric("&)))(((")).To(Equal(false))
		Expect(Alphanumeric("43534h56jmTHHF3k")).To(Equal(true))
	})
})
