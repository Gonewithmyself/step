package bitmap

type bitmap struct {
	m []uint8
	n int
}

func newBitmap(n int) *bitmap {
	slot := n/8 + 1
	return &bitmap{
		m: make([]uint8, slot),
		n: slot * 8,
	}
}

func (b *bitmap) set(x int) {
	if x <= 0 || x > b.n {
		return
	}

	x--
	slot := x / 8

	bit := x % 8
	bit = 1 << bit
	b.m[slot] |= uint8(bit)
}

func (b *bitmap) unSet(x int) {
	if x <= 0 || x > b.n {
		return
	}

	x--
	slot := x / 8

	bit := x % 8
	bit = 1 << bit
	b.m[slot] &= uint8(^bit)
}

func (b *bitmap) isSet(x int) bool {
	if x <= 0 || x > b.n {
		return false
	}

	x--
	slot := x / 8

	bit := x % 8
	bit = 1 << bit
	return b.m[slot]&uint8(bit) != 0
}
