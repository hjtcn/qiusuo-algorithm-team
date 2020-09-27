# 第二周周报
## 20200921-20200927
### 本周主要刷了六道数组相关的算法题，如下：
### 本周题目
- [1535] [找出数组游戏的赢家](https://leetcode-cn.com/problems/find-the-winner-of-an-array-game)
- [1438] [绝对差不超过限制的最长连续子数组](https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit)
- [1267] [统计参与通信的服务器](https://leetcode-cn.com/problems/count-servers-that-communicate)
- [840] [矩阵中的幻方](https://leetcode-cn.com/problems/magic-squares-in-grid)
- [969] [煎饼排序](https://leetcode-cn.com/problems/pancake-sorting/)
- [1424] [对角线遍历](https://leetcode-cn.com/problems/diagonal-traverse-ii)


### 我的题解
1.[找出数组游戏的赢家](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/8-1535%E6%89%BE%E5%87%BA%E6%95%B0%E7%BB%84%E6%B8%B8%E6%88%8F%E7%9A%84%E8%B5%A2%E5%AE%B6.js)

  - 记录最大值，如果出现比max还大的值就将count置为1，重新计数，出现count等于k，就返回最大值。

2.[绝对差不超过限制的最长连续子数组](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/9-1438%E7%BB%9D%E5%AF%B9%E5%B7%AE%E4%B8%8D%E8%B6%85%E8%BF%87%E9%99%90%E5%88%B6%E7%9A%84%E6%9C%80%E9%95%BF%E8%BF%9E%E7%BB%AD%E5%AD%90%E6%95%B0%E7%BB%84.js)

  - 采用双指针，滑动窗口的方式。不断的更新min数组和max数组，获取最大值和最小值之间的绝对差是否大于limit，如果大于limit，则left++，如果小于limit，则right++.
  - 如何维护min数组和max数组呢，如果出现一个比原有min还小的值，说明原有数组不可能成为最小值了。则min.pop()掉这个第二小的值，push最新的最小值。
  而在循环查询最大值最小值的绝对差是否大于limit时，如果当前最小元素等于min的第一个值，就将第一个值也给shift掉。
  - 不断更新不超过limit的最大长度。

3.[统计参与通信的服务器](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/10-1267%E7%BB%9F%E8%AE%A1%E5%8F%82%E4%B8%8E%E9%80%9A%E4%BF%A1%E7%9A%84%E6%9C%8D%E5%8A%A1%E5%99%A8.js)

 1. 暴力解决法。记录出现i，j值的个数，最后再次双层遍历。
 如果该元素grid[i][j]的i，j值都只出现了一次，则该服务器无法通信。记录服务器总个数然后减去无法通信的服务器个数
 2. dfs深搜。整行整列搜索当前包含几个服务器,如果该行该列的服务器个数大于1，则该行该列的服务器都能通信，则res加该个数。
   
4.[矩阵中的幻方](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/11-840%E7%9F%A9%E9%98%B5%E4%B8%AD%E7%9A%84%E5%B9%BB%E6%96%B9.js)

 - 双层遍历行列，计算三行三列两斜线的和，判断和都为15，特殊情况，判断各元素中是否有包含大于9或者小于1的数。以及所有元素为5的情况
  
5.[煎饼排序](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/12-969%E7%85%8E%E9%A5%BC%E6%8E%92%E5%BA%8F.js)

 - 该题的重点在于提示:A[i] 是 [1, 2, ..., A.length] 的排列.以及找寻最大值。
   一开始的最大值为数组长度len，然后利用js的indexOf的api，查找当前所在位置n，通过将前n个数翻转，则最大值就跑第一个了，然后一次全部数组的翻转，它就跑最后一位了，剩下的需要排序的数组就剩len-1个了,最大值为len-1,思路如下进行翻转

6.[对角线遍历](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/13-1424%E5%AF%B9%E8%A7%92%E7%BA%BF%E9%81%8D%E5%8E%86.js)
 - 根据对角线元素nums[i][j]的索引之和（i+j）都相等，双层遍历后将和相同的元素放入一个数组。
   这样经过遍历赋值后的res数组就变为期望数组元素顺序排列的多维数组。
   然后利用flat方法进行无穷极降维到一维数组，并且返回


### 一周总结

   这周的题太沉重了，差点背不动。

   1. js-Api
     - 使用率比较高。数组的push(),pop(),shift(),unshift(),indexOf()
     - 最新了解。flat()：将数组降维
   2. 本周很多题目都可以使用暴力解决，但也都有一下相对的小规律，比如查看索引，索引的和，值的和等等，找到规律，然后边记录边更新。
   3. 双指针比较,滑动窗口
     - 不断更新满足条件的最长字符串。
   4. 个人感觉很多情况下深搜、广搜也可以解决，但是没写出来，有点小忧伤，记下来，来日再啃。
    
### 除了坚持，别无它想
