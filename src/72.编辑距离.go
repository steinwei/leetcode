/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */
package leetcode

import "math"

// @lc code=start
func minDistance(word1 string, word2 string) int {
	var (
		n, m = len(word1), len(word2)
		f = make([][]int, n+1)
		min func(x,y int) int
	)

	min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i := range f {
		f[i] = make([]int, m+1)
		for j := range f[i] {
			f[i][j] = math.MaxInt64
		}
	}

	f[0][0] = 0
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if j < m {
				f[i][j+1] = min(f[i][j+1], f[i][j]+1)
			}
			if i < n {
				f[i+1][j] = min(f[i+1][j], f[i][j]+1)
			}
			if (i < n && j < m) && word1[i] == word2[j] {
				f[i+1][j+1] = min(f[i+1][j+1], f[i][j])
			}
			if i<n&&j<m{
				f[i+1][j+1] = min(f[i+1][j+1], f[i][j]+1)
			}
		}
	}

	return f[n][m]
}

// @lc code=end
