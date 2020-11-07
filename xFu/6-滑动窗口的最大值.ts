// 剑指 Offer 59 - I. 滑动窗口的最大值
// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

// 示例:

// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
// 输出: [3,3,5,5,6,7] 
// 解释: 

//   滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7
 

// 提示：

// 你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

function maxSlidingWindow(nums: number[], k: number): number[] {
    let list= [];
    for(let i=0;i< nums.length;i++) {
        let flag = 0;
        let maxVal = -Infinity;

        // 若当前滑动窗口中的最后一个值不存在，就结束滑动 
        if(nums[i+k-1] === undefined) break;

        // 循环当前窗口中的值，找出最大值
        while(flag < k){
            let currentNum = nums[i+flag];
            if(maxVal < currentNum){
                maxVal = currentNum;
            }
            flag++;
        }

        list.push(maxVal);
        
    }

    return list;
};

// 执行用时：156 m, 在所有 typescript 提交中击败了53.85%的用户
// 内存消耗45.6 M, 在所有 typescript 提交中击败了46.15%的用户

// 总结: 
// 1.这种方法没有使用到队列的知识
// 2.这种方法在 i k flag （窗口起始下标， 窗口长度， 窗口内数据循环下标）的处理上容易绕晕