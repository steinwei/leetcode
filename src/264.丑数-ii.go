/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */
package leetcode

// @lc code=start
func nthUglyNumber(n int) int {
	var (
		f = make([]int, n+1)
		min func(x, y int) int
		p1, p2, p3 = 1, 1, 1
	)

	min = func(x, y int) int {
		if x<y {
			return x
		}
		return y
	}

	// f(i) 其中 i 代表 n，结果存储对应的丑数
	f[1] = 1
	for i := 2; i <= n; i++ {
		// 这里 dp 作为记忆缓存，存数据
		x1, x2, x3 := f[p1] * 2, f[p2] * 3, f[p3] * 5
		f[i] = min(x1, min(x2, x3))

		if f[i] == x1 {
			p1 ++
		}

		if f[i] == x2 {
			p2 ++
		}

		if f[i] == x3 {
			p3 ++
		}
	}

	return f[n]
}
// @lc code=end

