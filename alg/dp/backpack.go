package dp

import (
	"step/misc/randtools"
)

type packet struct {
	cap int

	items  []int
	values []int
}

func newPacket() *packet {
	n := randtools.Range(10, 11)
	items := make([]int, n)
	vals := make([]int, n)
	total := 0
	for i := 0; i < n; i++ {
		items[i] = randtools.Range(1, 10)
		vals[i] = randtools.Range(1, 10)
		total += items[i]
	}

	cap := randtools.Range(total/2, total)
	return &packet{cap, items, vals}
}

func (p *packet) put() int {
	states := make([][]bool, len(p.items))
	for i := range states {
		states[i] = make([]bool, p.cap+1)
	}

	states[0][0] = true
	states[0][p.items[0]] = true

	for i := 1; i < len(p.items); i++ {
		for j := 0; j <= p.cap; j++ {
			if states[i-1][j] {
				states[i][j] = true
			}
		}

		for j := 0; j <= p.cap-p.items[i]; j++ {
			if states[i-1][j] {
				states[i][j+p.items[i]] = true
			}
		}
	}

	for i := p.cap; i >= 0; i-- {
		if states[len(p.items)-1][i] {
			return i
		}
	}
	return 0
}

// state[i] = state[i]
func (p *packet) putDownUp() int {
	states := make([]bool, p.cap+1)
	states[p.items[0]] = true
	for i := 1; i < len(p.items); i++ {
		for j := p.cap - p.items[i]; j >= 0; j-- {
			if states[j] {
				states[j+p.items[i]] = true
			}
		}
	}

	for i := p.cap; i >= 0; i-- {
		if states[i] {
			return i
		}
	}
	return 0
}

// state[i] = state[i]
func (p *packet) putValue() int {
	states := make([][]int, len(p.items))
	for i := range states {
		states[i] = make([]int, p.cap+1)
	}

	states[0][0] = 0
	states[0][p.items[0]] = p.values[0]

	for i := 1; i < len(p.items); i++ {
		for j := 0; j <= p.cap; j++ {
			if states[i-1][j] >= 0 {
				states[i][j] = states[i-1][j]
			}
		}

		for j := 0; j <= p.cap-p.items[i]; j++ {
			if states[i-1][j] >= 0 {
				v := states[i-1][j] + p.values[i]
				if v > states[i][j+p.items[i]] {
					states[i][j+p.items[i]] = v
				}
			}
		}
	}

	max := -1
	for i := p.cap; i >= 0; i-- {
		if states[len(p.items)-1][i] >= max {
			max = states[len(p.items)-1][i]
		}
	}
	return max
}
