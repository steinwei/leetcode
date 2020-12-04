/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */
package leetcode

// @lc code=start
func twoSum(numbers []int, target int) []int {
	ans := []int{}

	lo, hi := 0, len(numbers)-1

	for hi > lo {
		sum := numbers[hi] + numbers[lo]
		if sum == target {
			break
		} else if sum > target {
			hi--
		} else {
			lo++
		}
	}

	ans = append(ans, lo+1, hi+1)

	return ans
}

// @lc code=end
