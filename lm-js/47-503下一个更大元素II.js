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
 * @lc app=leetcode.cn id=503 lang=javascript
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number[]}
 */
var nextGreaterElements = function(nums) {
    //首先一直目标数组长度为len，先填充为-1
    let len=nums.length,stack=[],res=new Array(len).fill(-1)
    //未查到结果，可以从头循环开始查找，扩容nums数组
    for(let i=len;i<len*2;i++){
        nums.push(nums[i%len])
    }
    //单调栈，遍历扩容后的数组
    for(let i=0;i<nums.length;i++){
        //循环判断当前元素大于栈顶元素，出栈。
        while(stack.length&&nums[stack[stack.length-1]]<nums[i]){
            let index=stack.pop()    
            //这个判断是因为nums数组扩容了嘛！因此我们的目标数组只赋值前一段，防止被重复赋值    
            if(index<len)
            {
                res[index]=nums[i]
            }
        }
        //栈中存的是索引，索引！！！！
        stack.push(i)
 
    }
    return res
};
// @lc code=end

/** 题解
    单调栈。原理还是和前面做的更大元素一致。。
    利用栈存储元素的索引，当当前元素比栈顶元素大，则弹出栈顶元素(索引index)，并直接进行res赋值.res[index]=nums[i]
    但是由于可循环查询，开始前我就将nums数组2倍扩容，而在赋值时却要注意，nums是2倍扩容后的，我们只赋值前一段，防止重复赋值
复杂度分析：
  时间复杂度：O(N)
  遍历一遍nums数组，内部虽然包含while循环，但元素的入栈出栈操作只有一次

  空间复杂度：O(n)
  变量声明目标数组
*/