package skiplist

import (
	"step/misc/randtools"
	"testing"
)

func Test_newSkiplist(t *testing.T) {
	sk := newSkiplist()

	for i := 0; i < 5; i++ {
		num := randtools.Range(1, 1000)
		n := sk.insert(num, i)
		t.Log(n)
	}
}
