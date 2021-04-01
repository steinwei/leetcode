/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */
package leetcode

import "strings"

// @lc code=start
func wordPattern(pattern string, s string) bool {
	w2ch := map[string]byte{}
	ch2w := map[byte]string{}

	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	for index, word := range words {
		ch := pattern[index]
		// if
		// set map
		if w2ch[word] > 0 && w2ch[word] != ch || ch2w[ch] != "" && ch2w[ch] != word {
			return false
		}
		w2ch[word] = ch
		ch2w[ch] = word
	}

	return true
}

// @lc code=end
