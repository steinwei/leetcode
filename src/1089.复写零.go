/*
 * @lc app=leetcode.cn id=1089 lang=golang
 *
 * [1089] 复写零
 */
package leetcode

// @lc code=start
func duplicateZeros(arr []int)  {
	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			for j:=n-1;j>i+1;j--{
				arr[j] = arr[j-1]
			}
			if i<n-1{
				arr[i+1]=0
			}			
			i++
		}		
	}

	return
}
// @lc code=end

