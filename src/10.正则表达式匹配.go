/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */
package leetcode

// @lc code=start
func isMatch(s string, p string) bool {
	var (
		n,m = len(s), len(p)
		dp = make([][]bool, 21)
	)

	// initial
	for i := range dp {
		dp[i] = make([]bool, 31)
	}

	dp[0][0] = true
	// loop
	for i := 0; i <= n; i++ {
		for j := 0; j < m; j++ {
			if dp[i][j] {
				if p[j] == '.' && i != n {
					dp[i+1][j+1] = true
				}
				if p[j] >= 'a' && p[j] <= 'z' && i != n && p[j] == s[i] {
					dp[i+1][j+1] = true
				}
				if j < m-1 && p[j+1] == '*' {
					dp[i][j+2] = true
				}
				if p[j] == '*' {
					dp[i][j+1] = true
					if p[j-1] == '.' && i != n {
						dp[i+1][j] = true
					}
					if p[j-1] >= 'a' && p[j-1] <= 'z' && i != n && s[i] == p[j-1] {
						dp[i+1][j] = true
					}
				}
			}
		}
	}

	return dp[n][m]
}
// @lc code=end

