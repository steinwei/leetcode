/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] 独一无二的出现次数
 */

package leetcode

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	ans := true

	// sort.Ints(arr)

	// n := len(arr)

	// maps := map[int]bool{}
	// count := 1

	// for i := 1; i < n; {
	// 	if arr[i] == arr[i-1] {
	// 		i++
	// 		count++
	// 	} else {
	// 		if _, ok := maps[count]; ok {
	// 			ans = false
	// 			break
	// 		} else {
	// 			maps[count] = true
	// 			i++
	// 			count = 1
	// 		}
	// 	}
	// }
	// fmt.Println(arr, maps)

	n, maps := len(arr), map[int]int{}

	for i := 0; i < n; i++ {
		if _, ok := maps[arr[i]]; ok {
			maps[arr[i]]++
		} else {
			maps[arr[i]] = 1
		}
	}

	res := map[int]bool{}
	for _, val := range maps {
		if _, ok := res[val]; ok {
			ans = false
			break
		} else {
			res[val] = true
		}
	}
	// fmt.Println(res, maps)

	return ans
}

// @lc code=end
