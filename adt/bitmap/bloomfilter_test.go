package bitmap

import (
	"testing"
)

func Test_newBloom(t *testing.T) {
	b := newBloom(20)

	b.add(10)

	t.Log(b)
}
