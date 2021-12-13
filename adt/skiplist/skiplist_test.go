package skiplist

import (
	"step/misc/randtools"
	"testing"
)

func Test_newSkiplist(t *testing.T) {
	sk := newSkiplist()
	nodecount := 50

	dels := make([]int, 0, 10)
	for i := 0; i < nodecount; i++ {
		num := randtools.Range(1, 1000)
		n := sk.insert(num, i)
		t.Log("add", n, num)
		dels = append(dels, num)
	}
	t.Logf("after insert: %v\n", sk)

	for i := range dels {
		n := sk.search(dels[i])
		if n.score != dels[i] {
			t.Error("search failed")
		}
	}

	for i := range dels {
		n := sk.delete(dels[i])
		t.Log("del", dels[i], n)
	}

	t.Logf("after delete: %v\n", sk)

	// for i := range dels {
	// 	n := sk.search(dels[i])
	// 	t.Log(n, dels[i])
	// }
}
