/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 */
package leetcode

// @lc code=start
func validMountainArray(A []int) bool {

	n := len(A)
	i := 0

	for ; i+1 < n && A[i] < A[i+1]; i++ {
	}

	if i == n-1 || i == 0 {
		return false
	}

	for ; i+1 < n && A[i+1] < A[i]; i++ {
	}

	return i == n-1

}

// @lc code=end
