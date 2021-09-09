/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package leetcode

import "sort"

// @lc code=start
func threeSum(nums []int) (ans [][]int) {
	n := len(nums)

	// sort
	sort.Ints(nums)

	for first := 0; first < n; first++ {
		// jump 
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}

		third := n - 1
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}

			for second < third && nums[second] + nums[third] > target {
				third --
			}
			
			result := nums[first] + nums[second] + nums[third]

			if second == third {
				break
			}
			
			if result == 0 {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}


	return
}
// @lc code=end

