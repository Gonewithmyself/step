package greedy

import (
	"math/rand"
	"testing"
	"time"
)

func TestCandyChildren_assign(t *testing.T) {
	c := newCandyChildren()
	t.Log(c)
	c.assign()
}

func init() {
	rand.Seed(time.Now().Unix())
}

func TestCurrancy_Change(t *testing.T) {
	c := newCurrancy()
	amt := rand.Intn(c.total)
	t.Log(c, amt)
	c.Change(amt)
}

func Test_newInternalCover(t *testing.T) {
	tc := newInternalCover()
	ni := tc.nonIntersection()
	t.Log(tc, ni)

	tc = &IntervalCover{
		intervals: []interval{{1, 5}, {2, 4}, {3, 5}, {5, 9}, {6, 8}, {8, 10}},
	}
	ni = tc.nonIntersection()
	t.Log(tc, ni)
}
