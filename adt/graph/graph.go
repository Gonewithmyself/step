package graph

import (
	"container/list"
	"fmt"
)

type graph struct {
	data map[int]map[int]struct{}
}

func newGraph() *graph {
	return &graph{
		data: make(map[int]map[int]struct{}),
	}
}

func (g *graph) add(me, friend int) {
	g.doAdd(me, friend)
	// g.doAdd(friend, me)
}

func (g *graph) doAdd(me, friend int) {
	l := g.data[me]
	if l == nil {
		l = map[int]struct{}{}
		g.data[me] = l
	}

	l[friend] = struct{}{}
}

func (g *graph) bfs() {
	for id, mp := range g.data {
		fmt.Printf("%d: ", id)
		for pid := range mp {
			fmt.Printf("%d ", pid)
		}
		fmt.Println("")
	}
}

type vt struct {
	p int
	c int
}

func (g *graph) dfs() {
	l := list.New()

	for id := range g.data {
		l.PushBack(&vt{p: -1, c: id})
	}

	set := map[vt]struct{}{}
	for l.Len() != 0 {
		e := l.Front()
		l.Remove(e)
		node := e.Value.(*vt)
		if _, ok := set[*node]; ok {
			continue
		}
		set[*node] = struct{}{}
		if node.p == -1 {
			fmt.Println()
		}
		fmt.Printf("%d->%d ", node.p, node.c)
		mp := g.data[node.c]
		for next := range mp {
			l.PushFront(&vt{p: node.c, c: next})
		}
	}
}
