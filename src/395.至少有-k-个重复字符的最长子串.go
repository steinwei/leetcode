/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有 K 个重复字符的最长子串
 */
package leetcode

// @lc code=start
func longestSubstring(s string, k int) int {

	n := len(s)

	ans := 0

	cnt := [26]int{}
	
	for _, v := range s {
		cnt[v-'a']++
	}

	// 找到第一个 分割字符
	var split byte
	for i := 0; i < 26; i++ {
		if cnt[i] < k && cnt[i] > 0 {
			split = byte(i) + 'a'
			break
		}
	}

	
	if split == 0 {
		return n
	}

	l := 0
	r := n-1

	for l <= r {
		// 过滤掉不合适的字符
		for l <= r && s[l] == split {
			l ++
		}
		
		if l > r {
			break
		}

		start := l

		for l <= r && s[l] != split {
			l ++
		}

		ans = max(ans, longestSubstring(s[start:l], k))
	}
	

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
