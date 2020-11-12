/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */
package leetcode

// @lc code=start
func nextPermutation(nums []int) {
	hi := len(nums) - 1
	lo := hi - 1
	for lo >= 0 && nums[lo] >= nums[lo+1] {
		lo--
	}
	if lo >= 0 {
		j := hi
		for j >= 0 && nums[j] <= nums[lo] {
			j--
		}
		nums[lo], nums[j] = nums[j], nums[lo]
	}
	reverse(nums[lo+1:])
}
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// @lc code=end
