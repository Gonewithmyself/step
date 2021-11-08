package greedy

import (
	"fmt"
	"math/rand"
	"sort"
)

// 总限制  极值
// 每次选择 同等限制 对极值贡献最大
// 每次选择 不影响后续选择

type CandyChildren struct {
	candiesSize    []int
	childrenDemand []int
}

func newCandyChildren() *CandyChildren {
	c := &CandyChildren{
		candiesSize:    make([]int, 7),
		childrenDemand: make([]int, 10),
	}

	sz := [...]int{1, 2, 3}

	for i := range c.candiesSize {
		c.candiesSize[i] = sz[rand.Intn(2)]
	}
	sort.Ints(c.candiesSize)

	for i := range c.childrenDemand {
		c.childrenDemand[i] = sz[rand.Intn(3)]
	}
	sort.Ints(c.childrenDemand)

	return c
}

func (c *CandyChildren) assign() {
	j := 0
	for i := range c.childrenDemand {
		if j == len(c.candiesSize) {
			break
		}

		for ; j < len(c.candiesSize); j++ {
			if c.candiesSize[j] >= c.childrenDemand[i] {
				fmt.Println("assign ", i, j)
				j++
				break
			}
		}
	}
}

type Currancy struct {
	value [7]int
	count [7]int
	total int
}

func newCurrancy() *Currancy {
	c := &Currancy{
		value: [7]int{1, 2, 5, 10, 20, 50, 100},
	}

	for i := range c.count {
		if i == 0 {
			c.count[0] = 100000
			continue
		}

		c.count[i] = rand.Intn(100)
		c.total += c.count[i] * c.value[i]
	}
	return c
}

func (c Currancy) Change(amt int) {
	for amt > 0 {
		for i := 6; i >= 0; i-- {
			count := amt / c.value[i]
			if count == 0 {
				continue
			}

			if count > c.count[i] {
				count = c.count[i]
			}

			amt -= c.value[i] * count
			fmt.Println("give", c.value[i], count)
		}
	}
}
