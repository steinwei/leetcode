/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */
package leetcode

// @lc code=start
func largestRectangleArea(heights []int) int {
	var (
		n = len(heights)
		stack = []int{}
		left = make([]int, n)
		right = make([]int, n)
		maxn = 0
	)
	
	// 填补左哨兵
	for i := 0; i < n; i++ {
		for len(stack)>0 && heights[stack[len(stack)-1]]>=heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	// clear
	stack = []int{}

	// 填补右哨兵
	for i := n-1; i > -1; i-- {
		for len(stack)>0 && heights[stack[len(stack)-1]]>=heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		// 减1是因为前面哨兵的left索引事前-1了
		maxn = max(maxn, heights[i] * (right[i]-left[i]-1))
	}

	return maxn
}

func max(x, y int) int {
	if x>y{
		return x
	}
	return y
}
// @lc code=end

