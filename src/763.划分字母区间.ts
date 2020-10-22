/*
 * @lc app=leetcode.cn id=763 lang=typescript
 *
 * [763] 划分字母区间
 */

// @lc code=start
function partitionLabels(S: string): number[] {
    let res :number[] = []
    let max = 0
    let count = 0
    const len = S.length

    const backtracks = (index: number, str: string) => {
        if(count >= len) {
            return
        }

        const lastIndex = str.lastIndexOf(str[index])
        max = Math.max(max, lastIndex)
        // console.log(
        //     lastIndex, max
        // )
        if(index < max) {
            backtracks(index + 1, str)
        } else if (index == max) {
            count += index + 1
            res.push(index + 1)
            // reset
            max = 0
            backtracks(0, str.substr(index + 1))
        }
    }

    backtracks(0, S)

    return res
};
// @lc code=end

