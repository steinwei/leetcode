/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */
package leetcode

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// base
	if len(obstacleGrid) == 0  {
		return 0
	}

	n, m := len(obstacleGrid), len(obstacleGrid[0])


	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}

	for j := 0; j < m && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[n-1][m-1]
}
// @lc code=end

