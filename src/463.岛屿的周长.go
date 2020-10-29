/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 */
package leetcode

// @lc code=start
func islandPerimeter(grid [][]int) int {
	ans := 0

	n := len(grid)

	if n == 0 {
		return ans
	}

	m := len(grid[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				ans += dfs1(i, j, grid)
				// fmt.Println(ans)
			}
		}
	}

	return ans
}

func dfs1(x int, y int, grid [][]int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return 1
	} else if grid[x][y] == 0 {
		return 1
	} else if grid[x][y] == 2 {
		return 0
	}

	grid[x][y] = 2

	// if

	return dfs1(x-1, y, grid) + dfs1(x+1, y, grid) + dfs1(x, y-1, grid) + dfs1(x, y+1, grid)
}

// @lc code=end
