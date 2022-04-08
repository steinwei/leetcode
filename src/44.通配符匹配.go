/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */
package leetcode

// @lc code=start
func isMatch(s string, p string) bool {
	var (
		n, m = len(s), len(p)
		dp = make([][]bool, n+1)
	)

	for i := range dp {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dp[i][j] {
				if p[j] <= 'z' && p[j] >= 'a' && p[j] == s[i] {
					dp[i+1][j+1] = true
				}
				if p[j] == '*' {
					dp[i][j+1] = true
				}
				if p[j] == '*' && i < n {
					dp[i+1][j] = true
				}
				if p[j] == '?' {
					dp[i+1][j+1] = true
				}
			}
		}
	}

	return dp[n][m]
}
// @lc code=end

