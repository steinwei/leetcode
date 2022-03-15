/*
 * @lc app=leetcode.cn id=1380 lang=golang
 *
 * [1380] 矩阵中的幸运数
 */
package leetcode

// @lc code=start
func luckyNumbers (matrix [][]int) (ans []int) {
	var (
		n, m = len(matrix), len(matrix[0])
	)

	for i := 0; i < n; i++ {
		minX := 0
		maxY := i
		isMax := true

		for j := 0; j < m; j++ {
			if matrix[i][j] < matrix[i][minX] {
				minX = j
				maxY = i
			}
		}
		
		for k := 0; k < n; k++ {
			if k == maxY {
				continue
			}
			if matrix[k][minX] > matrix[maxY][minX] {
				isMax = false
			}
		}

		if isMax == true {
			ans = append(ans, matrix[maxY][minX])
		}
	}

	return
}
// @lc code=end

