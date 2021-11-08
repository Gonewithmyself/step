package truth

import "testing"

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
