/*
 * @lc app=leetcode.cn id=289 lang=golang
 *
 * [289] 生命游戏
 */
package leetcode

// @lc code=start
func gameOfLife(board [][]int)  {
	var (
		ans = [][]int{}
		n, m = len(board), len(board[0])
		dirs = [][2]int{{-1,0},{1,0},{0,-1},{0,1},{1,-1},{1,1},{-1,1},{-1,-1}}
	)

	copy(board, ans)

	mid := m/2

	for x := mid; x < n; x++ {
		live := 0
		dead := 0
		for _, v := range dirs {
			dx := x+v[0]
			dy := mid+v[1]
			if board[dx][dy] == 1 {
				live ++
			} else {
				dead ++
			}
		}
	}
	
	return
}
// @lc code=end

