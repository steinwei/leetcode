/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 */
package leetcode

// @lc code=start
func longestPalindromeSubseq(s string) int {
	n := len(s)

	dp := make([][]int, n)

	for i := range s {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i > -1; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			//  
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				// 看丢弃哪个值
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

