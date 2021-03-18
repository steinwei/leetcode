/*
 * @lc app=leetcode.cn id=1027 lang=golang
 *
 * [1027] 最长等差数列
 */
package leetcode

// @lc code=start
func longestArithSeqLength(A []int) int {

	n := len(A)

	if n < 3 {
		return 0
	}

	dp := make([]map[int]int, n+1)
	ans := 0

	for i := 0; i < len(dp); i++ {
		dp[i] = make(map[int]int)
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := A[i] - A[j]
			if _, ok := dp[j][diff]; ok {
				dp[i][diff] = dp[j][diff] + 1
			} else {
				dp[j][diff] = 1
				dp[i][diff] = dp[j][diff] + 1
			}
			// fmt.Println(dp)
			ans = max(ans, dp[i][diff])
		}
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
