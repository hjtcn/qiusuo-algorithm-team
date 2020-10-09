# 第三周周报
## 20200928-20201004
### 本周主要刷了五道链表相关的算法题，如下：


### 本周题目
- [剑指Offer22] [链表中倒数第k个节点](hhttps://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/)
- [2] [两数相加](https://leetcode-cn.com/problems/add-two-numbers/)
- [21] [合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/)
- [445] [两数相加II](https://leetcode-cn.com/problems/add-two-numbers-ii/)
- [剑指Offer52] [两个链表的第一个公共节点](https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/)


### 我的题解
1.[链表中倒数第k个节点](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/15-%E5%89%91%E6%8C%87offer22%E9%93%BE%E8%A1%A8%E4%B8%AD%E5%80%92%E6%95%B0%E7%AC%ACk%E4%B8%AA%E8%8A%82%E7%82%B9.js)

  - 保持链表长度为k，少于k,将当前指针后移，一旦等于k,则把后移当前指针的同时，头指针也要后移，直至链表走完。

2.[两数相加](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/16-2%E4%B8%A4%E6%95%B0%E7%9B%B8%E5%8A%A0.js)

  - 这个题目主要考虑进位以及遍历两个链表的过程。
  - 主要要考虑两条链表长度不一时的用与来控制没有值赋值为0的情况。

3.[合并两个有序链表](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/17-21%E5%90%88%E5%B9%B6%E4%B8%A4%E4%B8%AA%E6%9C%89%E5%BA%8F%E9%93%BE%E8%A1%A8.js)

 1. 遍历。创建一个目标空链表，每次指向最小的链表节点，等到某个链表遍历结束，即刻将另一个链表的剩余节点缀到目标链表中。
 2. 递归。某一条当前指针指向的节点value较小，就移动指针到下一个节点，递归调用合并函数，直到该条链表遍历结束。
   
4.[两数相加II](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/18-445%E4%B8%A4%E6%95%B0%E7%9B%B8%E5%8A%A0II.js)

 - 对于长的链表，在长度差之外这些数字即为求和之后的值(即len之前，和即本身)，然后遍历两个数组，进行求和赋值
 - 借用数组，进行翻转进位，再进行构建链表
  
5.[两个链表的第一个公共节点 ](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/19-160%E7%9B%B8%E4%BA%A4%E9%93%BE%E8%A1%A8.js)

 1.  利用map对象，记录遍历过的节点，设置为true，再次遍历第二个链表时，出现过的节点直接返回
 2.  双指针实现法。设置两条链表的头指针，让其分别走一遍两条链表，最后相交，两个头指针肯定会走到最后一个节点，否则不相交，返回为null



### 一周总结

   这周小涛出的题，题目难度不高，不过都很有质量，赞赞赞。

   1. 这周主要了解链表的数据结构，同时也深入了解了使用js语言是如何实现链表的过程，如何创建一个空链表以及当链表经过遍历已经指向根节点了，需要提前赋值记录一个空对象，指向头节点，最后返回目标链表。
   2. 使用到了map对象存储出现情况
   3. 使用到了递归，进行链表合并
   4. 双指针方法，实现链表长度的控制
   5. 会使用&&，考虑节点是否为对象，是否含有value值和next节点
    
### 除了坚持，别无它想
