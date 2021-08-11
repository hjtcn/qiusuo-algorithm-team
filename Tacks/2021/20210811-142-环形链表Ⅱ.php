<?php
/*
 * @Descripttion: 周三干
 * @Author: tacks321@qq.com
 * @Date: 2021-08-11 11:00:24
 * @LastEditTime: 2021-08-11 14:36:02
 */


/*
 * @lc app=leetcode.cn id=142 lang=php
 *
 * [142] 环形链表 II
 *
 * https://leetcode-cn.com/problems/linked-list-cycle-ii/description/
 *
 * algorithms
 * Medium (54.75%)
 * Likes:    1101
 * Dislikes: 0
 * Total Accepted:    269.2K
 * Total Submissions: 488.9K
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
 * 
 * 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是
 * -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
 * 
 * 说明：不允许修改给定的链表。
 * 
 * 进阶：
 * 
 * 
 * 你是否可以使用 O(1) 空间解决此题？
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：head = [3,2,0,-4], pos = 1
 * 输出：返回索引为 1 的链表节点
 * 解释：链表中有一个环，其尾部连接到第二个节点。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：head = [1,2], pos = 0
 * 输出：返回索引为 0 的链表节点
 * 解释：链表中有一个环，其尾部连接到第一个节点。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 
 * 
 * 输入：head = [1], pos = -1
 * 输出：返回 null
 * 解释：链表中没有环。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中节点的数目范围在范围 [0, 10^4] 内
 * -10^5 
 * pos 的值为 -1 或者链表中的一个有效索引
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


/**
 * 快慢双指针
 * 
 * 执行用时：12 ms, 在所有 PHP 提交中击败了88.57%的用户
 * 内存消耗：18.2 MB, 在所有 PHP 提交中击败了70.00%的用户
 * 
 * @date  2021-08-11 14:26:34
 */
class Solution {
    /**
     * @param ListNode $head
     * @return ListNode
     */
    function detectCycle($head) {
        // 快慢指针
        $slow = $head;
        $fast = $head;
        
        while($fast != null && $fast->next !=null ) {
            $slow = $slow->next;
            $fast = $fast->next->next; // 快指针一次走两步
            // 第一次相遇的时候
            if($slow == $fast) {
                break;
            }
        }
        // 如果遍历结束遇到 null， 那么无环
        if($fast == null || $fast->next == null) {
            return null;
        }

        // 如果有环,计算入环点，第二次相遇
        $slow = $slow; // 另一个指针保持在首次相遇点
        $fast = $head; // 一个指针从头跑
        while($fast != $slow) {
            $slow = $slow->next;
            $fast = $fast->next;
        }
        return $fast;
    }
}
// @lc code=end

/*

【 Clarification 】
题目：
    给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
    说明：不允许修改给定的链表。    
含义：
    判断链表是否有环，并求出链表的入环点。

【 Possible-Solution 】
https://pic.leetcode-cn.com/1623750708-inrvVm-1623750459597.jpg

1、快慢指针
    假设 一段路程，
    快慢指针，开始到入环点为D，入环点到首次相遇点为S1, 首次相遇点到入环点为 S2

    指针 slow 每次走一步 所走的距离为 D+S1
    指针 fast 每次走两部 所走的距离为 D+S1+n(S1+S2)

    因为速度关系，fast是slow的两倍

    2(D+S1) = D + S1 + n(S1+S2)

    从而 D = (n-1)(S1+S2) + S2

    也就是说链表头节点开始到入环点的距离，等于首次相遇点绕n-1圈再回到入环点的距离


*/