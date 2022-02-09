/*
 * @lc app=leetcode.cn id=1447 lang=golang
 *
 * [1447] 最简分数
 */
package leetcode

import "strconv"

// @lc code=start
func simplifiedFractions(n int) (ans []string) {
	var (
		helper func(a, b int) int
	)

	// isSimpleEnough?
	helper = func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}

		return b
	}

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if helper(j, i) == 1 {
				ans = append(ans, strconv.Itoa(j) + "/" + strconv.Itoa(i))
			}
		}
	}

	return
}
// @lc code=end

