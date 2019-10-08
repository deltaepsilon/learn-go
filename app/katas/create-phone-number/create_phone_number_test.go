package create_phone_number_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/create-phone-number"
)

var _ = Describe("CreatePhoneNumber", func() {
	It("basic test", func() {
    Expect(CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})).To(Equal("(123) 456-7890"))
  })
})
