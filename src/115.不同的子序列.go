/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */
package leetcode

// @lc code=start
func numDistinct(s string, t string) int {
	var (
		n = len(s)
		m = len(t)
		f = make([][]int, 1001)
	)

	for i := range f {
		f[i] = make([]int, 1001)
	}

	f[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= m; j++ {
			f[i+1][j] += f[i][j]
			if j < m && s[i] == t[j] {
				f[i+1][j+1] += f[i][j]
			}
		}
	}

	return f[n][m]
}
// @lc code=end

