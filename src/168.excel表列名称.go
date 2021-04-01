/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */
// package leetcode

// @lc code=start
func convertToTitle(n int) string {

	if n <= 26 {
		return string('A' + (n-1)%26)
	}

	return convertToTitle((n-1)/26) + convertToTitle((n-1)%26+1)
}

// @lc code=en
