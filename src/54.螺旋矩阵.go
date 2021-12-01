/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */
package leetcode

// @lc code=start
func spiralOrder(matrix [][]int) (ans []int) {
	if len(matrix) == 0 {
		return
	}

	if len(matrix[0]) == 0 {
		return
	}

	var (
		n, m = len(matrix), len(matrix[0])
		total = n * m
		directions = [][2]int{{0,1}, {1,0}, {0,-1}, {-1,0}}
		directionIndex = 0
		row, col = 0, 0	
	)

	for i := 0; i < total; i++ {
		ans = append(ans, matrix[row][col])
		matrix[row][col] = -101

		dx := row + directions[directionIndex][0]
		dy := col + directions[directionIndex][1]

		if dx < 0 || dx >= n || dy < 0 || dy >= m || matrix[dx][dy] == -101 {
			directionIndex = (directionIndex+1)%4
		}

		row += directions[directionIndex][0]
		col += directions[directionIndex][1]
	}
	
	return
}
// @lc code=end

