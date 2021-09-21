/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */
package leetcode

// @lc code=start
func countSubstrings(s string) int {
	n := len(s)

	dp:= make([][]bool, n)

	for i := range dp {
		dp[i] = make([]bool, n)
	}

	cnt := 0

	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if s[i] == s[j] && (j - i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				cnt++
			}
		}
	}

	return cnt
}
// @lc code=end

