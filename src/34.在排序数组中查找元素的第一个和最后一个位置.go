/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */
package leetcode

// @lc code=start
func searchRange(nums []int, target int) (ans []int) {

	lo, hi := 0, len(nums)-1

	if lo > hi {
		ans = append(ans, -1, -1)
		return
	}

	ans = append(ans, searchB(nums, target, lo, hi), searchD(nums, target, lo, hi))

	return
}

func searchB(nums []int, target, lo, hi int) int {
	if lo > hi {
		return -1
	}

	mid := lo + (hi-lo)/2

	// fmt.Println(mid)
	// 确立上下边界
	if target == nums[mid] {
		if (mid == 0) || nums[mid-1] < target {
			return mid
		}
	}

	if target <= nums[mid] {
		return searchB(nums, target, lo, mid-1)
	} else {
		return searchB(nums, target, mid+1, hi)
	}
}

func searchD(nums []int, target, lo, hi int) int {
	if lo > hi {
		return -1
	}

	mid := lo + (hi-lo)/2

	// fmt.Println(mid)

	// 确立上下边界
	if target == nums[mid] && (mid == len(nums)-1 || nums[mid+1] > target) {
		return mid
	}

	if target < nums[mid] {
		return searchD(nums, target, lo, mid-1)
	} else {
		return searchD(nums, target, mid+1, hi)
	}
}

// @lc code=end
