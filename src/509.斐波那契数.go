/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */
package leetcode

// @lc code=start
// func fib(N int) int {
// dp[n] = dp[n-1] + dp[n-2]

// 	x, y := 0, 1
// 	for i := 0; i < N; i++ {
// 		x, y = y, y+x
// 	}
// 	return x
// }

func fib(N int) int {

	dp:= make([]int, N + 1)

	if N == 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[N]
}

// @lc code=end
