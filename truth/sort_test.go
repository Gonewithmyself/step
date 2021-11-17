package truth

import (
	"sort"
	"testing"
)

func BenchmarkSort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort1()
	}
}

func BenchmarkSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort2()
	}
}

func TestSearch(t *testing.T) {
	data := []int{1, 1, 2, 2, 2, 4, 5, 5}

	n := len(data)
	sort.Ints(data)
	t.Log(data, n)

	x := 3
	idx := sort.Search(n, func(i int) bool {
		return data[i] >= x
	})
	t.Log(idx, data[idx])

	x = 2
	idx = sort.Search(n, func(i int) bool {
		return data[i] >= x
	})
	t.Log(idx, data[idx])

	x = 7
	idx = sort.Search(n, func(i int) bool {
		return data[i] >= x
	})
	t.Log(idx, data[idx])
}
