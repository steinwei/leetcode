/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */
package leetcode

// @lc code=start
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}

	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}

	return true
}

// @lc code=end
