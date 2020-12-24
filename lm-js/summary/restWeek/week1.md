# 单调栈
### 定义

栈内元素递增(递减)顺序排列的栈

### 应用

1. 向右(或向左)比它大(或小)的元素的距离

- [739] [每日温度](https://leetcode-cn.com/problems/daily-temperatures/)
- [496] [下一个更大元素I](https://leetcode-cn.com/problems/next-greater-element-i/)
- [503] [下一个更大元素II](https://leetcode-cn.com/problems/next-greater-element-ii/)


2. 矩形面积类型

- [84] [柱状图中最大的矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)

- [42] [接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)



代码模版：
``` 
let stack = [];
    for(let i = 0; i < T.length; i++) {
        // 如果栈长度不为空，且当前元素大于栈顶元素
        while(stack.length && T[i] > T[stack[stack.length - 1]]) {
            //弹出栈顶元素
            let topIdx = stack.pop();
            //记录栈顶元素所对应的结果值
            res[topIdx] = i - topIdx;
        }
        //当前元素入栈
        stack.push(i);
        // console.log(stack,res)
}
``` 

### 心得体会

1. 感觉挺好玩的，但是如果没接触这种题之前是真的没有思路，除非暴力N^2进行查询
2. 多锻炼几遍，熟能生巧
3. 矩形面积难度要更大点，找元素知道怎么暴力，但是面积就更加需要想象能力，面积的更新(索引*高度)，单调栈还是记录索引。

