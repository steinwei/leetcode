/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

package leetcode

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	var (
		n = len(s)
		rk = -1
		kmap = map[byte]int{}
		ans = 0
	)

	for i := 0; i < n; i++ {
		if i > 0 {
			delete(kmap, s[i-1])
		}

		for rk+1 < n && kmap[s[rk+1]] == 0  {
			kmap[s[rk+1]] ++
			rk++
		}

		ans = max(ans, rk - i + 1)
	}
	return ans
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end
