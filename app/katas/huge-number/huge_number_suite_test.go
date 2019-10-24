package huge_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHugeNumber(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HugeNumber Suite")
}
