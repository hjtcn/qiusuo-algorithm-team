# 第六周周报
## 20201026-20201101
### 本周主要刷了五道栈相关的算法题，如下：


### 本周题目

- [155] [最小栈](https://leetcode-cn.com/problems/min-stack/)

- [150] [逆波兰表达式求值](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)

- [844] [比较含退格的字符串](https://leetcode-cn.com/problems/backspace-string-compare/)

- [316] [去除重复字母](https://leetcode-cn.com/problems/remove-duplicate-letters/[[503]])

- [503] [下一个更大元素II](https://leetcode-cn.com/problems/next-greater-element-ii/)

### 我的题解
1.[最小栈](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/43-155%E6%9C%80%E5%B0%8F%E6%A0%88.js)

 1. 利用栈维持连续递减的较小元素。空间复杂度O(N)
 2. 仅设置变量记录最小元素，而执行入栈出栈操作的栈记录原有元素减去最小值的差。当栈顶元素(差)<0时，pop的栈顶元素为差+最小元素，更新最小元素。空间复杂度为O(1)

2.[逆波兰表达式求值](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/44-150%E9%80%86%E6%B3%A2%E5%85%B0%E8%A1%A8%E8%BE%BE%E5%BC%8F%E6%B1%82%E5%80%BC.js)

 - 利用栈的入栈出栈操作，将运算值入栈，遇到运算符，则将左右运算值出栈，计算后的结果入栈，重复操作 

3.[比较含退格的字符串](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/45-844%E6%AF%94%E8%BE%83%E5%90%AB%E9%80%80%E6%A0%BC%E7%9A%84%E5%AD%97%E7%AC%A6%E4%B8%B2.js)

 1. 暴力处理两个字符串，将字符串变为数组遇到#则出栈，最后将数组拼接为字符串对比是否相同
 2. 双指针。两个指针同时指向字符串结尾出，向前移动，当指针遇到'#'时，指针左移，相应的标记(用来记录退格的影响力)++。当退格影响力的标记大于0时，继续不停左移。战胜这两种情况，对比当前指针指向的元素是否相等，不相等则返回false,相等则指针继续同时左移。跳出循环，返回true。
   
4.[去除重复字母](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/46-316%E5%8E%BB%E9%99%A4%E9%87%8D%E5%A4%8D%E5%AD%97%E6%AF%8D.js)

 - 贪心。循环判断当前元素小于栈顶元素，且当前元素后面还会存在栈顶元素相同的元素，则出栈。
  
5.[下一个更大元素 II](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/47-503%E4%B8%8B%E4%B8%80%E4%B8%AA%E6%9B%B4%E5%A4%A7%E5%85%83%E7%B4%A0II.js)

- 单调栈。可以循环查找，则将nums作2倍扩容，利用栈存储数组的索引，循环比较如果当前元素比栈顶元素(索引)所指向的元素大，则pop栈顶元素topIdx(索引)，并赋值目标数组res[topIdx]=nums[i]。注意，赋值过程中topIdx要小于原nums的长度，因为这个数组被2倍扩容过，防止重复赋值。



### 一周总结

   这是第二周刷栈的题目了。

   1. 前三题还是满基础的，第四题贪心算法，基本没啥思路，第五题之前做过类似的两道，也算是重新复习了一把。
   2. 有些简单题暴力做也是可以做出来的，但是做完翻翻题解，也会发现时间空间复杂度更优化一步的处理方法。
   3. 去除重复字母，看完鹏飞学长的题解，觉得没思路还是不能慌，一步步去接近答案，就算只解决部分问题，也会更加深印象。
   4. 单调栈确实有点酷，感觉以后还会遇到不少。
    
### 除了坚持，别无它想
