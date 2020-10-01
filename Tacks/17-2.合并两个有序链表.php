<?php
/*
 * @lc app=leetcode.cn id=21 lang=php
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (64.42%)
 * Likes:    1287
 * Dislikes: 0
 * Total Accepted:    379K
 * Total Submissions: 588.4K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
 */
class ListNode {
    public $val  = 0;
    public $next = null;
    function __construct($val = 0, $next = null) {
        $this->val  = $val;
        $this->next = $next;
    }
}

class Solution {

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    // 迭代1
    // 执行用时：8 ms, 在所有 PHP 提交中击败了86.16%的用户
    // 内存消耗：14.5 MB, 在所有 PHP 提交中击败了52.70%的用户
    function mergeTwoLists1($l1, $l2) {
        $preHead = new ListNode(null); // 声明一个头指针用来等待返回最后得到整个链表。
        $prev    = $preHead; // 这里是引用
        while($l1 !== null && $l2 !== null) {
            if($l1->val <= $l2->val) {
                $prev->next = $l1; // prev向后移动,向较小的值移动
                $l1 = $l1->next;
            }else{
                $prev->next = $l2; 
                $l2 = $l2->next;
            }
            $prev = $prev->next; // 当前节点也需要向后移动
        }
        // 如果两个链表中还有一个没遍历完，直接加到后面
        $prev->next = ($l1 === null ? $l2 : $l1);

        return $preHead->next;
    }

    // 迭代2
    // 执行用时：8 ms, 在所有 PHP 提交中击败了86.16%的用户
    // 存消耗：14.4 MB, 在所有 PHP 提交中击败了87.84%的用户
    function mergeTwoLists2($l1, $l2) {
        $preHead = new ListNode(0); // 声明一个头指针用来等待返回最后得到整个链表。
        $prev    = $preHead; // 这里是引用
        // 采用或的关系，这样只要有一个不为空就一直遍历
        // 如果采用全等判断，性能也会慢下来，所以这里采用!=
        while($l1 != null || $l2 != null) {
            // 如果l1可能为空，那么就直接把l2赋值过去
            if($l1 == null) {
                $prev->next = $l2;
                break;// 终止循环
            }
            // 同理，l2也可能为空
            if($l2 == null) {
                $prev->next = $l1;
                break;
            }
            // 如果两个都不是空的时候在进行比较
            if($l1->val <= $l2->val) {
                $prev->next = $l1; // prev向后移动,向较小的值移动
                $l1 = $l1->next;
            }else{
                $prev->next = $l2; 
                $l2 = $l2->next;
            }
            $prev = $prev->next; // 当前节点也需要向后移动
        }
        return $preHead->next;
    }

    // 递归
    // 执行用时：8 ms, 在所有 PHP 提交中击败了86.16%的用户
    // 内存消耗：14.4 MB, 在所有 PHP 提交中击败了81.08%的用户
    function mergeTwoLists($l1, $l2) {
        // 边界判断 有可能两个链表其中一个先遍历结束
        if($l1 === null) {
            return $l2;
        }else if($l2 == null){
            return $l1;
        }
        $list = null;
        // 判断两边谁大
        if($l1->val <= $l2->val) {
            $list = $l1; // 如果l1大，那么l1的链表就向后移动一个节点，
            $list->next = $this->mergeTwoLists($l1->next, $l2); // 然后继续递归比较
        }else{
            $list = $l2;
            $list->next = $this->mergeTwoLists($l1, $l2->next);
        }
        return $list;

    }
}
// @lc code=end


// @tacks think=start
/*
 

题意 meaning 

    合并两个有序链表，其实跟合并两个有序数组差不多，思路一样，只是数据结构有所不同，但是思路是一样的。

    而且都是升序链表，如果一个升序，一个降序，可能就要反着比较。这里都是升序，会相对简单一些。

关键 key 

    "有序的升序链表"
    "链表的遍历"

想法 idea

【1】链表遍历比较法
    时间复杂度是:O(N)   (N是两个链表最短长度)
    空间复杂度是:O(N+M) (声明新的链表空间)

    关于链表迭代的方式，老黑表示：&&改用|| 可能会更快一些，但是我实践过后，还是没有&&快。

    另外比较的时候，采用!=比较就行，不用全等!==比较，这样也会更快一些

【2】链表递归比较法
    时间复杂度是：O(N+M) （都需要不断递归到最后空节点）
    空间复杂度是：O(N+M)  (声明新的链表空间)
     



*/
// @tacks think=end