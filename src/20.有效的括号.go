/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
package leetcode

// @lc code=start
func isValid(s string) bool {
	var (
		n = len(s)
		stack = []byte{s[0]}
		hashmap = map[byte]byte{
			')': '(',
			']': '[',
			'}': '{',
		}
	)
	
	if n & 1 == 1 {
		return false
	}

	for i := 1; i < n; i++ {
		cur := s[i]
		
		if len(stack) == 0 {
			stack = append(stack, cur)
			continue
		}

		top := stack[len(stack)-1]
		if p, ok := hashmap[cur]; ok && p == top {
			stack = stack[:len(stack)-1]
			continue
		} else if ok && p != top {
			return false
		}
		stack = append(stack, cur)
	}

	return len(stack) == 0
}
// @lc code=end

