/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */
package leetcode

// @lc code=start
func canPartition(nums []int) bool {
	var (
		sum ,avg = 0,0
	)

	for _, v := range nums {
		sum+=v
	}

	if sum & 0x01 == 1 {
		return false
	}
	avg = sum >> 1

	dp:= make([]bool, avg+1)
	dp[0] = true
	for _, v := range nums {
		for i := avg; i - v >= 0; i-- {
			newNode := i
			oldNode := i - v
			if dp[oldNode] {
				dp[newNode] = true
			}
		}
	}

	return dp[avg]
}

// @lc code=end

