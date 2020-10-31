/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */
package leetcode

// @lc code=start
func generate(numRows int) [][]int {
	ans := [][]int{}

	if numRows == 0 {
		return ans
	}
	ans = append(ans, []int{1})
	// 取巧的做法
	for i := 1; i < numRows; i++ {
		m := []int{0}
		m = append(m, ans[len(ans)-1]...)
		for j := 0; j < i; j++ {
			m[j] = m[j] + m[j+1]
		}
		ans = append(ans, m)
	}

	return ans
}

// @lc code=end
