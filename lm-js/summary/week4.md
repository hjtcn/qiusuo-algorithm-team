# 第三周周报
## 20201012-20201018
### 本周主要刷了五道链表相关的算法题，如下：


### 本周题目

- [876][链表的中间结点](https://leetcode-cn.com/problems/middle-of-the-linked-list/)

- [剑指 Offer 24] [反转链表](https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/)

- [92][反转链表 II](https://leetcode-cn.com/problems/reverse-linked-list-ii/)

- [234][回文链表](https://leetcode-cn.com/problems/palindrome-linked-list/)

- [82][删除排序链表中的重复元素 II](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)

### 我的题解
1.[链表的中间结点](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/29-876%E9%93%BE%E8%A1%A8%E7%9A%84%E4%B8%AD%E9%97%B4%E8%8A%82%E7%82%B9.js)

 1. 暴力解决。遍历获取链表长度，再找中间节点
 2. 快慢指针，快指针遍历整个链表结束，慢指针正好走到中间位置。

2.[反转链表](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/30-%E5%89%91%E6%8C%87offer24%E5%8F%8D%E8%BD%AC%E9%93%BE%E8%A1%A8.js)

 1. 迭代，每次首先将下一个节点存储起来，然后开始反转，最后哦借用提前存起来的下一个节点来进行迭代。
 2. 递归，只关注目标节点和它的下一个节点，然后反转。下一个节点的next指向目标节点，目标节点的next指向null.

3.[反转链表 II](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/31-92%E5%8F%8D%E8%BD%AC%E9%93%BE%E8%A1%A8II.js)

 - 这个是反转链表的深入篇，指定位置的反转链表。
 - 将链表分为三段，头尾不必反转，但要记住头的最后一个节点front和尾部的第一个节点prev。然后中间位置反转，再利用提前记录的front和prev进行链表拼接。
   
4.[回文链表](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/32-234%E5%9B%9E%E6%96%87%E9%93%BE%E8%A1%A8.js)

 - 数组及头尾指针.遍历链表，存储至新数组中，然后头尾指针对比，出现不等的就返回false，直至头尾指针中某一个已到中间位置，依然相等，则返回true
  
5.[删除排序链表中的重复元素 II ](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/33-82%E5%88%A0%E9%99%A4%E6%8E%92%E5%BA%8F%E9%93%BE%E8%A1%A8%E4%B8%AD%E7%9A%84%E9%87%8D%E5%A4%8D%E5%85%83%E7%B4%A0.js)

- HashMap。遍历链表，完成value和出现次数的存储，然后创建新的链表，仅存储出现次数等于1的节点。



### 一周总结

   这周依然是链表的题目。总的感觉就是反转链表确实是有点刺激，总是拐着拐着就晕了。

   1. 没思路时，第一想法就是暴力迭代找答案。
   2. 快慢指针有的时候用起来确实很巧妙，但是全凭自己思考是考虑不到这一点的，需要做的类型题多了，就熟能生巧了。
   3. 使用到了map对象存储出现情况。其实是因为后两天比较忙碌了，就没有深入扩展，就使用了自己比较擅长的方法去实现。
   4. 使用到了递归，递归真的感觉很神奇，思路一定要非常清晰，否则就很容易形成事件环。当然有些时候也容易造成超时。
   5. 反转链表的节点提前存储，反转过程，执行下一个链表的迭代过程，同样是要思路清晰 ，发现链表造成事件环的纪律太大了，我太南了。。
    
### 除了坚持，别无它想
