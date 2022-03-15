/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */
package leetcode

// @lc code=start
func missingNumber(nums []int) int {
	var (
		n = len(nums)
		ans = 0
		hashmap = map[int]bool{}
	)

	for _, v := range nums {
		hashmap[v] = true
	}

	for i := 0; i <= n; i++ {
		if !hashmap[i] {
			return i
		}
	}

	return ans
}
// @lc code=end

