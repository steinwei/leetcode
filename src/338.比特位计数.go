/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */
package leetcode

// @lc code=start
func countBits(n int) (ans []int) {
	// 动规解法
	ans = make([]int, n + 1)
	for i := 0; i < n + 1; i ++ {
		ans[i] = i & 1 + ans[i>>1]
	}

	return

	// 暴力解法
	// for i := 0; i <= n; i++ {
	// 	cnt := i
	// 	res := 0
	// 	for cnt > 0 {
	// 		res += cnt & 1
	// 		cnt >>= 1
	// 	}
	// 	ans = append(ans, res)
	// }

	// return
}
// @lc code=end

