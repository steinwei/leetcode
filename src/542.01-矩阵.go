/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */
package leetcode

// @lc code=start
func updateMatrix(mat [][]int) (ans [][]int) {
	// 
	// dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[[i][j+1]]) + 1
	n:=len(mat)

	dp:= make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			}

			if i == 0 || j == 0 || i == n-1 || j == n-1 {
				dp[i][j] = mat[i][j]
				continue
			}

			dp[i][j] = min(min(min(mat[i-1][j], mat[i+1][j]), mat[i][j-1]), mat[i][j+1]) + 1
		}
	}

	ans = dp

	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

