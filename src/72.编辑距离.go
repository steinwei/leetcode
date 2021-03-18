/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */
package leetcode

// @lc code=start
func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)

	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	for i := 0; i <= n1; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= n2; i++ {
		dp[0][i] = i
	}

	var Min func(arg ...int) int
	Min = func(arg ...int) int {
		min := arg[0]
		for _, val := range arg {
			if val < min {
				min = val
			}
		}
		return min
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(
					dp[i-1][j-1]+1,
					dp[i-1][j]+1,
					dp[i][j-1]+1)
			}
		}
	}

	return dp[n1][n2]
}

// @lc code=end
