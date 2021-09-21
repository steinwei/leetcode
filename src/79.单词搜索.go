/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */
package leetcode

// @lc code=start
func exist(board [][]byte, word string) bool {

	ans := false

	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	n, m := len(board), len(board[0])

	vis := make([][]bool, n)

	for i:= range vis {
		vis[i] = make([]bool, m)
	}

	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] !=  word[k] {
			return false
		}
		if k == len(word) - 1 {
			return true
		}

		vis[i][j] = true
		defer func ()  {vis[i][j] = false}()
		for _, v := range directions {
			newX := i + v[0]
			newY := j + v[1]
			if newX >= 0 && newX < n && newY >=0 && newY < m && !vis[newX][newY] {
				if check(newX, newY, k+1) {
					return true
				}
			}
		}

		// vis[i][j] = false

		return false
	}

	for i, rows := range board {
		for j := range rows {
			if check(i, j, 0) {
				return true
			}
		}
	}

	return ans
}
// @lc code=end

