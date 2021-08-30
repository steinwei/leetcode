/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */
package leetcode

// @lc code=start
func trap(height []int) int {

	ans := 0

	stack := []int{}

	for i, h := range height {
		for len(stack) > 0 &&  height[stack[len(stack) - 1]] < h {
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			if len(stack) <= 0 {
				break
			}

			left := stack[len(stack) - 1]
			curWidth := i - left - 1
			curHeight := min(h, height[left]) - height[top]
			ans += curHeight * curWidth
		}
		stack = append(stack, i)
	}

	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
// @lc code=end

