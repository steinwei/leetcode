/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */
package leetcode

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	n, m := len(matrix), len(matrix[0])

	dp := make([][]int, n)
	ans := 0

	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				ans = 1
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j],dp[i][j-1]), dp[i-1][j-1]) + 1
				if ans < dp[i][j] {
					ans = dp[i][j]
				}
			}
		}
	}

	return ans*ans
}

func min(x,y int) int {
	if x > y {
		return y
	}
	return x
}
// @lc code=end

