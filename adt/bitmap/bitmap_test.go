package bitmap

import (
	"testing"
)

func Test_newBitmap(t *testing.T) {
	b := newBitmap(9)

	for i := 0; i <= b.n+5; i++ {
		b.set(i)
		t.Logf("%b %d %v-%v-%v\n", b.m, i, b.isSet(i), b.isSet(i+1), b.isSet(i-1))
		b.unSet(i)
	}

	b.set(1)
	t.Logf("%b %d\n", b.m, b.n)
}
