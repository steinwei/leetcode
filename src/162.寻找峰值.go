/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */
package leetcode

// @lc code=start
func findPeakElement(nums []int) int {
	ans := 0

	for i, v := range nums {
		if v > nums[ans] {
			ans = i
		}
	}

	return ans
}
// @lc code=end

