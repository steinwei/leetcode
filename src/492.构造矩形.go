/*
 * @lc app=leetcode.cn id=492 lang=golang
 *
 * [492] 构造矩形
 */
package leetcode

import "math"

// @lc code=start
func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))
    for area%w > 0 {
        w--
    }
    return []int{area / w, w}
}
// @lc code=end

