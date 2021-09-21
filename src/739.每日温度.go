/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */
package leetcode

// @lc code=start
func dailyTemperatures(temperatures []int) (ans []int) {
	n := len(temperatures)

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if temperatures[j] - temperatures[i] > 0 {
				dp[i] = j - i
				break
			}
		}
	}

	ans = dp

	return 
}
// @lc code=end

