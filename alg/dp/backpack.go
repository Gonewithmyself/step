package dp

import (
	"step/misc/randtools"
)

type packet struct {
	cap int

	items []int
}

func newPacket() *packet {
	n := randtools.Range(10, 11)
	items := make([]int, n)
	total := 0
	for i := 0; i < n; i++ {
		items[i] = randtools.Range(1, 10)
		total += items[i]
	}

	cap := randtools.Range(total/2, total)
	return &packet{cap, items}
}

func (p *packet) put() int {
	// write your code here
	// f[i][j] 前i个物品，是否能装j
	// f[i][j] =f[i-1][j] f[i-1][j-a[i] j>a[i]
	// f[0][0]=true f[...][0]=true
	// f[n][X]
	f := make([][]bool, len(p.items)+1)
	for i := 0; i <= len(p.items); i++ {
		f[i] = make([]bool, p.cap+1)
	}
	f[0][0] = true

	for i := 1; i <= len(p.items); i++ {
		for j := 0; j <= p.cap; j++ {
			f[i][j] = f[i-1][j]
			if j-p.items[i-1] >= 0 && f[i-1][j-p.items[i-1]] {
				f[i][j] = true
			}
		}
	}

	for i := p.cap; i >= 0; i-- {
		if f[len(p.items)][i] {
			return i
		}
	}
	return 0
}
