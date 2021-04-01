/*
 * @lc app=leetcode.cn id=22 lang=typescript
 *
 * [22] 括号生成
 */

// @lc code=start
function generateParenthesis(n: number): string[] {
    let res: string[] = []

    // baoli
    const backtracks = (open: number, close: number, temp: string, max: number, res: string[]) => {
        if(temp.length >= max * 2) {
            res.push(temp)
            return
        }

        if(open < max) {
            backtracks(open + 1, close, temp + '(', max, res)
        }

        if( close < open) {
            backtracks(open, close + 1, temp + ')', max, res)
        }
    }

    backtracks(0, 0, '', n, res)

    return res
};
// @lc code=end

