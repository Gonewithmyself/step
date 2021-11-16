package dp

import (
	"fmt"
	"testing"
)

func Test_road_search(t *testing.T) {
	r := newRoad()
	t.Log(r.search())
}

func Test_robotpath_maxpath(t *testing.T) {
	r := newrobotpath()
	t.Log(r.maxpath(), uniquePaths(4, 4))

	stones := [][]int{
		{0, 0, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	t.Log(r.maxpathWithobstacles(stones), uniquePathsWithObstacles(stones))
}

func uniquePaths(m int, n int) int {
	// f[i][j] 表示i,j到0,0路径数
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if f[i] == nil {
				f[i] = make([]int, n)
			}
			f[i][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	fmt.Println(f)
	return f[m-1][n-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// f[i][j] = f[i-1][j] + f[i][j-1] 并检查障碍物
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if f[i] == nil {
				f[i] = make([]int, n)
			}
			f[i][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 || f[i-1][0] == 0 {
			f[i][0] = 0
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 || f[0][j-1] == 0 {
			f[0][j] = 0
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				f[i][j] = 0
			} else {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}
	fmt.Println(f)
	return f[m-1][n-1]
}
