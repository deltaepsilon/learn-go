package sudoku_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSudokuValidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SudokuValidator Suite")
}
