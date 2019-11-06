package skyscrapers_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSixBySixSkyscrapers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SixBySixSkyscrapers Suite")
}
