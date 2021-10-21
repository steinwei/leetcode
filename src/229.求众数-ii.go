/*
 * @lc app=leetcode.cn id=229 lang=golang
 *
 * [229] 求众数 II
 */
package leetcode

import "sort"

// @lc code=start
func majorityElement(nums []int) (ans []int) {
	n := len(nums)

	if n == 1 {
		ans = append(ans, nums...)
		return
	}

	sort.Ints(nums)

	cnt := 1

	for i := 1; i < n; i++ {
		if nums[i-1] == nums[i] {
			cnt ++
			if i == n - 1 {
				if cnt * 3 > n {
					ans = append(ans, nums[i-1])
				}
			}
		} else {
			if cnt * 3 > n {
				ans = append(ans, nums[i-1])
			}

			// reset
			cnt = 1
		}
	}
	
	return
}
// @lc code=end

