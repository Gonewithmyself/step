package backtrack

import (
	"testing"
)

func Test_newNQueen(t *testing.T) {
	nq := newNQueen()
	t.Log(nq)
}

func Test_packet(t *testing.T) {
	nq := newPacket()
	t.Log(nq)
}
