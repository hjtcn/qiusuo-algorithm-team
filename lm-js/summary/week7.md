# 第4周

## 20201102-20201108

> 本周主要刷的是队列，主要重难点在单调队列和循环队列的实现
### 本周题目

-  [[剑指Offer59]滑动窗口的最大值](https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/)
-  [[622]设计循环队列](https://leetcode-cn.com/problems/design-circular-queue/)
-  [[933]最近的请求次数](https://leetcode-cn.com/problems/number-of-recent-calls/)
-  [[641]设计双端循环队列](https://leetcode-cn.com/problems/design-circular-deque/)
-  [[面试题17.09]第k个数](https://leetcode-cn.com/problems/get-kth-magic-number-lcci/)

### 我的题解

1. [滑动窗口的最大值](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/50-%E5%89%91%E6%8C%87offer59%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3%E7%9A%84%E6%9C%80%E5%A4%A7%E5%80%BC.js)

    - 单调队列，保持队列单调递减。
    - 滑动窗口的前一个值和队列头相等，则说明跳出这个窗口，这个最大值马上要失效了。将队列进行shift操作，更新最大值了。

2. [设计循环队列](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/51-622%E8%AE%BE%E8%AE%A1%E5%BE%AA%E7%8E%AF%E9%98%9F%E5%88%97.js)
    - 插入操作，借助tail指针赋值
      删除操作，直接修改head指针即可。this.head=(this.head+1)%this.len
    - 头尾指针的初始化可分别初始为0,len初始为k+1,这样减少一些赋值和判断

3. [最近的请求次数](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/52-933%E6%9C%80%E8%BF%91%E7%9A%84%E8%AF%B7%E6%B1%82%E6%AC%A1%E6%95%B0.js)
    - 暴力。如果新增的数减去3000大于队列头，则需要从头部开始shift队列，
      然后将当前元素入队。返回队列长度即可
4. [设计双端循环队列](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/53-641%E8%AE%BE%E8%AE%A1%E5%BE%AA%E7%8E%AF%E5%8F%8C%E7%AB%AF%E9%98%9F%E5%88%97.js)

    - 借助头尾指针，判断队列为满或空的临界点
    - 头部增加：this.head-1(求余先减k)，头部删除：this.head+1(求余)
      尾部增加：this.tail+1(求余),尾部删除：this.tail-1(求余先加k)
    - 指针对应下标赋值

5. [第k个数](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/54-%E9%9D%A2%E8%AF%95%E9%A2%981709%E7%AC%ACk%E4%B8%AA%E6%95%B0.js)

    - 动态规划问题，主要是要找到思路，更新目标最小值
    - 选择三个指针dp[p3] * 3, dp[p5] * 5, dp[p7] * 7中的最小值作为当前最小乘积
    - 更新完最小值后，对应的指针后移

### 一周总结

   这周是队列的相关题目。
   1. 对于题目来说，队列的数据结构不是很明显。
   2. 第2题和第4题从构建(双端)循环队列来深入理解队列，其它题则是队列的应用题，做完这五道题，你确实会发现，队列就是一种特殊的链表结构，指针的的运用也是极为的重要。
   3. 单调队列的第一题还是没有自己做出来，做单调栈也已经经历三个类型题了，还会要多思考，开拓思维。
   4. 动态规划题的接触，让我感到我们再慢慢坚持，慢慢进步了，数据结构类型弄完，我们就开始各种应用算法啦，加油！！！
    
### 除了坚持，别无它想



