/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */
package leetcode

// @lc code=start
func mySqrt(x int) int {

	ans := -1
	l, r := 0, x

	for l <= r {
		mid := l + (r-l)/2
		if mid*mid > x {
			r = mid - 1
		} else {
			ans = mid
			l = mid + 1
		}
	}

	return ans
}

// @lc code=end
