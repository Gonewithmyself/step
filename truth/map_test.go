package truth

import (
	"testing"
	"unsafe"
)

func Test_ss(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss()
		})
	}
}

func Test_overLoadFactor(t *testing.T) {
	type args struct {
		count int
		B     uint8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := overLoadFactor(tt.args.count, tt.args.B); got != tt.want {
				t.Errorf("overLoadFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getB(t *testing.T) {
	prev := 0
	for i := 0; i < 1000; i++ {
		curr := getB(i)
		if curr != uint8(prev) {
			t.Log(curr, calculateBuckets(uint8(prev)), bucketShift(uint8(prev)), i)
			prev = int(curr)
		}
	}
}

func Test_bucketShift(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(i, bucketShift(uint8(i)))
	}
	t.Log(unsafe.Sizeof(bmap{}))
}

func Test_tooManyOverflowBuckets(t *testing.T) {
	for i := 1; i < 20; i++ {
		for j := 0; j < 100; j++ {
			if tooManyOverflowBuckets(uint16(j), uint8(i)) {
				t.Log(i, j)
				break
			}
		}
	}
}
