/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */
package leetcode

// @lc code=start
func setZeroes(matrix [][]int)  {
	var (
		m,n = len(matrix), len(matrix[0])
		col0 = false
	)

	for _, v := range matrix {
		if v[0] == 0 {
			col0 = true
		}

		for j := 1; j < n; j++ {
			if v[j] == 0 {
				v[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	
	for row := m-1; row >= 0; row-- {
		for col := 1; col < n; col++ {
			if matrix[0][col] == 0 || matrix[row][0] == 0 {
				matrix[row][col] = 0
			}
		}
		if col0 {
			matrix[row][0] = 0
		}
	}

	return
}
// @lc code=end

