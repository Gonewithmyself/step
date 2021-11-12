package truth

import (
	"fmt"
	"unsafe"
)

func ss() {
	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		m[i] = i * 20
	}

	fmt.Println(m)
}

const (
	bucketCnt     = 8
	loadFactorNum = 13
	loadFactorDen = 2
)

func getB(hint int) uint8 {
	b := uint8(0)
	for overLoadFactor(hint, b) {
		b++
	}
	return uint8(b)
}

func overLoadFactor(count int, B uint8) bool {
	return count > bucketCnt && uintptr(count) > loadFactorNum*(bucketShift(B)/loadFactorDen)
}

// bucketShift returns 1<<b, optimized for code generation.
func bucketShift(b uint8) uintptr {
	// Masking the shift amount allows overflow checks to be elided.
	return uintptr(1) << (b & (8*8 - 1))
}

func calculateBuckets(b uint8) uintptr {
	base := bucketShift(b)
	nbuckets := base
	if b >= 4 {
		// Add on the estimated number of overflow buckets
		// required to insert the median number of elements
		// used with this value of b.
		nbuckets += bucketShift(b - 4)
		sz := unsafe.Sizeof(bmap{}) * nbuckets
		up := roundupsize(sz)
		if up != sz {
			nbuckets = up / unsafe.Sizeof(bmap{})
		}
	}
	return nbuckets
}

func tooManyOverflowBuckets(noverflow uint16, B uint8) bool {
	// If the threshold is too low, we do extraneous work.
	// If the threshold is too high, maps that grow and shrink can hold on to lots of unused memory.
	// "too many" means (approximately) as many overflow buckets as regular buckets.
	// See incrnoverflow for more details.
	if B > 15 {
		B = 15
	}
	// The compiler doesn't see here that B < 16; mask B to generate shorter shift code.
	return noverflow >= uint16(1)<<(B&15)
}

type bmap struct {
	top  [8]uint8
	keys [8]int
	vals [8]int
	next *bmap
}
