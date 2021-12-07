package skiplist

import (
	"math"
	"step/misc/randtools"
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

	for i := 0; i < lv; i++ {
		n.forwards[i] = path[i].forwards[i]
		path[i].forwards[i] = n
	}
	return n
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
	if cur != nil && cur.score == score {
		return nil
	}

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
