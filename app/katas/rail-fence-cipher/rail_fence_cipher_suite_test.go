package cipher_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRailFenceCipher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RailFenceCipher Suite")
}
