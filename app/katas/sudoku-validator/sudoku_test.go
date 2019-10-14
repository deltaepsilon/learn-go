package sudoku_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/sudoku-validator"
)

var _ = Describe("Sudoku", func() {
	It("valid", func() {
		sudoku := [][]int{
			{5, 3, 4, 6, 7, 8, 9, 1, 2},
			{6, 7, 2, 1, 9, 5, 3, 4, 8},
			{1, 9, 8, 3, 4, 2, 5, 6, 7},
			{8, 5, 9, 7, 6, 1, 4, 2, 3},
			{4, 2, 6, 8, 5, 3, 7, 9, 1},
			{7, 1, 3, 9, 2, 4, 8, 5, 6},
			{9, 6, 1, 5, 3, 7, 2, 8, 4},
			{2, 8, 7, 4, 1, 9, 6, 3, 5},
			{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}

		Expect(ValidateSolution(sudoku)).To(Equal(true))
	})


	It("should validate a valid sudoku", func() {
		sudoku := [][]int{
			{8, 2, 6, 3, 4, 7, 5, 9, 1},
			{7, 3, 5, 8, 1, 9, 6, 4, 2},
			{1, 9, 4, 2, 6, 5, 8, 7, 3},
			{3, 1, 7, 5, 8, 4, 2, 6, 9},
			{6, 5, 9, 1, 7, 2, 4, 3, 8},
			{4, 8, 2, 9, 3, 6, 7, 1, 5},
			{9, 4, 8, 7, 5, 1, 3, 2, 6},
			{5, 6, 1, 4, 2, 3, 9, 8, 7},
			{2, 7, 3, 6, 9, 8, 1, 5, 4},
		}

		Expect(ValidateSolution(sudoku)).To(Equal(true))
	})

	It("valid rows but invalid grids", func() {
		sudoku := [][]int{
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{2, 3, 1, 5, 6, 4, 8, 9, 7},
			{3, 1, 2, 6, 4, 5, 9, 7, 8},
			{4, 5, 6, 7, 8, 9, 1, 2, 3},
			{5, 6, 4, 8, 9, 7, 2, 3, 1},
			{6, 4, 5, 9, 7, 8, 3, 1, 2},
			{7, 8, 9, 1, 2, 3, 4, 5, 6},
			{8, 9, 7, 2, 3, 1, 5, 6, 4},
			{9, 7, 8, 3, 1, 2, 6, 4, 5},
		}

		Expect(ValidateSolution(sudoku)).To(Equal(false))
	})

	It("valid grids, invalid row", func() {
		sudoku := [][]int{
			{5, 3, 4, 6, 7, 8, 9, 1, 2},
			{6, 7, 2, 1, 9, 5, 3, 4, 8},
			{1, 9, 8, 3, 4, 2, 5, 6, 7},
			{8, 5, 9, 7, 6, 1, 4, 2, 3},
			{4, 2, 6, 8, 5, 3, 7, 9, 1},
			{7, 1, 3, 9, 2, 4, 8, 5, 6},
			{9, 6, 1, 5, 3, 7, 2, 8, 4},
			{2, 8, 7, 4, 1, 9, 6, 3, 5},
			{3, 0, 0, 4, 8, 1, 1, 7, 9},
		}

		Expect(ValidateSolution(sudoku)).To(Equal(false))
	})
})
