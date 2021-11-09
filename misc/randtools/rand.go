package randtools

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Range [start, end)
func Range(start, end int) int {
	return rand.Intn(end-start) + start
}

// Range2 [start, end]
func Range2(start, end int) int {
	return rand.Intn(end-start+1) + start
}
