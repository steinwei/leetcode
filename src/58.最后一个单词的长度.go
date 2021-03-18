/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */
package leetcode

// @lc code=start
func lengthOfLastWord(s string) int {
	ans := 0

	len := len(s)
	n := len - 1

	for n >= 0 {
		if s[n] != ' ' {
			ans++
		}
		if s[n] == ' ' && ans != 0 {
			break
		}
		n--
	}

	return ans
}

// @lc code=end
