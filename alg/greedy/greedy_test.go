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
