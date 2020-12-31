/*
 * @lc app=leetcode.cn id=2 lang=php
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (38.18%)
 * Likes:    4973
 * Dislikes: 0
 * Total Accepted:    562.9K
 * Total Submissions: 1.5M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 * 
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 * 
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 * 
 * 示例：
 * 
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
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
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2) {
        $carry = 0; // 某个位置的进位
        $sum   = 0; // 某个位置的和
        $out   = new ListNode(0); // 头结点
        $cur   = $out; // 对象赋值，是引用赋值
        while($l1 != null || $l2 != null) {
            $x = !empty($l1) ? $l1->val : 0; // 获取当前节点的值
            $y = !empty($l2) ? $l2->val : 0; // 获取当前节点的值
            $sum = $x + $y + $carry; // 当前节点的和
            if($sum >= 10) {
                $carry = 1; // 需要进位的,因为两个个位数的和一定小于20，所以进位的话只能是 1
                $cur = $cur->next = new ListNode($sum-10); // 当前的值 赋值给当前节点
            }else{
                $carry = 0; // 重置为0
                $cur = $cur->next = new ListNode($sum); // 当前的值 赋值给当前节点
            }
            // 继续向下遍历直到为空
            $l1 = $l1->next; 
            $l2 = $l2->next; 
        }
        // 可能最后还有进位
        if($carry > 0) {
            $cur->next = new ListNode(1);
        }
        return $out->next;
    }
}
// @lc code=end


// @tacks think=start
/*
 

题意 meaning 

    其实就是两个数字相加，像我们平时自己计算一样，个位相加，然后十位相加。。。判断是否需要进位

关键 key 

    "carry 进位时机"
    "链表的遍历"

想法 idea

【1】链表遍历法
    时间复杂度是:O(N)
    空间复杂度是:O(N)
     



*/
// @tacks think=end