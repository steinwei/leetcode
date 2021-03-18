/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 */
package leetcode

// @lc code=start
func sortArrayByParityII(a []int) []int {
	for i, j := 0, 1; i < len(a); i += 2 {
		if a[i]%2 == 1 {
			for a[j]%2 == 1 {
				j += 2
			}
			a[i], a[j] = a[j], a[i]
		}
	}
	return a
}

// @lc code=end
