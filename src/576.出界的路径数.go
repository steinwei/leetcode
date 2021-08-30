/*
 * @lc app=leetcode.cn id=576 lang=golang
 *
 * [576] 出界的路径数
 */
package leetcode

// @lc code=start

const mod int = 1e9 + 7
var dirs = []struct{x,y int}{{-1,0},{1,0},{0,-1},{0,1}}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	ans := 0

	dp := make([][][]int, maxMove + 1)

	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}

	dp[0][startRow][startColumn] = 1
	for L := 0; L < maxMove; L++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				count := dp[L][i][j]
				if count > 0 {
					for _, dir := range dirs {
						i1, j1 := i+ dir.x, j+dir.y
						if i1 >= 0 && i1 <m && j1 >=0 && j1 <n {
							dp[L+1][i1][j1] = (dp[L+1][i1][j1] + count) % mod
						} else {
							ans = (ans + count) % mod
						}
					}
				}
			}
		}
	}

	return ans
}
// @lc code=end

