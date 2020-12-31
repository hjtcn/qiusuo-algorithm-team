<?php


/*
    2020-10-19 周一 ☀

    上周有点忙，积累了三天的题没刷，抓紧时间补补

    这个是上周五的题
*/


/*
 * @lc app=leetcode.cn id=82 lang=php
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (49.41%)
 * Likes:    384
 * Dislikes: 0
 * Total Accepted:    70.1K
 * Total Submissions: 141.7K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 * 
 * 示例 1:
 * 
 * 输入: 1->2->3->3->4->4->5
 * 输出: 1->2->5
 * 
 * 
 * 示例 2:
 * 
 * 输入: 1->1->1->2->3
 * 输出: 2->3
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
    // 执行用时：8 ms, 在所有 PHP 提交中击败了90.91%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了7.14%的用户
    function deleteDuplicates($head) {
        $cur = new ListNode(-1); 
        $dummy = $cur; // 虚拟头结点
        while($head != null) {
            $is_repeat = false; // 是否是重复数字
            while($head->val === $head->next->val) {
                $head = $head->next; // 重复的继续向后遍历
                $is_repeat = true;
            }
            // 如果不是重复的 向后遍历赋值给cur
            if(!$is_repeat) {
                // 尾插法
                $cur->next = $head;
                $cur = $cur->next;
            }
            $head = $head->next;
        }
        // 防止最后一个节点是重复的
        $cur->next = null;
        return $dummy->next; // 返回头节点
    }
}
// @lc code=end


// @tacks think=start
/*
题意 meaning 

    删除链表中重复出现的所有节点
    链表已经是排好序的

关键 key 

    “遍历链表”

想法 idea

    遍历链表删除重复的节点
 
*/
// @tacks think=end