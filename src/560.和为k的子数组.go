/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */
package leetcode

// @lc code=start
func subarraySum(nums []int, k int) int {
	ans := 0
	alias := map[int]int{}

	alias[0] = 1
	sum0 := 0
	for _, val := range nums {
		sum0 += val
		sum1 := sum0 - k
		if alias[sum1] > 0 {
			ans += alias[sum1]
		}
		alias[sum0]++
	}

	return ans
}

// @lc code=end
