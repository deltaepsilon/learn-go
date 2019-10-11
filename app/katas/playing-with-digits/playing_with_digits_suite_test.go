package digits_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPlayingWithDigits(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PlayingWithDigits Suite")
}
