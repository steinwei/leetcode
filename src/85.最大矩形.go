/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */
package leetcode

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	var (
		n, m = len(matrix), len(matrix[0])
		nums = make([]int, m)
		ans = 0
	)

	// col
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] - '0' == 1 {
				nums[j] += 1
			} else {
				nums[j] = 0
			}
		}
		ans = max(ans, largestRectangleArea(nums))
	}	
		
	return ans
}

func max(x, y int) int {
	if x>y {
		return x
	}
	return y
}

func largestRectangleArea(heights []int) int {
	var (
		n = len(heights)
		l = make([]int, n)
		r = make([]int, n)
		stack = []int{}
		ret = 0
	)
	
	// right
	for i := 0; i < n; i++ {
		for len(stack)>0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			l[i] = -1
		} else {
			l[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	// clear
	stack = []int{}
	// left
	for i := n-1; i > -1; i-- {
		for len(stack)>0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			r[i] = n
		} else {
			r[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		ret = max(ret, heights[i]*(r[i]-l[i]-1))
	}

	return ret
}

// @lc code=end

