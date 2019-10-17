package partitions_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIntegerPartitions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IntegerPartitions Suite")
}
