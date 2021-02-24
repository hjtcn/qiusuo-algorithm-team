// 503. 下一个更大元素 II
// 给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

// 示例 1:

// 输入: [1,2,1]
// 输出: [2,-1,2]
// 解释: 第一个 1 的下一个更大的数是 2；
// 数字 2 找不到下一个更大的数； 
// 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
// 注意: 输入数组的长度不会超过 10000。


/*
    好吧，我承认我又忘了代码模版了。。。。。
    害，只能顺着记忆走逻辑：1.循环，需要把原数组重复一遍扩大一倍
                        2.栈记录索引。单调递增栈
    出栈的索引就是目标的索引，它指向的当前值是第一个比它大的值。我把这个部分给忘了。
    首先当前值不断扫描是作为目标值-----单调递增栈记录的是目标索引。
    循环找到一个当前元素大于栈顶元素:栈pop(),目标数组也可以开始赋值。

    还要多做几遍呀
*/

/**
 * @param {number[]} nums
 * @return {number[]}
 */
var nextGreaterElements = function(nums) {
    let len=nums.length;
    nums=nums.concat(nums)
    let stack=[],res=new Array(len).fill(-1)
    for(let i=0;i<nums.length;i++){     
        while(stack.length&&nums[i]>nums[stack[stack.length-1]]){
            let index=stack.pop()
            if(index<len){
                res[index]=nums[i]
            }
        }
        stack.push(i)
    }
    return res
};