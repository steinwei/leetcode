"use strict";
function quicksort(nums) {
    let low = 0, high = nums.length - 1;
    const swap = (l, r) => {
        let temp = 0;
        temp = l;
        l = r;
        r = temp;
    };
    const partition = (l, r, arr) => {
        const random = l + Math.floor(Math.random() * r);
        let prev = random + 1;
        for (let i = l + 1; i < r; i++) {
            if (arr[i] > arr[prev]) {
                swap(i, prev);
            }
        }
        swap(l, prev);
    };
    const mid = low + (high - low) / 2;
    partition(low, mid - 1, nums);
    partition(mid, high, nums);
}
function test() {
    let nums = [2, 3, 5, 1, 4, 6];
    return quicksort(nums);
}
//# sourceMappingURL=0x00-quicksort.js.map