/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
package leetcode

// @lc code=start
func isValid(s string) bool {
	// var (
	// 	n = len(s)
	// 	stack = []byte{s[0]}
	// 	hashmap = map[byte]byte{
	// 		')': '(',
	// 		']': '[',
	// 		'}': '{',
	// 	}
	// )
	
	// if n & 1 == 1 {
	// 	return false
	// }

	// for i := 1; i < n; i++ {
	// 	cur := s[i]
		
	// 	if len(stack) == 0 {
	// 		stack = append(stack, cur)
	// 		continue
	// 	}

	// 	top := stack[len(stack)-1]
	// 	if p, ok := hashmap[cur]; ok && p == top {
	// 		stack = stack[:len(stack)-1]
	// 		continue
	// 	} else if ok && p != top {
	// 		return false
	// 	}
	// 	stack = append(stack, cur)
	// }

	// return len(stack) == 0

	var (
		stack = []string{}
		n = len(s)
		pairs = map[string]string{
			")": "(",
			"]": "[",
			"}": "{",
		}
	)

	// 每次都漏了判断是否奇偶的条件
	if n&1 == 1 {
		return false
	}

	for i := 0; i < n; i++ {
		// 这里应该要加多判断条件是否匹配括号来减少匹配遍历次数，具体参考上面的答案
		if len(stack)>0&&pairs[string(s[i])] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(s[i]))
		}
	}

	return len(stack) == 0
}
// @lc code=end

