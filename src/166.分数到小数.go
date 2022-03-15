/*
 * @lc app=leetcode.cn id=166 lang=golang
 *
 * [166] 分数到小数
 */
package leetcode

import "strconv"

// @lc code=start
func fractionToDecimal(numerator int, denominator int) string {
	s := []byte{}

	if numerator % denominator == 0 {
		return strconv.Itoa(numerator/denominator)
	}

	if numerator < 0 != (denominator < 0) {
		s = append(s, '-')
	}

	// 整数部分
	numerator = abs(numerator)
	denominator = abs(denominator)
	intergerPart := numerator/denominator
	s = append(s, strconv.Itoa(intergerPart)...)
	s = append(s, '.')

	// 小数部分
	indexMap := map[int]int{}
	reminder := numerator%denominator
	for reminder != 0 && indexMap[reminder]==0 {
		indexMap[reminder] = len(s)
		reminder *= 10
		s = append(s, '0'+byte(reminder/denominator))
		reminder = reminder % denominator
	}

	// 循环部分
	if reminder > 0 {
		insertIndex := indexMap[reminder]
		s = append(s[:insertIndex], append([]byte{'('}, s[insertIndex:]...)...)
		s = append(s, ')')
	}

	return string(s)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// @lc code=end

