/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */
package leetcode

// @lc code=start
func updateMatrix(mat [][]int) [][]int {
	n:=len(mat)
	m:= len(mat[0])

	directions := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	stack := [][]int{}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 0 {
				stack = append(stack, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	for len(stack) != 0 {
		point := stack[0]
		stack = stack[1:]

		for i := 0; i < 4; i++ {
			dx := point[0] + directions[i][0]
			dy := point[1] + directions[i][1]

			if dx < 0 || dx > n - 1 || dy < 0 || dy > m - 1 {
				continue
			}

			if mat[dx][dy] == -1 {
				mat[dx][dy] = mat[point[0]][point[1]] + 1
				stack = append(stack, []int{dx, dy})
			}
		}
	}

	return mat
}
// @lc code=end

