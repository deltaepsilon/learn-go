package squares_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/squares-in-rectangle"
)

var _ = Describe("SquaresInRectangle", func() {
	It("should handle basic cases", func() {
		dotest(5, 3, []int{3, 2, 1, 1})
		dotest(3, 5, []int{3, 2, 1, 1})
		dotest(5, 5, nil)        
		dotest(20, 14, []int{14, 6, 6, 2, 2, 2})        
	}) 
})

func dotest(lng int, wdth int, exp []int) {
	var ans = SquaresInRect(lng, wdth)
	Expect(ans).To(Equal(exp))
}
