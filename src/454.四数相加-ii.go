/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */
package leetcode

// @lc code=start
func fourSumCount(A []int, B []int, C []int, D []int) int {
	ans := 0

	countAB := map[int]int{}

	for _, i := range A {
		for _, j := range B {
			countAB[i+j]++
		}
	}

	for _, i := range C {
		for _, j := range D {
			ans += countAB[-i-j]
		}
	}

	// funny

	return ans
}

// @lc code=end
