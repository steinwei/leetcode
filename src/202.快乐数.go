/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */
package leetcode

// @lc code=start
func isHappy(n int) bool {
	var (
		m = map[int]bool{}
		get_sum func(n int) int
	)

	get_sum = func(n int) int {
		sum := 0
		for n > 0 {
			sum += (n%10)*(n%10)
			n /= 10
		}
		return sum
	}

	for ; n != 1 && !m[n]; n, m[n] = get_sum(n), true {}

	return n == 1
}
// @lc code=end

