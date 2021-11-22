package array

import (
	"fmt"
	"step/misc/randtools"
	"testing"
)

func Test_heap_build(t *testing.T) {
	h := heap{65, 72, 96, 91, 63, 96, 83, 91, 91, 65}
	h.build()

	n := len(h)
	for i := 0; i < n; i++ {
		x, ok := h.pop()
		fmt.Println(x, ok)
	}
	fmt.Println(h)

	for i := 10; i > 0; i-- {
		h.push(randtools.Range(1, 100))
	}
	fmt.Println(h)

	n = len(h)
	for i := 0; i < n; i++ {
		x, ok := h.pop()
		fmt.Println(x, ok)
	}

	fmt.Println(h)
}
