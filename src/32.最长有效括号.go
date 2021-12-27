/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */
package leetcode

// @lc code=start
func longestValidParentheses(s string) int {
	var (
		ans = 0
		n = len(s)
		stack = []int{-1}
	)

	for i := 0; i < n; i++ {
		cur := s[i]
		if cur == '(' {
			stack = append(stack, i)
		} else {
			top := stack[len(stack)-1]
			if top != -1 && s[top] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, i)
			}
			ans = max(ans, i-stack[len(stack)-1])
		}
	}

	return ans
}

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

