/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */
package leetcode

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	var (
		n, m, k = len(s1), len(s2), len(s3)
		f = make([][]bool, 101)
	)
	
	for i := range f {
		f[i] = make([]bool, 101)
	}

	if n + m != k {
		return false
	}

	f[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if f[i][j] {
				if i<n && s1[i] == s3[i+j] {
					f[i+1][j] = true
				}
				if j<m && s2[j] == s3[i+j] {
					f[i][j+1] = true
				}
			}
		}
	}

	return f[n][m]
}
// @lc code=end

