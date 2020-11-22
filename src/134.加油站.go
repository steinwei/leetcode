/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */
package leetcode

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {

	for i, n := 0, len(gas); i < n; i++ {
		dp_0_0 := 0
		sumOfGas, sumOfCost := 0, 0
		for dp_0_0 < n {
			j := (i + dp_0_0) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			// fmt.Println(i, dp_0_0, j)
			if sumOfCost > sumOfGas {
				break
			}
			dp_0_0++
		}

		if dp_0_0 == n {
			// todo
			return i
		}

		// i = dp_0_0
	}

	return -1
}

// @lc code=end
