<?php

/*

【第四周】：2020-10-13 周二 

北京: 降温 大风 (记得加衣服)

*/

/*
 * @lc app=leetcode.cn id=206 lang=php
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 * [剑指 Offer 24. 反转链表](https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/)
 * algorithms
 * Easy (70.69%)
 * Likes:    1271
 * Dislikes: 0
 * Total Accepted:    348.3K
 * Total Submissions: 492.7K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 * 
 * 示例:
 * 
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 * 
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
 * 
 */

// @lc code=start
/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val) { $this->val = $val; }
 * }
 */
class Solution {

    /**
     * @param ListNode $head
     * @return ListNode
     */
    // 迭代（遍历）
    // 时间复杂度：O(n)
    // 空间复杂度：O(1)
    // 执行用时：12 ms, 在所有 PHP 提交中击败了44.52%的用户
    // 内存消耗：15.9 MB, 在所有 PHP 提交中击败了65.77%的用户
    function reverseList1($head) {
        $pre = null; // 最后要反转链表
        while($head != null) {
            $next = $head->next; // 提前保存下一个结点
            $head->next = $pre; 
            $pre = $head;
            $head= $next; // 继续遍历下一个结点 
        }
        return $pre;
    }

    // 递归
    // 时间复杂度：O(n)
    // 空间复杂度：O(n) 由于使用递归，将会使用隐式栈空间。递归深度可能会达到 n 层
    // 执行用时：8 ms, 在所有 PHP 提交中击败了83.87%的用户
    // 内存消耗：16.8 MB, 在所有 PHP 提交中击败了5.41%的用户
    function reverseList($head) {
        // 防止链表中循环
        if ($head == null || $head->next == null) {
            return $head;
        }
        // 获得前面反转过的链表
        $pre = $this->reverseList($head->next);
        // 这里比较绕，建议看官方解法
        $head->next->next = $head;
        $head->next = null;
        return $pre;
    } 
}
// @lc code=end


// @tacks think=start
/*
题意 meaning 
  
    反转链表 迭代或递归

关键 key 

    “链表的迭代”
    “链表的递归”

想法 idea

【1】迭代法
    // 看代码吧 画图比较容易理解
    - 主要是记得提前保存下一个结点
【2】递归法
    如果链表如下
        N(1)->N(2)->N(k-1)->N(k)->N(k+1)->....->N(M)
    假设N(K+1) 到N(M) 已经反转成功
        N(1)->N(2)->N(k-1)->N(k)->N(k+1)<-.....<-N(M)

    那么我们希望N(k+1)的下面指针指向指向N(K)
    那么我们需要【N(K)->next->next】 = 【N(k)】
 
*/
// @tacks think=end