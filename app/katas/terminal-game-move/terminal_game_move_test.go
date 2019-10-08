package terminal_game_move_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deltaepsilon/learn-go/katas/terminal-game-move"
)

var _ = Describe("TerminalGameMove", func() {
	It("Move(0, 4)", func() { Expect(Move(0, 4)).To(Equal(8)) })
	It("Move(3, 6)", func() { Expect(Move(3, 6)).To(Equal(15)) })
	It("Move(2, 5)", func() { Expect(Move(2, 5)).To(Equal(12)) })
})
