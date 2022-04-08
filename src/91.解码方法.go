/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */
package leetcode

// @lc code=start
func numDecodings(s string) int {
	var (
		f = make([]int, 101)
		n = len(s)
	)

	f[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1]-'0' != 0 {
			f[i] += f[i-1]
		}

		if i>1 && (s[i-2]-'0')*10 + (s[i-1]-'0') >= 10 &&
		(s[i-2]-'0')*10 + (s[i-1]-'0') <= 26 {
			f[i] += f[i-2]
		}
	}

	return f[n]
}
// @lc code=end

