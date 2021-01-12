<?php
/*
 * @Descripttion: 
 * @Author: tacks321@qq.com
 * @Date: 2021-01-06 18:53:32
 * @LastEditTime: 2021-01-07 17:32:45
 */

/*
 * @lc app=leetcode.cn id=86 lang=php
 *
 * [86] 分隔链表
 *
 * https://leetcode-cn.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (60.01%)
 * Likes:    350
 * Dislikes: 0
 * Total Accepted:    82.9K
 * Total Submissions: 133.4K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * 给你一个链表和一个特定值 x ，请你对链表进行分隔，使得所有小于 x 的节点都出现在大于或等于 x 的节点之前。
 * 
 * 你应当保留两个分区中每个节点的初始相对位置。
 * 
 * 
 * 
 * 示例：
 * 
 * 
 * 输入：head = 1->4->3->2->5->2, x = 3
 * 输出：1->2->2->4->3->5
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
// 节点 数据结构
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
     * @param ListNode $head
     * @param Integer $x
     * @return ListNode
     */
    // 执行用时：4 ms, 在所有 PHP 提交中击败了95.18%的用户
    // 内存消耗：15.1 MB, 在所有 PHP 提交中击败了91.57%的用户
    function partition($head, $x) {
        // 哑节点定义
        // 即它们的 next 指针指向链表的头节点
        $smallDummy = new ListNode(0);
        $larageDummy = new ListNode(0);

        $smallHead = $smallDummy;
        $largeHead = $larageDummy;
        while($head) {
            $cur = $head; // 用cur来标识当前节点 会思路比较清晰一点
            
            // 与基准值比较进行分割
            if($cur->val < $x) {
                // 小于特定值的部分
                $smallHead->next = $cur;
                $smallHead = $cur;
            } else {
                // 大于等于特定值的部分
                $largeHead->next = $cur;
                $largeHead = $cur;
            }

            // 向下遍历
            $head = $head->next; 
        }
        
        // 合并
        $largeHead->next = null; // 细节点
        $smallHead->next = $larageDummy->next;
        
        return $smallDummy->next;
    }
}
// @lc code=end

/*

其实：题意就是说，给你一个基准值，然后将链表一分为二然后再连起来

链表的题就是要画图!!!!


遍历链表所有节点
维护两个链表 
    small 小于特定值的部分
    large 大于等于特定值的部分
链表哑节点使用
    哑节点（dummy node）是初始值为NULL的节点，创建在使用到链表的函数中，可以起到避免处理头节点为空的边界问题的作用，减少代码执行异常的可能性。

    


*/