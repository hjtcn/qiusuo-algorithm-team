<?php

/*

【第四周】：2020-10-12 周一

*/

/*
 * @lc app=leetcode.cn id=876 lang=php
 *
 * [876] 链表的中间结点
 *
 * https://leetcode-cn.com/problems/middle-of-the-linked-list/description/
 *
 * algorithms
 * Easy (69.47%)
 * Likes:    267
 * Dislikes: 0
 * Total Accepted:    77.1K
 * Total Submissions: 110.9K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给定一个带有头结点 head 的非空单链表，返回链表的中间结点。
 * 
 * 如果有两个中间结点，则返回第二个中间结点。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：[1,2,3,4,5]
 * 输出：此列表中的结点 3 (序列化形式：[3,4,5])
 * 返回的结点值为 3 。 (测评系统对该结点序列化表述是 [3,4,5])。
 * 注意，我们返回了一个 ListNode 类型的对象 ans，这样：
 * ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, 以及 ans.next.next.next
 * = NULL.
 * 
 * 
 * 示例 2：
 * 
 * 输入：[1,2,3,4,5,6]
 * 输出：此列表中的结点 4 (序列化形式：[4,5,6])
 * 由于该列表有两个中间结点，值分别为 3 和 4，我们返回第二个结点。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 给定链表的结点数介于 1 和 100 之间。
 * 
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
    // 单指针遍历法
    // 执行用时：4 ms, 在所有 PHP 提交中击败了94.55%的用户
    // 内存消耗：14.5 MB, 在所有 PHP 提交中击败了55.56%的用户
    // 时间复杂度：O(N)，其中 N 是给定链表的结点数目
    // 空间复杂度：O(1)，只需要常数空间存放变量和指针。
    function middleNode1($head) {
        $p1 = $head;
        $len= 0;// 统计链表长度
        $mid= 0;// 中间节点位置
        while($p1 != null) {
            $p1 = $p1->next;
            $len++;
        }
        // 判断链表长度奇偶性 并获取中间节点的位置
        /*
        if($len%2 == 0){
            $mid = $len/2 + 1;
        }else{
            $mid = ($len+1)/2;
        }
        $mid = $mid - 1;
        */
        // intval — 获取变量的整数值
        $mid = intval($len/2); // 获取中间位置
        while($mid) {
            $mid--;
            $head = $head->next;
        }
        return $head;
    }

    // 快慢指针法
    // 执行用时：8 ms, 在所有 PHP 提交中击败了61.82%的用户
    // 内存消耗：14.2 MB, 在所有 PHP 提交中击败了97.22%的用户
    function middleNode2($head) {
        $slow = $head;
        $fast = $head;
        // while($fast && $fast->next) {
        while($fast->next) {
            $slow = $slow->next;       // 慢指针一次走一步
            // 当然可能fast会走到null,但是我们最后只需要slow就行
            $fast = $fast->next->next; // 快指针一次走两步
        }
        return $slow;
    }


}
// @lc code=end


// @tacks think=start
/*
 

题意 meaning 

    也就是想找到链表的中间位置的，并用链表指针进行返回。
    如果有两个中间结点，则返回第二个中间结点。这个也要注意！！！

关键 key 

    “奇偶性”
    “链表的遍历”
    “如何找到中间结点的位置”
    
想法 idea
 
【1】单指针遍历法
    - 第一次遍历记录链表的长度（因为链表结构特性，不能像数组一样访问下标）
    - 通过链表长度len获取中间节点的位置
    - 第二次遍历遍历到中间节点位置即可
 时间复杂度：O(N)
 空间复杂度：O(1)
【2】快慢指针法
    - 快指针一次走两步
    - 慢指针一次走一步
    - 等差数列，那么遍历到最后，慢指针就是走到中间的位置
    - 但是要注意遍历的条件
 时间复杂度：O(N)
 空间复杂度：O(1)


补充 replenish

    小马姐对我的评论让我意识到了，原来其实可以不分奇偶性

    采用  intval 函数 — 获取变量的整数值
        $mid = intval($len/2); // 获取中间位置
    比如 [1,2,3,4] 中间位置就是 intval(len/2) = 4/2 =  2
    比如 [1,2,3]   中间位置就是 intval(len/2) = intval(3/2) = 1

    因为PHP默认/除法是会有小数的，3/2=1.5 弱语言类型。
    
*/
// @tacks think=end