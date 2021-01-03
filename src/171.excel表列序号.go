/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 */
package leetcode

// @lc code=start
func titleToNumber(s string) int {
	ret := 0
	runes := []rune(s)
	for _, c := range runes {
		ret = 26*ret + (int(c-'A') + 1)
	}
	return ret
}

// @lc code=end
