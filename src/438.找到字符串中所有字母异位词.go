/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */
package leetcode

// @lc code=start
func findAnagrams(s string, p string) (ans []int) {

	var (
		n, m = len(s), len(p)
		sCount = [26]int{}
		pCount = [26]int{}
	)

	if n < m {
		return
	}

	// first window
	for i,v := range p {
		pCount[v-'a']++
		sCount[s[i]-'a']++
	}

	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i := range s[:n-m] {
		sCount[s[i]-'a']--
		sCount[s[i+m]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}

	return
}
// @lc code=end

