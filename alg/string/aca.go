package cmp

import (
	"sort"
	"strings"
	"unicode/utf8"
)

type node struct {
	children   map[rune]*node
	fail       *node
	wordLength int
}

type ACA struct {
	root      *node
	nodeCount int
}

// New returns an empty aca.
func New() *ACA {
	return &ACA{root: &node{}, nodeCount: 1}
}

// Add adds a new word to aca.
// After Add, and before Find,
// MUST Build.
func (a *ACA) Add(word string) {
	n := a.root
	for _, r := range word {
		if n.children == nil {
			n.children = make(map[rune]*node)
		}
		if n.children[r] == nil {
			n.children[r] = &node{}
			a.nodeCount++
		}
		n = n.children[r]
	}
	n.wordLength = len(word)
}

// Del delete a word from aca.
// After Del, and before Find,
// MUST Build.
func (a *ACA) Del(word string) {
	rs := []rune(word)
	stack := make([]*node, len(rs))
	n := a.root

	for i, r := range rs {
		if n.children[r] == nil {
			return
		}
		stack[i] = n
		n = n.children[r]
	}

	// if it is NOT the leaf node
	if len(n.children) > 0 {
		n.wordLength = 0
		return
	}

	// if it is the leaf node
	for i := len(rs) - 1; i >= 0; i-- {
		stack[i].children[rs[i]].children = nil
		stack[i].children[rs[i]].fail = nil

		delete(stack[i].children, rs[i])
		a.nodeCount--
		if len(stack[i].children) > 0 ||
			stack[i].wordLength > 0 {
			return
		}
	}
}

// Build builds the fail pointer.
// It MUST be called before Find.
func (a *ACA) Build() {
	// allocate enough memory as a queue
	q := append(make([]*node, 0, a.nodeCount), a.root)

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		for r, child := range node.children {
			q = append(q, child)

			p := node.fail
			for p != nil {
				// ATTENTION: nil map cannot be writen
				// but CAN BE READ!!!
				if p.children[r] != nil {
					child.fail = p.children[r]
					break
				}
				p = p.fail
			}
			if p == nil {
				child.fail = a.root
			}
		}
	}
}

func (a *ACA) find(s string, cb func(start, end int)) {
	n := a.root
	for i, r := range s {
		for n.children[r] == nil && n != a.root {
			n = n.fail
		}
		n = n.children[r]
		if n == nil {
			n = a.root
			continue
		}

		end := i + utf8.RuneLen(r)
		for t := n; t != a.root; t = t.fail {
			if t.wordLength > 0 {
				cb(end-t.wordLength, end)
			}
		}
	}
}

// Find finds all the words contains in s.
// The results may duplicated.
// It is caller's responsibility to make results unique.
func (a *ACA) Find(s string) (words []string) {
	a.find(s, func(start, end int) {
		words = append(words, s[start:end])
	})
	return
}

// Block records the start and end position
// that words appear, namely s[start:end].
type Block struct {
	Start, End int
}

// Blocks returns the blocks that words in aca appear.
func (a *ACA) Blocks(s string) (blocks []Block) {
	a.find(s, func(start, end int) {
		blocks = append(blocks, Block{Start: start, End: end})
	})
	return
}

// wrapper of Blocks , union sequent blocks
func (a *ACA) SequentBlocks(s string) (blocks []Block) {
	return UnionBlocks(a.Blocks(s))
}

// Replace replace all sensitive words in s with replaceTo
func (a *ACA) Replace(s, replaceTo string) string {
	var result strings.Builder

	blocks := a.SequentBlocks(s)
	prePos := 0
	for _, block := range blocks {
		result.WriteString(s[prePos:block.Start])
		result.WriteString(replaceTo)
		prePos = block.End
	}

	if prePos > 0 {
		result.WriteString(s[prePos:])
	} else {
		return s
	}
	return result.String()
}

type byPos []Block

func (bs byPos) Len() int { return len(bs) }

func (bs byPos) Swap(i, j int) { bs[i], bs[j] = bs[j], bs[i] }

func (bs byPos) Less(i, j int) bool {
	if bs[i].Start < bs[j].Start {
		return true
	}
	if bs[i].Start == bs[j].Start {
		return bs[i].End < bs[j].End
	}
	return false
}

func UnionBlocks(blocks []Block) []Block {
	if len(blocks) == 0 {
		return blocks
	}

	sort.Sort(byPos(blocks))
	n := 0
	for i := 1; i < len(blocks); i++ {
		if blocks[i].Start <= blocks[n].End {
			if blocks[i].End > blocks[n].End {
				blocks[n].End = blocks[i].End
			}
		} else {
			n++
			blocks[n] = blocks[i]
		}
	}
	return blocks[:n+1]
}
