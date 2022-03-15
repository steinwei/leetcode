/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */
package leetcode

import (
	"strconv"
	"strings"
)

// @lc code=start
func fizzBuzz(n int) (ans []string) {
	for i := 0; i < n; i++ {
		sb := &strings.Builder{}
		cur := i+1
		isThirdMul:= cur % 3 == 0
		isFiveMul:= cur % 5 == 0

		if isThirdMul {
			sb.WriteString("Fizz")
		} 
		if isFiveMul {
			sb.WriteString("Buzz")
		} 
		if sb.Len() == 0 {
			sb.WriteString(strconv.Itoa(cur))
		}
		ans = append(ans, sb.String())
	}

	return
}
// @lc code=end

