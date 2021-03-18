/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */
package leetcode

// @lc code=start
func longestPalindrome(s string) int {

	hashmap := map[rune]int{}

	for _, v := range s {
		hashmap[v]++
	}

	count := 0
	for _, v := range hashmap {
		count += v / 2 * 2
		if v&1 == 1 && count&1 == 1 {
			count += v
		}
	}

	if len(s) == count {
		return count
	} else {
		return count + 1
	}
}

// @lc code=end
