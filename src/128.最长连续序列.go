/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */
package leetcode

// @lc code=start
func longestConsecutive(nums []int) int {
	ans := 0

	seen := map[int]bool{}

	for _, v := range nums {
		seen[v] = true
	}

	for i := range seen {

		if !seen[i-1] {

			curr := i

			for seen[curr+1] {
				curr ++
			}

			if curr - i + 1 > ans {
				ans = curr - i + 1
			}
		}
	}

	return ans
}
// @lc code=end

