/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] 扰乱字符串
 */
package leetcode

// @lc code=start
func isScramble(s1 string, s2 string) bool {
	var (
		n = len(s1)
		f = make([][][]int, 31)
		dp func(l1,r1,l2,r2 int) int
	)

	// initial
	for i := range f {
		f[i] = make([][]int, 31)
		for j := range f[i] {
			f[i][j] = make([]int, 31)
			for k := range f[i][j] {
				f[i][j][k] = -1
			}
		}
	}

	dp = func(l1, r1, l2, r2 int) int {
		if f[l1][r1][l2] != -1 {
			return f[l1][r1][l2]
		}

		if l1 == r1 {
			if s1[l1-1] == s2[l2-1] {
				f[l1][r1][l2] = 1
			} else {
				f[l1][r1][l2] =  0
			}
			return f[l1][r1][l2]
		}

		f[l1][r1][l2] = 0
		for i := l1; i < r1; i++ {
			if dp(l1, i, l2, i-l1+l2) == 1 && dp(i+1, r1, i-l1+l2+1, r2) == 1{
				f[l1][r1][l2] = 1
			}
			if dp(l1, i, r1-i+l2, r2) == 1 && dp(i+1, r1, l2, r1-i-1+l2) == 1 {
				f[l1][r1][l2] = 1
			}
		}

		return f[l1][r1][l2]
	}

	return dp(1, n, 1, n) == 1
}
// @lc code=end

