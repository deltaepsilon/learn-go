package cipher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/rail-fence-cipher"
)

var _ = Describe("Cipher", func() {
	It("Example #1", func() { Expect(Encode("WEAREDISCOVEREDFLEEATONCE", 3)).To(Equal("WECRLTEERDSOEEFEAOCAIVDEN")) })
	It("Example #1 with four columns", func() { Expect(Encode("WEAREDISCOVEREDFLEEATONCE!", 4)).To(Equal("WIREEEDSEEEAC!AECVDLTNROFO")) })
	It("Example #2 with four columns", func() { Expect(Decode("WIREEEDSEEEAC!AECVDLTNROFO", 4)).To(Equal("WEAREDISCOVEREDFLEEATONCE!")) })
	It("Example #2", func() { Expect(Decode("WECRLTEERDSOEEFEAOCAIVDEN", 3)).To(Equal("WEAREDISCOVEREDFLEEATONCE")) })
	It("Example #3", func() { Expect(Encode("Hello, World!", 3)).To(Equal("Hoo!el,Wrdl l")) })
})
