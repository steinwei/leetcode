/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */
package leetcode

// @lc code=start
func hammingDistance(x int, y int) int {
	ans := 0

	for s := x ^ y; s > 0; s >>= 1 {
		ans += s & 1
	}

	return ans
}
// @lc code=end

