/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */
package leetcode

// @lc code=start
func canPartition(nums []int) bool {
	// n := len(nums)
	s:= 0

	// get sum
	for _, v := range nums {
		s+=v
	}

	// base
	if s&1 == 1 {
		return false
	}

	// 平分
	v := s>>1

	dp := make([]bool, v + 1)
	dp[0] = true

	// 从set取值
	for _, x := range nums {
	for node := v; node - x> 0; node-- {
			newNode := node
			oldNode := node - x
			if dp[oldNode] {
				dp[newNode] =  true
			}
		}
	}

	return dp[v]
}

// @lc code=end

