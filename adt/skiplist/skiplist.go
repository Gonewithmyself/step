package skiplist

import (
	"fmt"
	"math"
	"step/misc/randtools"
	"strings"
)

type skiplist struct {
	lv   int
	len  int
	head *node
}

type node struct {
	forwards []*node
	score    int
	data     interface{}
}

func newSkiplist() *skiplist {
	return &skiplist{
		head: &node{
			forwards: make([]*node, maxLevel),
		},
	}
}

func (s *skiplist) search(score int) *node {
	cur := s.head
	for i := s.lv - 1; i >= 0; i-- {
		for cur.forwards[i] != nil && cur.forwards[i].score < score {
			cur = cur.forwards[i]
		}
	}

	cur = cur.forwards[0]
	if cur != nil && cur.score == score {
		return cur
	}
	return nil
}

func (s *skiplist) insert(score int, val interface{}) *node {
	cur := s.head
	path := make([]*node, maxLevel)
	for i := s.lv - 1; i >= 0; i-- {
		for cur.forwards[i] != nil && cur.forwards[i].score < score {
			cur = cur.forwards[i]
		}
		path[i] = cur
	}

	cur = cur.forwards[0]
	if cur != nil && cur.score == score {
		return nil
	}

	lv := randLevel()
	if lv > s.lv {
		path[s.lv] = s.head
		s.lv++
		lv = s.lv
	}

	n := &node{
		score:    score,
		data:     val,
		forwards: make([]*node, lv),
	}
	s.len++

	for i := 0; i < lv; i++ {
		n.forwards[i] = path[i].forwards[i]
		path[i].forwards[i] = n
	}
	return n
}

func (s *skiplist) String() string {
	var buf strings.Builder
	buf.WriteString(fmt.Sprintf("lv(%v) len(%v) nodes:",
		s.lv, s.len))

	cur := s.head
	m := make(map[int]int)
	for cur.forwards[0] != nil {
		lv := 0
		for i := 0; i < s.lv; i++ {
			if i == len(cur.forwards) ||
				cur.forwards[i] == nil {
				break
			}

			if cur.forwards[i].score != cur.forwards[0].score {
				m[cur.forwards[i].score] = i + 1
				continue
			}

			lv++
		}
		cur = cur.forwards[0]
		if x, ok := m[cur.score]; ok {
			lv = x
		}
		fx := " "
		if lv != 1 {
			fx = "*"
		}
		buf.WriteString(fmt.Sprintf("%v,%v%v ", cur.score, lv, fx))
	}

	return buf.String()
}

func (s *skiplist) delete(score int) *node {
	cur := s.head
	path := make([]*node, maxLevel)
	for i := s.lv - 1; i >= 0; i-- {
		for cur.forwards[i] != nil && cur.forwards[i].score < score {
			cur = cur.forwards[i]
		}
		path[i] = cur
	}

	cur = cur.forwards[0]
	if cur == nil || cur.score != score {
		return nil
	}

	s.len--
	for i := 0; i < s.lv; i++ {
		if path[i].forwards[i] != cur {
			return nil
		}
		path[i].forwards[i] = cur.forwards[i]
	}
	return cur
}

const (
	maxLevel         = 16
	p        float32 = 0.25
)

func randLevel() int {
	lv := 1
	for float32(randtools.Range(1, math.MaxInt32))/float32(math.MaxInt32) < p && lv < maxLevel {
		lv++
	}

	return lv
}
