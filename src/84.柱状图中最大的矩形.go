/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */
package leetcode

// @lc code=start
func largestRectangleArea(heights []int) int {
	var (
		ans = 0
		stack = []int{}
		n = len(heights)
		left = make([]int, n)
		right = make([]int, n)
	)
	for i := 0; i < n; i++ {
		// 单调栈 遇到小的全部出栈
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		// 遇到大的 入栈
		stack = append(stack, i)
	}
	stack = []int{}
	for i := n-1; i > -1; i-- {
		// 单调栈 遇到小的全部出栈
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		// 遇到大的 入栈
		stack = append(stack, i)
	}
	for i := 0; i < n; i++ {
		ans = max(ans, heights[i]*(right[i]-left[i]-1))
	}
	return ans
}

func max(x, y int) int {
	if x>y{
		return x
	}
	return y
}
// @lc code=end

