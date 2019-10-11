package josephus_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJosephusSurvivor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JosephusSurvivor Suite")
}
