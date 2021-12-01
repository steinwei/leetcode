/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] 扰乱字符串
 */
package leetcode

// @lc code=start
func isScramble(s1 string, s2 string) bool {
	var (
		s1len = len(s1)
		s2len = len(s2)
		dp = make([][][]bool, s1len+1)
	)

	if s1len != s2len {
		return false
	}

	// initial
	for i := 0; i < s1len; i++ {
		dp[i] = make([][]bool, s1len+1)
		for j := 0; j < s1len; j++ {
			dp[i][j] = make([]bool, s1len+1)
			dp[i][j][1] = s1[i] == s2[j]
		}
	}

	// 
	for len := 2; len <= s1len; len++ {
		for i := 0; i+len <= s1len; i++ {
			for j := 0; j+len <= s1len; j++ {
				for leftLen := 1; leftLen < len && !dp[i][j][len]; leftLen++ {
					c1 := dp[i][j][leftLen] && dp[i+leftLen][j+leftLen][len-leftLen]
					c2 := dp[i][j+len-leftLen][leftLen] && dp[i+leftLen][j][len-leftLen]
					dp[i][j][len] = c1 || c2
				}
			}
		}
	}

	return dp[0][0][s1len]
}
// @lc code=end

