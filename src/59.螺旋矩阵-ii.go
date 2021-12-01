/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */
package leetcode

// @lc code=start
func generateMatrix(n int) [][]int {
	var (
		total = n * n
		row, col = 0, 0
		matrix = make([][]int, n)
		directions = [][2]int {{0,1},{1,0},{0,-1},{-1,0}}
		directionIndex = 0
	)

	// 建 matrix
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// 
	for i := 1; i < total+1; i++ {
		// ans
		matrix[row][col] = i
		
		dx := row + directions[directionIndex][0]
		dy := col + directions[directionIndex][1]
		
		if dx < 0 || dx >= n || dy < 0 || dy >= n || matrix[dx][dy] > 0 {
			directionIndex = (directionIndex+1)%4
		}
		
		row += directions[directionIndex][0]
		col += directions[directionIndex][1]
	}


	return matrix
}
// @lc code=end

