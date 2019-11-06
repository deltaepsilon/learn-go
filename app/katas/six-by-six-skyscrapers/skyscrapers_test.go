package skyscrapers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/six-by-six-skyscrapers"
)

var _ = Describe("Skyscrapers", func() {
	// It("can solve 6x6 puzzle 1", func() {
	// 	var clues = []int{
	// 		3, 2, 2, 3, 2, 1,
	// 		1, 2, 3, 3, 2, 2,
	// 		5, 1, 2, 2, 4, 3,
	// 		3, 2, 1, 2, 2, 4,
	// 	}

	// 	var expected = [][]int{
	// 		{2, 1, 4, 3, 5, 6},
	// 		{1, 6, 3, 2, 4, 5},
	// 		{4, 3, 6, 5, 1, 2},
	// 		{6, 5, 2, 1, 3, 4},
	// 		{5, 4, 1, 6, 2, 3},
	// 		{3, 2, 5, 4, 6, 1},
	// 	}

	// 	Expect(SolvePuzzle(clues)).To(Equal(expected))
	// })

	It("can solve 6x6 puzzle 2", func() {
		var clues = []int{
			0, 0, 0, 2, 2, 0,
			0, 0, 0, 6, 3, 0,
			0, 4, 0, 0, 0, 0,
			4, 4, 0, 3, 0, 0,
		}
		var expected = [][]int{
			{5, 6, 1, 4, 3, 2},
			{4, 1, 3, 2, 6, 5},
			{2, 3, 6, 1, 5, 4},
			{6, 5, 4, 3, 2, 1},
			{1, 2, 5, 6, 4, 3},
			{3, 4, 2, 5, 1, 6},
		}

		Expect(SolvePuzzle(clues)).To(Equal(expected))
	})

	// It("can solve 6x6 puzzle 3", func() {
	// 	var clues = []int{
	// 		0, 3, 0, 5, 3, 4,
	// 		0, 0, 0, 0, 0, 1,
	// 		0, 3, 0, 3, 2, 3,
	// 		3, 2, 0, 3, 1, 0,
	// 	}
	// 	var expected = [][]int{
	// 		{5, 2, 6, 1, 4, 3},
	// 		{6, 4, 3, 2, 5, 1},
	// 		{3, 1, 5, 4, 6, 2},
	// 		{2, 6, 1, 5, 3, 4},
	// 		{4, 3, 2, 6, 1, 5},
	// 		{1, 5, 4, 3, 2, 6},
	// 	}

	// 	Expect(SolvePuzzle(clues)).To(Equal(expected))
	// })

})
