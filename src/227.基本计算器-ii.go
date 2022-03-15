/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */
package leetcode

// @lc code=start
func calculate(s string) int {
	var (
		ans = 0
		stack = []int{}
		num = 0
		preSign = '+'
	)

	for i, ch := range s {
		isDigit := ch >= '0' && ch <= '9'
		if isDigit {
			num = num * 10 + int(ch-'0')
		}

		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
				case '+':
					stack = append(stack, num)
				case '-':
					stack = append(stack, -num)
				case '*':
					stack[len(stack)-1] *= num
				case '/':
					stack[len(stack)-1] /= num
				}
			preSign = ch
			num = 0
		}
	}

	for _, v := range stack {
		ans += v
	}

	return ans
}
// @lc code=end

