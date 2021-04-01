/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */
package leetcode

import "strings"

// @lc code=start
func wordBreak(s string, wordDict []string) (sentences []string) {
	wordSet := map[string]struct{}{}
	// 建立字典索引
	for _, w := range wordDict {
		wordSet[w] = struct{}{}
	}

	n := len(s)
	// memo
	dp := make([][][]string, n)
	var backtrack func(index int) [][]string
	backtrack = func(index int) [][]string {
		// memo字典
		if dp[index] != nil {
			return dp[index]
		}
		// tmp
		wordsList := [][]string{}
		for i := index + 1; i < n; i++ {
			// 取词
			word := s[index:i]
			// 字典是否含有
			if _, has := wordSet[word]; has {
				// 递归
				for _, nextWords := range backtrack(i) {
					// 有点难懂
					// 看完js解法，就懂了
					// 这里等同于 js 的， wordsList = [wordsList, ...nextWords]
					wordsList = append(wordsList, append([]string{word}, nextWords...))
				}
			}
		}
		word := s[index:]
		if _, has := wordSet[word]; has {
			// 更新tmp
			wordsList = append(wordsList, []string{word})
		}
		// 添加memo索引
		dp[index] = wordsList
		return wordsList
	}
	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}
	// 其实思路差不多，不过这里是做了出来，而我是没做出来
	// 跟 js 有点不同的是,这里 return 不用 return value,而是把引用的结果给改了, 自动同步到对应地址上
	return
}

// @lc code=end
