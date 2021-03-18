/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */
package leetcode

// @lc code=start
func fib(N int) int {
	x, y := 0, 1
	for i := 0; i < N; i++ {
		x, y = y, y+x
		// fmt.Println(x, y, i)
	}
	return x
}

// @lc code=end
