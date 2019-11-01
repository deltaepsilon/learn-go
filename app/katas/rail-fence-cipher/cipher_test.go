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

	It("Should handle the ABCs", func() {
		testAbc(2)
		testAbc(3)
		testAbc(4)
		testAbc(5)
		testAbc(6)
		testAbc(7)
		testAbc(8)
		testAbc(9)
		testAbc(10)
		testAbc(11)
		testAbc(12)
	})
})

func testAbc(rows int) {
	abc := "abcdefghijklmnopqrstuvwxyz"
	encoded := Encode(abc, rows)
	decoded := Decode(encoded, rows)

	Expect(decoded).To(Equal(abc))
}
