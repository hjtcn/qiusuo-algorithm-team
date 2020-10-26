/*
 * @lc app=leetcode.cn id=496 lang=javascript
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var nextGreaterElement = function(nums1, nums2) {
    let stack=[],map=new Map(),res=[]
    //先遍历nums2,意图记录下每个元素所对应的后一个大于它的元素
    for(let i=0;i<nums2.length;i++){
        //当当前元素比栈顶元素(索引)所对应的元素大
        while(stack.length&&nums2[i]>nums2[stack[stack.length-1]]){
            //弹出栈顶元素，map对象以key,value的形式记录栈顶元素(索引)所对应的元素和当前元素
            let index=stack.pop()
            map.set(nums2[index],nums2[i])
        }
        //索引入栈
        stack.push(i)
    }
    //遍历num1,查询每个元素所对应的比它大的值，如果没找到，即为-1，push进入res数组
    for(let i=0;i<nums1.length;i++){
        res.push(map.get(nums1[i])||-1)
    }
    return res
};


/*
    早上为了评论，发现把那道温度题给遗忘了，后来从新整理整理了思路。
    然后就看到了这道题.卡卡卡，五分钟就做出来了，哈哈哈哈
*/
/** 题解
    单调栈。利用栈存储元素的索引，当当前元素比栈顶元素大，则弹出栈顶元素(索引)，并利用map存储(栈顶元素(索引)所指向的元素，当前元素)
    最后遍历num1查询，每个元素对应的下一个更大元素，push进入目标数组，如果不存在，则向目标数组中push(-1)
复杂度分析：
  时间复杂度：O(N)
  遍历一遍nums2数组，内部虽然包含while循环，但元素的入栈出战操作只有一次

  空间复杂度：O(n)
  变量声明数组和map对象
*/
// @lc code=end

