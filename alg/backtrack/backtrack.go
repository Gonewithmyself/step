package backtrack

import (
	"fmt"
	"step/misc/randtools"
	"strings"
)

type nQueen struct {
	n          int
	chessboard []int
	res        [][]int
}

func newNQueen() *nQueen {
	nq := &nQueen{
		n: randtools.Range(8, 9),
	}

	nq.chessboard = make([]int, nq.n)
	nq.placeQueen(0)
	return nq
}

func (nq *nQueen) placeQueen(row int) {
	if row == nq.n {
		board := make([]int, nq.n)
		copy(board, nq.chessboard)
		nq.res = append(nq.res, board)
		return
	}

	for col := 0; col < nq.n; col++ {
		if nq.canPlace(row, col) {
			nq.chessboard[row] = col
			nq.placeQueen(row + 1)
		}
	}
}

func (nq *nQueen) canPlace(row, col int) bool {
	leftup := col - 1
	rightup := col + 1

	for row = row - 1; row >= 0; row-- {
		if nq.chessboard[row] == col {
			return false
		}

		if leftup >= 0 && nq.chessboard[row] == leftup {
			return false
		}

		if rightup < nq.n && nq.chessboard[row] == rightup {
			return false
		}

		leftup--
		rightup++
	}
	return true
}

func (nq *nQueen) String() string {
	var buf strings.Builder
	for i := range nq.res {
		for j := range nq.res[i] {
			for k := 0; k < nq.n; k++ {
				if nq.res[i][j] == k {
					buf.WriteString("Q ")
				} else {
					buf.WriteString("x ")
				}
			}
			buf.WriteString("\n")
		}

		buf.WriteString("-----\n\n")
	}
	return buf.String()
}

type packet struct {
	cap int

	total int
	items []int
}

func newPacket() *packet {
	pt := &packet{}

	n := randtools.Range(30, 33)

	total := 0
	for i := 0; i < n; i += 7 {
		weight := randtools.Range(i, i+3) + 1
		pt.items = append(pt.items, weight)
		total += weight
	}
	pt.cap = randtools.Range(total/2, total)
	pt.total = total
	pt.put(0, 0, 0)
	return pt
}

func (pt *packet) put(i, cw, start int) {
	if i == len(pt.items) || cw == pt.cap {
		fmt.Println("done", start)
		return
	}

	pt.put(i+1, cw, start)
	if cw+pt.items[i] <= pt.cap {
		if cw == 0 {
			fmt.Printf("%v start \n", pt.items[i])
			start = pt.items[i]
		}

		// pt.gots[start] = append(pt.gots[start], pt.items[i])
		fmt.Printf("%v %v ", pt.items[i], cw)
		pt.put(i+1, cw+pt.items[i], start)
	}
}

func (pt *packet) String() string {
	var buf strings.Builder
	buf.WriteString(fmt.Sprintf("items(%v) total(%v) cap(%v)\n", pt.items, pt.total, pt.cap))
	return buf.String()
}

type rematch struct {
	patt   string
	matchd bool
}

func newReMatch(patt string) *rematch {
	return &rematch{
		patt: patt,
	}
}

func (re *rematch) match(src string) bool {
	re.matchd = false
	return re.matchd
}

func (re *rematch) doMatch(i, j, n int, src string) {
	if re.matchd {
		return
	}

	if j == len(re.patt) {
		if i == n {
			re.matchd = true
			return
		}
	}

	switch {
	case re.patt[j] == '*':
		for k := 0; k < n-i; k++ {
			re.doMatch(i+k, j+1, n, src)
		}

	case re.patt[j] == '?':
		re.doMatch(i, j+1, n, src)
		re.doMatch(i+1, j+1, n, src)

	case re.patt[j] == src[i] &&
		i < n:
		re.doMatch(i+1, j+1, n, src)
	}
}
