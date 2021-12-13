/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */
package leetcode

import "sort"

// @lc code=start
func fourSum(nums []int, target int) (ans [][]int) {
	var (
		n = len(nums)
	)

	// sort 
	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i-1]==nums[i] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i+1; j < n-2; j++ {
			if j > i+1 && nums[j-1]==nums[j] || nums[j]+nums[n-3]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for l,r:=j+1,n-1;l<r; {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					ans = append(ans, []int{nums[i],nums[j],nums[l],nums[r]})
					for r--;l<r&&nums[r] == nums[r+1]; r-- {}
					for l++;l<r&&nums[l] == nums[l-1]; l++ {}
				}
	
				if sum > target {
					r --
				}
	
				if sum < target {
					l ++
				}
			}
		}
	}

	return
}
// @lc code=end

