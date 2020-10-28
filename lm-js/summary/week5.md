# 第五周周报
## 20201019-20201025
### 本周主要刷了五道栈相关的算法题，如下：


### 本周题目

- [20] [有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)

- [739] [每日温度](https://leetcode-cn.com/problems/daily-temperatures/)

- [剑指offer09] [用两个栈实现队列](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/)

- [402] [移掉K位数字](https://leetcode-cn.com/problems/remove-k-digits/)

- [496] [下一个更大元素I](https://leetcode-cn.com/problems/next-greater-element-i/)


### 我的题解
1.[有效的括号](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/36-20%E6%9C%89%E6%95%88%E7%9A%84%E6%8B%AC%E5%8F%B7.js)

 - 利用数组模拟栈的操作，去匹配左右括号
 - 若左括号入栈，若右括号和栈顶元素不是相匹配，则返回false,若匹配，且则pop掉栈顶元素，迭代完毕后栈为空，则返回true,返回咋faflse
 - 学到了利用map进行括号匹配的技巧。以右括号为key，左括号为value，判断右括号和栈顶元素是否匹配。

2.[每日温度](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/37-739%E6%AF%8F%E6%97%A5%E6%B8%A9%E5%BA%A6.js)

 1. 暴力解决。两层for循环，查询到数组最后，再回过头查询下一个。
 2. 单调栈。利用栈存储数组的索引，循环比较如果当前元素比栈顶元素(索引)所指向的元素大，则pop栈顶元素topIdx(索引)，并赋值res[topIdx]等于索引差，直到栈空。

3.[用两个栈实现队列](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/38-%E5%89%91%E6%8C%87offer09%E7%94%A8%E4%B8%A4%E4%B8%AA%E6%A0%88%E6%9E%84%E5%BB%BA%E9%98%9F%E5%88%97.js)

 - 刚开始不了解输入输出，思路是非常散的，后来看了题解，了解到使用js去创建原型对象，和原型方法。
 - 使用两个栈实现队列，栈的特点是什么？先进后出。则只能使用数组的push和pop方法。而使用两个栈模拟队列的先进先出特点。则需要在pop的时候考虑如何翻转栈。
 - 思路：stack1用来添加，stack2用来删除，删除的时候。如果不为空，则直接返回stack2的栈顶元素；如果stack2为空，则不断的pop掉stack1，且stack2不断的push该值，最后stack2弹出栈顶元素，若不存在，返回-1
   
4.[移掉K位数字](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/39-402%E7%A7%BB%E6%8E%89K%E4%BD%8D%E6%95%B0%E5%AD%97.js)

 - 贪心。循环判断如果当前元素比栈顶元素小，则pop栈顶元素，同时k--，试图保持栈从小到大的排列，除非k已经减至0
 - 跳出循环后，k如果没有删除完指定数量，则开始弹出栈顶元素
 - 可以跳出循环后，去除前导零，也可以在遍历元素的时候，栈为空，当前元素为0时，不放入栈中。
 - 拼接字符串返回,栈为空则返回'0'
  
5.[下一个更大元素 I](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/40-496%E4%B8%8B%E4%B8%80%E4%B8%AA%E6%9B%B4%E5%A4%A7%E5%85%83%E7%B4%A0.js)

- 单调栈，思路几乎和每日温度一样，只是最后这道题记录的目标数组记录的是元素，因此会需要map来映射当前元素和下一个比它大的元素。
- 利用栈存储数组的索引，循环比较如果当前元素比栈顶元素(索引)所指向的元素大，则pop栈顶元素topIdx(索引)，并利用map存储(栈顶元素(索引)所指向的元素，当前元素)，直到栈空。



### 一周总结

   这周是栈的题目。大部分都是考验来栈的push，pop,翻转，排列等操作。

   1. 第一次真正刷栈的相关题，感觉就是有思路来，就肯定能刷出来，没思路就果断放弃吧。
   2. 像第三题用两个栈模拟队列，当时输入输出都没看懂，都不知道从何做起。
   3. 像单调栈这种思路，不存元素本身，开始存索引了，不遇到这些题，很少有那个思维扩展往这上面思考。
   4. 借用栈。进行比较、翻转操作，这种方式也是值得积累经验的
    
### 除了坚持，别无它想
