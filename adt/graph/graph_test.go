package graph

import (
	"step/misc/randtools"
	"testing"
)

func Test_newGraph(t *testing.T) {
	g := newGraph()
	for i := 0; i < 3; i++ {
		// n := randtools.Range(3, 9)
		for j := 0; j < 3; j++ {
			g.add(i, randtools.Range(0, 10))
		}
	}

	g.bfs()
	g.dfs()
}
