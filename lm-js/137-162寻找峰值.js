// 162. 寻找峰值
// 峰值元素是指其值大于左右相邻值的元素。

// 给你一个输入数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

// 你可以假设 nums[-1] = nums[n] = -∞ 。

 

// 示例 1：

// 输入：nums = [1,2,3,1]
// 输出：2
// 解释：3 是峰值元素，你的函数应该返回其索引 2。
// 示例 2：

// 输入：nums = [1,2,1,3,5,6,4]
// 输出：1 或 5 
// 解释：你的函数可以返回索引 1，其峰值元素为 2；
//      或者返回索引 5， 其峰值元素为 6。
 

// 提示：

// 1 <= nums.length <= 1000
// -231 <= nums[i] <= 231 - 1
// 对于所有有效的 i 都有 nums[i] != nums[i + 1]

/*
    思路：这道题讲真的，二分不太会。我没有找到nums[mid]去和谁对比
        看了题解，说根据与nums[mid+1]对比发现趋势来分辨
        如nums[mid]>nums[mid+1],说明是下降趋势，让r=mid
          nums[mid]<nums[mid+1],说明上升趋势，让l=mid
         一直维持峰值位置在l.
         为什么l是峰值的位置，因为最大的情况就是跳出循环前一个操作，还进行了个l=mid+1的赋值，
         r每次更新都是mid，

         按趋势来二分，我自己确实不敢想，怕有遗漏的峰值.后来我自己用测试用例运行了一遍
         [1,2,1,3,4,5,6]
         预期结果为6
         说明，如果最后一位数右边没值，且比左边值大，也可以输出。
         所以不怕遗漏，只要测试范围一直有一个最大值就行。


*/
/*
 * @lc app=leetcode.cn id=162 lang=javascript
 *
 * [162] 寻找峰值
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number}
 */

 var findPeakElement = function(nums) {
    let l=0,r=nums.length-1
    while(l<r){
        let mid=parseInt((l+r)/2)
        if(nums[mid]>nums[mid-1]&&nums[mid]>nums[mid+1]){
            return mid
        }
        if(nums[mid]<nums[mid+1]){
            l=mid+1
        }
        if(nums[mid]>nums[mid+1]){
            r=mid
        }
    }
    return l   
};
// @lc code=end

/*
    时间复杂度：O(logN)
    空间复杂度：O(1)
*/