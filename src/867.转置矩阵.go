/*
 * @lc app=leetcode.cn id=867 lang=golang
 *
 * [867] 转置矩阵
 */
package leetcode

// @lc code=start
func transpose(A [][]int) (ans [][]int) {
	if len(A) == 0 {
		return ans
	}
	m, n := len(A), len(A[0])

	for i := 0; i < n; i++ {
		temp := []int{}
		for j := 0; j < m; j++ {
			temp = append(temp, A[j][i])
		}
		ans = append(ans, temp)
	}

	return ans
}

// @lc code=end
