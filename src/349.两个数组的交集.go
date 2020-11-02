/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */
package leetcode

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	ans := []int{}

	// 存索引
	memo := map[int]bool{}

	for i := 0; i < len(nums1); i++ {
		if memo[nums1[i]] != true {
			memo[nums1[i]] = true
		}
	}

	for i := 0; i < len(nums2); i++ {
		if memo[nums2[i]] == true {
			ans = append(ans, nums2[i])
			memo[nums2[i]] = false
		}
	}

	return ans
}

// @lc code=end
