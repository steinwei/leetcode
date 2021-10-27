/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */
package leetcode

// @lc code=start
func findDisappearedNumbers(nums []int) (ans []int) {
	n := len(nums)

	for _, v := range nums {
		v = (v-1) % n
		nums[v] += n
	}

	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}

	return
}
// @lc code=end

