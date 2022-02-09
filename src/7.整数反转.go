/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
package leetcode

import "math"

// @lc code=start
func reverse(x int) int {
	ans := 0

	for x!=0 {
		if math.MinInt32/10 > ans || ans > math.MaxInt32/10 {
			return 0
		}
		digit := x%10
		x/=10
		ans = ans*10 + digit
	}

	return ans
}
// @lc code=end

