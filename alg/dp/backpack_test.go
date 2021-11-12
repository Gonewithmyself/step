package dp

import "testing"

func Test_packet_put(t *testing.T) {
	p := newPacket()
	t.Log(p.put(), p)
}
