<?php
/*
 * @Descripttion: 周二的题，追一追
 * @Author: tacks321@qq.com
 * @Date: 2021-08-10 18:10:08
 * @LastEditTime: 2021-08-11 10:50:22
 */


/*
 * @lc app=leetcode.cn id=19 lang=php
 *
 * [19] 删除链表的倒数第 N 个结点
 *
 * https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (42.01%)
 * Likes:    1497
 * Dislikes: 0
 * Total Accepted:    463.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 * 
 * 进阶：你能尝试使用一趟扫描实现吗？
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：head = [1,2,3,4,5], n = 2
 * 输出：[1,2,3,5]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：head = [1], n = 1
 * 输出：[]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：head = [1,2], n = 1
 * 输出：[1]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中结点的数目为 sz
 * 1 
 * 0 
 * 1 
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

/**
 * 先后双指针
 * 
 * 执行用时：0 ms, 在所有 PHP 提交中击败了100.00%的用户
 * 内存消耗：15.2 MB, 在所有 PHP 提交中击败了36.55%的用户
 * 
 * @date  2021-08-11 10:13:23
 */
class Solution {

    /**
     * @param ListNode $head
     * @param Integer $n
     * @return ListNode
     */
    // 先后双指针
    function removeNthFromEnd($head, $n) {
        // 声明哑节点
        $dummy = new ListNode(0);
        $dummy->next = $head;

        // 定义 先后指针
        $front = $head;
        $after = $dummy; // 这样是为了直接找到删除节点的前驱节点

        // after - front 

        // front 先指针，首先前进n
        while($n--) {
            $front = $front->next;
        }

        // after - 0 - 1 - 2 ... - n(front)  这个时候相隔 n个节点 

        // front after 先后指针 再一起走
        while($front != null) {
            $front = $front->next;
            $after = $after->next; 
        }
        // after是 要删除节点的前驱节点
        $after->next = $after->next->next;
        return $dummy->next;
    }
}


/**
 * 两次遍历
 * 
 *  执行用时：8 ms, 在所有 PHP 提交中击败了54.82%的用户
 *  内存消耗：15.2 MB, 在所有 PHP 提交中击败了57.36%的用户
 * @date  2021-08-11 10:13:23
 */
class Solution2 {

    /**
     * @param ListNode $head
     * @param Integer $n
     * @return ListNode
     */
    function removeNthFromEnd($head, $n) {
        $size = 0;

        // 声明哑节点
        $dummy = new ListNode(0);
        $dummy->next = $head;

        // 计算链表长度
        while($head != null) {
            $head = $head->next;
            $size++;
        }

        $head = $dummy; // 链表重置

        for ($i=0; $i < $size-$n; $i++) { 
            $head = $head->next;
        }

        // 删除节点
        $head->next = $head->next->next;

        return $dummy->next;
    }
}

// @lc code=end


/*

先后双指针+两次遍历
【 Clarification 】

题目：
    给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
    进阶：你能尝试使用一趟扫描实现吗？
示例：
    输入：head = [1,2,3,4,5], n = 2
    输出：[1,2,3,5]
    

哑节点的用法：======》 强调哑节点的用法
    在对链表进行操作时，添加一个哑节点（dummy node），它的 next 指针指向链表的头节点。 这样就不需要对头节点进行特殊的判断了。
    
【 Possible-Solution 】
1、先后指针
        通过定义 先后两个指针，从而实现一次扫描进行处理，从而找到要删除节点的前驱节点，这样可以更方便的删除节点
    时间复杂度 O(N)
    空间复杂度 O(1)
2、两次遍历
        试试两次遍历看看效果如何。先通过第一次遍历获取链表长度，然后再计算出来到倒数第n个节点，是正数第几个。
    时间复杂度 O(N)
    空间复杂度 O(1)
    
*/