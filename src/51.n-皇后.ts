/*
 * @lc app=leetcode.cn id=51 lang=typescript
 *
 * [51] N 皇后
 */

// @lc code=start
function solveNQueens(n: number): string[][] {
    let ans: string[][] = []

    const check = (row: number, col: number, columns: number[]): boolean => {
        for (let r = 0; r < row; r++) {
            if(columns[r] == col || row - r == Math.abs(columns[r] - col)) {
                return false
            }            
        }
        return true
    }

    const backtracks = (max: number, row: number, columns: number[]) => {
        if(row == max) {
            ans.push(createTable(n, columns))
            return
        }

        for (let col = 0; col < max; col++) {
            columns[row] = col

            if(check(row, col, columns)){
                backtracks(max, row + 1, columns)
            }

            // columns[row] = -1            
        }
    }

    backtracks(n, 0, [])

    return ans

    function createTable(max: number, columns: number[]): string[] {
        let table: string[] = []
        
        for (let i = 0; i < n; i++) {
            let point = ''
            for (let j = 0; j < n; j++) {
                if(columns[i] == j) {
                    point += 'Q'
                } else {
                    point += '.'
                }
            }
            table.push(point)
        }

        return table
    }

};
// @lc code=end

