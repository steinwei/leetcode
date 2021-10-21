/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */
package leetcode

// @lc code=start
func plusOne(digits []int) (ans []int) {
	n := len(digits)

	flag := 1

	for i := n-1; i >= 0; i-- {
		if flag > 0 {
			digits[i] += flag
			digits[i] %= 10
			if digits[i] == 0 {
				flag = 1
			} else {
				flag = 0
			}
		}
	}

	if flag == 1 {
		digits[0] = 1
		digits = append(digits, 0)
	}

	for _, v := range digits {
		ans = append(ans, v)
	}

	return
}
// @lc code=end

