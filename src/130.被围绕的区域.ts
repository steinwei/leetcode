/*
 * @lc app=leetcode.cn id=130 lang=typescript
 *
 * [130] 被围绕的区域
 */

// @lc code=start
/**
 Do not return anything, modify board in-place instead.
 */
function solve(board: string[][]): void {
    const n = board.length

    if(n == 0) {
        return
    }

    const m = board[0].length

    for (let i = 0; i < n; i++) {
        dfs(board, i, 0)
        dfs(board, i, m - 1)
    }

    for (let i = 0; i < m; i++) {
        dfs(board, 0, i)
        dfs(board, n - 1, i)
    }

    for (let i = 0; i < n; i++) {
        for (let j = 0; j < m; j++) {
            if(board[i][j] == 'A'){
                board[i][j] = 'O'
            }else if(board[i][j] == 'O'){
                board[i][j] = 'X'
            }
        }
    }


    function dfs(dp: string[][], x: number, y: number){
        if(x < 0 || x >= n || y<0 || y>=m || dp[x][y] !== 'O'){
            return
        }
        dp[x][y] = 'A'
        dfs(dp, x + 1, y)
        dfs(dp, x - 1, y)
        dfs(dp, x, y + 1)
        dfs(dp, x, y - 1)
    }
};
// @lc code=end

