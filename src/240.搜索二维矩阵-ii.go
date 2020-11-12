/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package leetcode

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	ans := false

	for _, i := range matrix {
		for _, j := range i {
			if ans == true {
				continue
			}
			if j == target {
				ans = true
			}
		}
	}

	return ans
}

// @lc code=end
