package cmp

import (
	"math"
)

func BFcmp(s, p string) int {
	n := len(s)
	m := len(p)
	for i := range s {
		curr := i
		if curr+m > n {
			break
		}

		idx := i
		for j := range p {
			if s[curr] != p[j] {
				idx = -1
				break
			}
			curr++
		}

		if idx != -1 {
			return idx
		}
	}
	return -1
}

func BMcmp(s, p string) int {
	n := len(s)
	m := len(p)
	for i := 0; i < n; {
		if i >= n {
			break
		}

		curr := i + m - 1
		if curr > n-1 {
			break
		}

		step := 1
		idx := i
		for j := m - 1; j >= 0; j-- {
			if s[curr] != p[j] {
				idx = -1
				xi := -1
				for k := m - 1; k >= 0; k-- {
					if p[k] == s[curr] {
						xi = k
						break
					}
				}

				step = j - xi
				if step < 1 {
					step = 1
				}
				break
			}
			curr--
		}

		if idx != -1 {
			return idx
		}
		i += step
	}
	return -1
}

func BMcmpWithCache(s, p string) int {
	n := len(s)
	m := len(p)

	bc := map[byte]int{}
	for i := range p {
		bc[p[i]] = i
	}

	for i := 0; i < n; {
		if i >= n {
			break
		}

		curr := i + m - 1
		if curr > n-1 {
			break
		}

		step := 1
		idx := i
		for j := m - 1; j >= 0; j-- {
			if s[curr] != p[j] {
				idx = -1
				xi := -1
				if k, ok := bc[s[curr]]; ok {
					xi = k
				}
				step = j - xi
				if step < 1 {
					step = 1
				}
				break
			}
			curr--
		}

		if idx != -1 {
			return idx
		}
		i += step
	}
	return -1
}

func BMcmpWithGoodSuffix(s, p string) int {
	n := len(s)
	m := len(p)

	bc := map[byte]int{}
	for i := range p {
		bc[p[i]] = i
	}

	bcget := func(char byte) int {
		if idx, ok := bc[char]; ok {
			return idx
		}
		return -1
	}

	suffix := make([]int, m)
	prefix := make([]bool, m)
	for i := 0; i < m; i++ {
		suffix[i] = -1
	}

	for i := 0; i < m-1; i++ {
		j := i
		k := 0
		for j >= 0 && p[j] == p[m-1-k] {
			k++
			j--
			suffix[k] = j + 1
		}

		if j == -1 {
			prefix[k] = true
		}
	}

	moveByGS := func(j int) int {
		k := m - 1 - j
		if suffix[k] != -1 {
			return j - suffix[k] + 1
		}

		for r := j + 1; r <= m-1; r++ {
			if prefix[m-r] {
				return r
			}
		}
		return m
	}
	for i := 0; i <= n-m; {

		j := 0
		for j = m - 1; j >= 0; j-- {
			if s[i+j] != p[j] {
				break
			}
		}

		if j < 0 {
			return i
		}

		var (
			x = j - bcget(s[i+j])
			y = 0
		)

		if j < m-1 {
			y = moveByGS(j)
		}
		i += int(math.Max(float64(x), float64(y)))
	}
	return -1
}

func KMPcmp(s, p string) int {
	nexts := buildNexts(p)
	tar := 0
	pos := 0

	for tar < len(s) {
		if s[tar] == p[pos] {
			tar++
			pos++
		} else if pos != 0 {
			pos = nexts[pos-1]
		} else {
			tar++
		}

		if pos == len(p) {
			return tar - pos
		}
	}
	return -1
}

func buildNexts(p string) []int {
	// k-前缀 k-后缀
	// next[x] 使得p[:k] p[len(p)-k:len(p)]相等 中的最大k
	nexts := make([]int, len(p)) // 好前缀长度 -> 前后缀重叠长度
	nexts[0] = 0
	now := 0
	x := 1
	for x < len(p) {
		if p[now] == p[x] {
			now++
			nexts[x] = now
			x++
		} else if now != 0 {
			now = nexts[now-1]
		} else {
			nexts[x] = 0
			x++
		}
	}

	return nexts
}
