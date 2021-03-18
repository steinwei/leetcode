/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */
package leetcode

// @lc code=start
func removeElement(nums []int, val int) int {
	n := len(nums)

	fast, slow := n-1, 0
	for fast >= slow {
		if nums[slow] == val {
			nums[slow] = nums[fast]
			// fmt.Println(nums, slow, fast)
			fast--
		} else {
			slow++
		}
	}
	return fast + 1
}

// @lc code=end
