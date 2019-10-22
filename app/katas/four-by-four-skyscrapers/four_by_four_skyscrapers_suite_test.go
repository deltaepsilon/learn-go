package skyscrapers_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFourByFourSkyscrapers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FourByFourSkyscrapers Suite")
}
