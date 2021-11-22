package array

type heap []int

func (h heap) build() {
	for i := (len(h) - 2) / 2; i >= 0; i-- {
		h.adjustDown(i, len(h))
	}
}

func (h *heap) push(x int) {
	*h = append(*h, x)
	h.adjustUp()
}

func (h *heap) pop() (int, bool) {
	n := len(*h)
	if n < 1 {
		return -1, false
	}

	top := (*h)[0]
	(*h)[0] = (*h)[n-1]
	*h = (*h)[0 : n-1]

	if n > 2 {
		h.adjustDown(0, n-1)
	}

	return top, true
}

func (h heap) adjustDown(parent, n int) {
	child := 2*parent + 1
	x := h[parent]
	for child < n {
		right := child + 1
		if right < n && h[right] < h[child] {
			child = right
		}

		if x <= h[child] {
			break
		}

		h[parent] = h[child]
		parent = child
		child = 2*child + 1
	}
	h[parent] = x
}

func (h heap) adjustUp() {
	child := len(h) - 1
	parent := (child - 1) / 2
	x := h[child]
	for child > 0 && x < h[parent] {
		h[child] = h[parent]
		child = parent
		parent = (child - 1) / 2
	}
	h[child] = x
}
