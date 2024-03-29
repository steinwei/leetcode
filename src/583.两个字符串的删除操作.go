/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] 两个字符串的删除操作
 */
package leetcode

// @lc code=start
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m + 1)

	for i := range dp {
		dp[i] = make([]int, n + 1)
		dp[i][0] = i
	}

	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	return dp[m][n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
// @lc code=end

