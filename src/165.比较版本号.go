/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */
package leetcode

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	// 双指针
	i, j:=0,0
	n,m:=len(version1),len(version2)

	for i<n || j<m {
		x, y := 0,0
		for ; i < n && version1[i] != '.'; i++ {
			x = x*10
		}

		for ; j < m && version2[j] != '.'; j++ {
			y= y*10
		}

		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}

	return 0
}
// @lc code=end

