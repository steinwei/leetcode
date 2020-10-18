/*
 * @lc app=leetcode.cn id=52 lang=typescript
 *
 * [52] N皇后 II
 */

// @lc code=start
function totalNQueens(n: number): number {
    let ans = 0;

    const check = (row: number, col: number, columns: number[]) : boolean=> {
        for (let r = 0; r < row; r++) {
            if( columns[r] == col || row - r == Math.abs(columns[r] - col)) {
                return false
            }            
        }
        return true
    }

    const backtracks = (max: number, row: number, columns: number[]) => {
        if(row == max) {
            ans ++;
            return;
        }

        for (let col = 0; col < n; col++) {
            columns[row] = col
            if(check(row, col, columns)){
                backtracks(max, row + 1, columns)
            }
            columns[row] = -1
        }
    }

    backtracks(n, 0, [])

    return ans
};
// @lc code=end

