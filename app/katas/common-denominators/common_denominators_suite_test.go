package denominators_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCommonDenominators(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CommonDenominators Suite")
}
