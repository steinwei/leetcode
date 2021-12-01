/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */
package leetcode

// @lc code=start
func reverseWords(s string) string {
	var (
		stack = []string{}
		n = len(s)
		ans = ""
	)

	for i := 0; i < n; i++ {
		t := s[i]
		if t == ' ' {
			for len(stack) > 0 {
				v := stack[len(stack)-1]
				ans += v
				stack = stack[:len(stack)-1]
			}
			// 
			ans += " "
		} else {
			stack = append(stack, string(t))
		}
	}

	for len(stack) > 0 {
		v := stack[len(stack)-1]
		ans += v
		stack = stack[:len(stack)-1]
	}

	return ans
}
// @lc code=end

