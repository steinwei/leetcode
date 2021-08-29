/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */
package leetcode

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	ans, count, n:= 0, 0, len(nums)

	if n == 1 {
		return ans
	}

	d:=nums[1] - nums[0]
	for i := 2; i < n; i++ {

		if d == nums[i] - nums[i-1] {
			count ++
		} else {
			count = 0
			d = nums[i] - nums[i-1]
		}
		ans += count
	}

	return ans
}
// @lc code=end

