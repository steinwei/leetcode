/*
 * @lc app=leetcode.cn id=388 lang=golang
 *
 * [388] 文件的最长绝对路径
 */
package leetcode

// @lc code=start
func lengthLongestPath(input string) int {
	var (
		stack = []int{}
		max func(x, y int) int
		n = len(input)
		ans = 0
	)

	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 0; i < n; {
		depth := 1
		for ; i < n && input[i] == '\t'; i++ {
			depth++
		}

		isFile := false
		length := 0
		for ; i < n && input[i] != '\n'; i++ {
			if input[i] == '.' {
				isFile = true
			}
			length ++
		}

		i++

		for len(stack) >= depth {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			length += stack[len(stack)-1] + 1
		}

		if isFile {
			ans = max(ans, length)
		} else {
			stack = append(stack, length)
		}
		
	}

	return ans
}
// @lc code=end

