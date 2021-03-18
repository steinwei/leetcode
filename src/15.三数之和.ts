/*
 * @lc app=leetcode.cn id=15 lang=typescript
 *
 * [15] 三数之和
 */

// @lc code=start
function threeSum(nums: number[]): number[][] {
	const res: Array<Array<number>> = [];

	// sort
	nums.sort((a,b)=>a-b)

	const len:number=nums.length
	const end:number=len-2

	// boundary
	if (nums[0] <= 0 && nums[len - 1] >= 0) {
		for (let i:number = 0; i < end;) {
			// 取负数
			if(nums[i]>0) break
			let head:number = i+1
			let tail:number = len-1
			while(head<tail){
				let first:number=nums[head]
				let second:number=nums[tail]
				const result:number = nums[i]+first+second
				if(result == 0) {
					res.push([nums[i], first, second])
				}
				if(result<=0){
					while(nums[head]==nums[++head]){}
				}else if(result>0){
					while(nums[tail]==nums[--tail]){}
				}else{
					break
				}
			}
			while(nums[i]==nums[++i]){}
		}
	}
	// 双指针

	return res;
}
// @lc code=end
// threeSum([-1, 0, 1, 2, -1, -4])