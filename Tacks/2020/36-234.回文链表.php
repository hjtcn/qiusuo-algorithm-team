<?php

/*
    2020-10-19 周一 ☀

    上周有点忙，积累了三天的题没刷，抓紧时间补补

    这个是上周四的题
*/

/*
 * @lc app=leetcode.cn id=234 lang=php
 *
 * [234] 回文链表
 *
 * https://leetcode-cn.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (43.50%)
 * Likes:    667
 * Dislikes: 0
 * Total Accepted:    135.8K
 * Total Submissions: 312K
 * Testcase Example:  '[1,2]'
 *
 * 请判断一个链表是否为回文链表。
 * 
 * 示例 1:
 * 
 * 输入: 1->2
 * 输出: false
 * 
 * 示例 2:
 * 
 * 输入: 1->2->2->1
 * 输出: true
 * 
 * 
 * 进阶：
 * 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
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
     * @return Boolean
     */
    // PHP SPL标准库 SplStack()
    // 执行用时：24 ms, 在所有 PHP 提交中击败了84.78%的用户
    // 内存消耗：29.9 MB, 在所有 PHP 提交中击败了13.34%的用户
    function isPalindrome1($head) {
        $p = $head;
        // SplStack类通过使用一个双向链表来提供栈的主要功能
        $stack = new SplStack();
        while($p) {
            $stack->push($p->val); // 压入栈
            $p = $p->next;
        }
        while(!$stack->isEmpty()) {
            // 弹出栈顶元素，与当前head值进行比较，是否相同
            if($stack->pop() != $head->val) {
                return false;
            }
            $head = $head->next;
        }
        return true;
    }

    // 快慢指针
    // 执行用时：28 ms, 在所有 PHP 提交中击败了68.48%的用户
    // 内存消耗：24.3 MB, 在所有 PHP 提交中击败了60.00%的用户
    function isPalindrome2($head) {
        // 边界判断
        if($head == null || $head->next == null) {
            return true;
        }
        $fast = $head;// 快指针
        $slow = $head;// 慢指针
        $part = null; // 目标链表前半部分倒序链表
        // slow遍历前半部分
        while($fast !=null && $fast->next != null) {
            $pre = $slow;
            $slow= $slow->next;       // 慢指针一次走一步
            $fast= $fast->next->next; // 快指针一次走两步
            // 头插法
            $pre->next = $part;
            $part = $pre;
        }
        // 中间位置
        // 如果链表是偶数 fast==null
        // 如果链表是奇数 fast指向最后一个节点，slow刚好是中间节点，在向下移动一个
        if($fast) {
            $slow = $slow->next;
        }
        // slow遍历后半部分
        while($slow || $part) {
            if($slow->val != $part->val) {
                return false;
            }
            $slow = $slow->next;
            $part = $part->next;
        }
        return true;
    }

    // 数组化
    // 执行用时：24 ms, 在所有 PHP 提交中击败了84.78%的用户
    // 内存消耗：27.4 MB, 在所有 PHP 提交中击败了17.34%的用户
    function isPalindrome3($head) {
        $list = [];
        while($head) {
            $list[] = $head->val;
            $head   = $head->next;
        }
        return $list == array_reverse($list);
    }

}
// @lc code=end



// @tacks think=start
/*
题意 meaning 

    链表是否是回文链表

关键 key 

    “回文”
    “链表”
    “快慢指针”
    “栈”

想法 idea

【1】采用内置栈解决
    - SplStack类通过使用一个双向链表来提供栈的主要功能
    - 遍历链表，将节点压入栈
    - 再次遍历链表，弹出栈顶元素进行表比较是否相等
    - 时间复杂度O(n)
    - 空间复杂度O(n)
【2】采用快慢链表
    - 通过快慢指针定位到中间节点的位置，同时得到一半的链表
    - 然后再一起遍历后半部分，判断是否相等
    - 时间复杂度O(n)
    - 空间复杂度O(1)
【3】链表数组化
    - 遍历链表将其赋值到数组中
    - 采用array_reverse数组反转进行判断
 
*/
// @tacks think=end