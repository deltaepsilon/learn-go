package terminal_game_move_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTerminalGameMove(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TerminalGameMove Suite")
}
