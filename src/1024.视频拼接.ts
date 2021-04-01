/*
 * @lc app=leetcode.cn id=1024 lang=typescript
 *
 * [1024] 视频拼接
 */

// @lc code=start
function videoStitching(clips: number[][], T: number): number {
    const maxn = new Array<number>(T).fill(0);
    let last = 0, ret = 0, pre = 0;
    for (const clip of clips) {
        if (clip[0] < T) {
            maxn[clip[0]] = Math.max(maxn[clip[0]], clip[1]);
        }
    }
    for (let i = 0; i < T; i++) {
        last = Math.max(last, maxn[i]);
        if (i == last) {
            return -1;
        }
        if (i == pre) {
            ret++;
            pre = last;
        }
    }
    return ret;
};
// @lc code=end

