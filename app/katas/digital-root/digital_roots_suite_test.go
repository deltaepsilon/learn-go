package digital_root_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDigitalRoot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DigitalRoot Suite")
}
