/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */
package leetcode

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for  left <= right{
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1	
		} else {
			left = mid + 1
		}
	}
	return -1
}
// @lc code=end

