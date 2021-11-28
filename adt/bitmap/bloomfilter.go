package bitmap

import (
	"fmt"
	"step/misc/randtools"
)

type bloom struct {
	*bitmap
	seeds []int
}

func newBloom(x int) *bloom {
	seeds := make([]int, 3)
	f := [...]int{2, 3, 5}
	for i := range f {
		seeds[i] = randtools.Range(x*f[i], x*(f[i]+1))
	}

	return &bloom{
		bitmap: newBitmap(x * 5),
		seeds:  seeds,
	}
}

func (b *bloom) add(x int) {
	x *= 5
	for i := range b.seeds {
		hx := x % b.seeds[i]
		fmt.Println(hx)
		b.set(hx)
	}
}

func (b *bloom) isSet(x int) bool {
	x *= 5
	for i := range b.seeds {
		hx := x % b.seeds[i]
		if b.bitmap.isSet(hx) {
			return true
		}
	}
	return false
}

func (b *bloom) String() string {
	return fmt.Sprintf("seeds(%v) %b", b.seeds, b.bitmap.m)
}
