package twice_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTwiceLinear(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TwiceLinear Suite")
}
