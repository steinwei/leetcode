/*
 * @lc app=leetcode.cn id=1020 lang=golang
 *
 * [1020] 飞地的数量
 */
package leetcode

// @lc code=start
func numEnclaves(grid [][]int) int {
	// var (
	// 	ans = 0
	// 	dfs func(x, y int)
	// 	n, m = len(grid), len(grid[0])
	// 	vis = make([][]bool, n)
	// )
	// for i := range vis {
	// 	vis[i] = make([]bool, m)
	// }
	// dfs = func(x, y int) {
	// 	if x < 0 || y < 0 || x > n-1 || y > m-1 || grid[x][y] == 0 || vis[x][y] {
	// 		return
	// 	}
	// 	vis[x][y] = true
	// 	dfs(x+1,y)
	// 	dfs(x,y+1)
	// 	dfs(x-1,y)
	// 	dfs(x,y-1)
	// }
	// for i := range grid {
	// 	dfs(i, 0)
	// 	dfs(i, m-1)
	// }
	// for j := 1; j < m; j++ {
	// 	dfs(0, j)
	// 	dfs(n-1, j)
	// }
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < m; j++ {
	// 		if grid[i][j] == 1 && !vis[i][j] {
	// 			ans++
	// 		}
	// 	}
	// }
	// return ans
		// 上述为 深度优先， 下面看下 广度优先

	var (
		ans = 0
		n, m = len(grid), len(grid[0])
		vis = make([][]bool, n)
		queue = [][2]int{}
		dirs = [][2]int{{0,1},{1,0},{-1,0},{0,-1}}
	)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	for i := range grid {
		if grid[i][0] == 1 {
			vis[i][0] = true
			queue = append(queue, [2]int{i,0})
		}
		if grid[i][m-1] == 1 {
			vis[i][m-1] = true
			queue = append(queue, [2]int{i,m-1})
		}
	}
	for j := 1; j < m-1; j++ {
		if grid[0][j] == 1 {
			vis[0][j] = true
			queue = append(queue, [2]int{0,j})
		}
		if grid[n-1][j] == 1 {
			vis[n-1][j] = true
			queue = append(queue, [2]int{n-1,j})
		}
	}
	for len(queue) != 0 {
		top := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for _, v := range dirs {
			dx := top[0] + v[0]
			dy := top[1] + v[1]
			if dx < 0 || dx > n-1 || dy < 0 || dy > m-1 || grid[dx][dy] == 0 || vis[dx][dy] {
				continue
			}
			vis[dx][dy] = true
			queue = append(queue, [2]int{dx,dy})
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && !vis[i][j] {
				ans ++
			}	
		}
	}

	return ans
}
// @lc code=end

