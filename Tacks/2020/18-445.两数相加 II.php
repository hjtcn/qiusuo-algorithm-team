<?php

/*
 * @lc app=leetcode.cn id=445 lang=php
 *
 * [445] 两数相加 II
 */

/*
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶：如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

示例：
    输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
    输出：7 -> 8 -> 0 -> 7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。



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
    // 栈方法
    // 执行用时：20 ms, 在所有 PHP 提交中击败了90.00%的用户
    // 内存消耗：14.6 MB, 在所有 PHP 提交中击败了16.67%的用户
    function addTwoNumbers($l1, $l2) {
        $num1 = []; // 链表1的栈
        $num2 = []; // 链表2的栈
        while($l1 != null) {
            $num1[] = $l1->val; //将当前的值进入数组，最高位的先进栈
            $l1 = $l1->next;
        }
        while($l2 != null) {
            $num2[] = $l2->val; //将当前的值进入数组
            $l2 = $l2->next;
        }
        $carry = 0; // 某个位置的进位
        $sum   = 0; // 某个位置的和

        $head = null; // 最后返回的链表
        // 直到两个栈都不为空，或者最后还有进位
        while(!empty($num1) || !empty($num2) || $carry) {
            // 当前位置两数之和 (上一次是否有进位)
            // array_pop 弹出栈，首先弹出来的是最低位的
            $sum = (array_pop($num1)??0) + (array_pop($num2)??0) + $carry;
            if($sum >= 10) {
                $carry = 1;
                $sum = $sum - 10;
            }else{
                $carry = 0;
                $sum = $sum;
            }
            // 采用头插法 插入新的节点
            $cur = new ListNode($sum);
            $cur->next = $head; // 先将当前节点插入 链表头部
            $head = $cur;       // 然后再用链表头部指向
        }
        return $head;
    }
}
// @lc code=end



// @tacks think=start
/*
 

题意 meaning 

    这次的两个链表相加求和的链表，当然与上次不同的是，上次两数相加本来就是低位在链表的头节点位置，这次是高位在链表头节点

    如果我们直接反转链表，那么其实就跟【2.两数相加】的题目一样，所以这里我们只能另外考虑。

    我们可以使用栈来存储逆序的链表，这样就可以倒着计算，先计算个位，然后十位，等等


关键 key 

    "链表的遍历"
    "栈的使用：先进后出"

想法 idea

【1】栈
    
    我们可以将所有数字压入栈中，然后再依次取出来进行相加求和，同时考虑是否进位。
    PHP这里采用数组来实现栈的功能。

    (array_pop($num1)??0) 这个代码就是为了弹出num1的元素，因为可能为空也就是null,所以采用三元运算符判断一下

    PHP7中可以这样用  (expr1) ?? (expr2) 
        等同于
    isset(expr1) ? expr1 : expr2 ;   => isset(array_pop($num1)) ? array_pop($num1) : 0
   
时间复杂度：O(max(N,M)) （N，M为两个链表长度，需要遍历到都是空）
空间复杂度：O(N+M) (N，M存储两个数组栈)


*/
// @tacks think=end