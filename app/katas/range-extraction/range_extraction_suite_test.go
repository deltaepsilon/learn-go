package extraction_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRangeExtraction(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RangeExtraction Suite")
}
