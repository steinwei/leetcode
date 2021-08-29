/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
package leetcode

// @lc code=start
var phoneMap map[string]string = map[string]string{
    "2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
}

func letterCombinations(digits string) (ans []string) {

	n := len(digits)

	if n == 0 {
		return
	}

	var backtracks func(index int, temp string)
	backtracks = func(index int, temp string) {
		if index >= n {
			ans = append(ans, temp)
			return
		}

        digit := string(digits[index])
        letters := phoneMap[digit]
        lettersCount := len(letters)
        for i := 0; i < lettersCount; i++ {
            backtracks(index + 1, temp + string(letters[i]))
        }
	}

	backtracks(0, "")

	return
}
// @lc code=end

