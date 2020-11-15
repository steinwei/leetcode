/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 */
package leetcode

import (
	"sort"
)

// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) (ans []int) {

	rank := map[int]int{}

	for i, num := range arr2 {
		rank[num] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		// fmt.Println(i, j, x, y)
		rankX, hasX := rank[x]
		rankY, hasY := rank[y]

		if hasX && hasY {
			return rankY > rankX
		}

		if hasX || hasY {
			return hasX
		}

		return x < y

	})

	ans = append(ans, arr1...)

	return
}

// @lc code=end
