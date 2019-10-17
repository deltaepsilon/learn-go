package partitions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/integer-partitions"
)

var _ = Describe("Partitions", func() {
	It("1", func() {
		Expect(Partitions(1)).To(Equal(1))
	})
	
	It("2", func() {
		Expect(Partitions(2)).To(Equal(2))
	})
	
	It("3", func() {
		Expect(Partitions(3)).To(Equal(3))
	})
	
	It("4", func() {
		Expect(Partitions(4)).To(Equal(5))
	})
	
	It("5", func() {
		Expect(Partitions(5)).To(Equal(7))
	})
	
	It("6", func() {
		Expect(Partitions(6)).To(Equal(11))
	})
	
	It("7", func() {
		Expect(Partitions(7)).To(Equal(15))
	})
	
	It("8", func() {
		Expect(Partitions(8)).To(Equal(22))
	})
	
	It("9", func() {
		Expect(Partitions(9)).To(Equal(30))
	})

	It("10", func() {
		Expect(Partitions(10)).To(Equal(42))
	})

	It("25", func() {
		Expect(Partitions(25)).To(Equal(1958))
	})
})
