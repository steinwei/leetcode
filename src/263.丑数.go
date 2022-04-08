/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */
package leetcode

// @lc code=start
func isUgly(n int) bool {
	var (
		factors = []int{2,3,5}
	)

	if n <= 0 {
		return false
	}

	for _, v := range factors {
		for n%v == 0 {
			n/=v
		}
	}

	return n == 1
}
// @lc code=end

