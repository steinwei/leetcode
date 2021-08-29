/*
 * @lc app=leetcode.cn id=1588 lang=golang
 *
 * [1588] 所有奇数长度子数组的和
 */
package leetcode

// @lc code=start
func sumOddLengthSubarrays(arr []int) int {
	ans := 0

	n := len(arr)

	windowSize := 1

	for windowSize <= n {
		slow := 0
		for slow + windowSize <= n {
			for i := slow; i < slow + windowSize; i++ {
				ans += arr[i]
			}
			slow ++
		}
		windowSize +=2
		
	}

	return ans
}
// @lc code=end

