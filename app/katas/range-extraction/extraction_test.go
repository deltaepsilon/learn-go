package extraction_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/range-extraction"
)

var _ = Describe("Extraction", func() {

	It("should match the example", func() {
		input := []int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}
		expected := "-6,-3-1,3-5,7-11,14,15,17-20"

		Expect(Solution(input)).To(Equal(expected))
	})

	It("should match the example", func() {
		input := []int{59, 60, 61, 62, 68, 69, 70, 71, 80, 81, 82, 83, 84, 85, 88, 97}
		expected := "59-62,68-71,80-85,88,97"

		Expect(Solution(input)).To(Equal(expected))
	})
})
