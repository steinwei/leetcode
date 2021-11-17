/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */
package leetcode

// @lc code=start
func maximalRectangle(matrix [][]byte) int {

	var (
		n = len(matrix)
		stack = []int{}
		left = make([]int, n)
		right = make([]int, n)
		maxn = 0
		heights = make([]int, n)
	)

	for i := range matrix {
		sum:=0
		for _, v := range matrix[i] {
			sum += int(v)
		}
	}

	for i := 0; i < n; i++ {
		for len(stack)>0 && matrix
	}

	return maxn
}

func max(x,y int) int{
	if x > y {
		return x
	}
	return y
}
// @lc code=end

