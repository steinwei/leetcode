/*
 * @lc app=leetcode.cn id=973 lang=golang
 *
 * [973] 最接近原点的 K 个点
 */
package leetcode

import "sort"

// @lc code=start
func kClosest(points [][]int, K int) (ans [][]int) {
	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	})
	ans = points[:K]

	return
}

// @lc code=end
