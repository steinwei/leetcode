/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */
package leetcode

// @lc code=start
func rotate(matrix [][]int)  {
	n := len(matrix)

	tmp := make([][]int, n)

	for rows := range tmp {
		tmp[rows] = make([]int, n)
	}

	for i := range matrix {
		for j := range matrix[i] {
			tmp[j][n - i - 1] = matrix[i][j]
		}
	}

	copy(matrix, tmp)

	return
}
// @lc code=end

