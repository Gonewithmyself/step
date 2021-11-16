package dp

import (
	"fmt"
	"math"
	"step/misc/randtools"
)

type road struct {
	pmap [][]int
}

func newRoad() *road {

	pmap := [][]int{
		{1, 3, 5, 9},
		{2, 1, 3, 4},
		{5, 2, 6, 7},
		{6, 8, 4, 3},
	}
	return &road{pmap}
}

// dp[i][j] = min(dp[i-1][j], dp[i,j-1])
func (r *road) search() int {
	dp := make([][]int, len(r.pmap))
	for i := range dp {
		dp[i] = make([]int, len(r.pmap[0]))
	}

	dp[0][0] = r.pmap[0][0]
	for i := 0; i < len(r.pmap); i++ {
		for j := 0; j < len(r.pmap[0]); j++ {
			var (
				up   = math.MaxInt
				left = math.MaxInt
			)

			if i-1 >= 0 {
				up = dp[i-1][j]
			}

			if j-1 >= 0 {
				left = dp[i][j-1]
			}

			min := up
			if left < up {
				min = left
			}

			if min == math.MaxInt {
				min = 0
			}

			dp[i][j] = min + r.pmap[i][j]
		}
	}
	fmt.Println(dp)
	return dp[len(r.pmap)-1][len(r.pmap[0])-1]
}

type robotpath struct {
	m, n int
}

func newrobotpath() *robotpath {
	n := randtools.Range(4, 5)
	m := n
	return &robotpath{
		m: m,
		n: n,
	}
}

// dp[i][j]= dp[i-1][j] + dp[i][j-1]
func (r *robotpath) maxpath() int {
	dp := make([][]int, r.m)
	for i := 0; i < r.m; i++ {
		dp[i] = make([]int, r.n)
	}

	dp[0][0] = 1
	for i := 1; i < r.m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < r.n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < r.m; i++ {
		for j := 1; j < r.n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	fmt.Println(dp)
	return 0
}

// dp[i][j]= dp[i-1][j] + dp[i][j-1]
func (r *robotpath) maxpathWithobstacles(stones [][]int) int {
	dp := make([][]int, r.m)
	for i := 0; i < r.m; i++ {
		dp[i] = make([]int, r.n)
	}

	dp[0][0] = 1
	for i := 1; i < r.m; i++ {
		if stones[i][0] == 0 && stones[i-1][0] == 0 {
			dp[i][0] = 1
		}

	}

	for i := 1; i < r.n; i++ {
		if stones[0][i] == 0 && stones[0][i-1] == 0 {
			dp[0][i] = 1
		}

	}

	for i := 1; i < r.m; i++ {
		for j := 1; j < r.n; j++ {
			up := dp[i-1][j]
			if stones[i-1][j] != 0 {
				up = 0
			}

			left := dp[i][j-1]
			if stones[i][j-1] != 0 {
				left = 0
			}
			dp[i][j] = up + left
		}
	}
	fmt.Println(dp)
	return 0
}
