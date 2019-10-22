package skyscrapers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/four-by-four-skyscrapers"
)

var _ = Describe("Skyscrapers", func() {
	It("should solve static puzzel 1", func() {
		var clues = []int{
			0, 0, 1, 2,
			0, 2, 0, 0,
			0, 3, 0, 0,
			0, 1, 0, 0}

		var outcome = [][]int{
			{2, 1, 4, 3},
			{3, 4, 1, 2},
			{4, 2, 3, 1},
			{1, 3, 2, 4},
		}

		Expect(SolvePuzzle(clues)).To(Equal(outcome))
	})

	It("should solve static puzzel 2", func() {
		var clues = []int{
			2, 2, 1, 3,
			2, 2, 3, 1,
			1, 2, 2, 3,
			3, 2, 1, 3,
		}

		var outcome = [][]int{
			{1, 3, 4, 2},
			{4, 2, 1, 3},
			{3, 4, 2, 1},
			{2, 1, 3, 4},
		}

		Expect(SolvePuzzle(clues)).To(Equal(outcome))
	})

	It("should solve static puzzel 3", func() {
		var clues = []int{
			0, 0, 1, 2,
			0, 2, 0, 0,
			0, 3, 0, 0,
			0, 1, 0, 0,
		}

		var outcome = [][]int{
			{2, 1, 4, 3},
			{3, 4, 1, 2},
			{4, 2, 3, 1},
			{1, 3, 2, 4},
		}

		Expect(SolvePuzzle(clues)).To(Equal(outcome))
	})

	It("should solve static puzzel 4", func() {
		var clues = []int{
			1, 2, 4, 2,
			2, 1, 3, 2,
			3, 1, 2, 3,
			3, 2, 2, 1,
		}

		var outcome = [][]int{
			{4, 2, 1, 3},
			{3, 1, 2, 4},
			{1, 4, 3, 2},
			{2, 3, 4, 1},
		}

		Expect(SolvePuzzle(clues)).To(Equal(outcome))
	})
})
