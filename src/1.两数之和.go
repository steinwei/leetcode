/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package leetcode

// @lc code=start
func twoSum(nums []int, target int) (ans []int) {
	var (
		hashmap = map[int]int{}
	)

	for i, v := range nums {
		if p, ok:= hashmap[v]; ok {
			ans = append(ans, p, i)
			return
		} else {
			hashmap[target - v] = i
		}
	}

	return
}
// @lc code=end

