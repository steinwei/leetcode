/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

package leetcode

// @lc code=start
func search(nums []int, target int) int {
	n:=len(nums)

	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	lo,hi:=0,n-1

	for lo <= hi {
		mid := (hi+lo)/2

		// get result
		if nums[mid] == target {
			return mid
		}
		
		// 判断单调区间
		if nums[0] <= nums[mid] {
			if nums[mid] > target && nums[0] <= target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if nums[mid] < target && nums[n-1] >= target  {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return -1
}
// @lc code=end

