<?php
/*
 * @lc app=leetcode.cn id=160 lang=php
 *
 * [160] 相交链表
 * 
 * 
 * 
注意：
    如果两个链表没有交点，返回 null.
    在返回结果后，两个链表仍须保持原有的结构。
    可假定整个链表结构中没有循环。
    程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists
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
     * @param ListNode $headA
     * @param ListNode $headB
     * @return ListNode
     */
    // 暴力法   
    // 超时！！！
    function getIntersectionNode1($headA, $headB) {
        // 为空判断
        if($headA === null || $headB === null) {
            return null;
        }
        $ha = $headA;
        $hb = $headB;
        while($ha != null) {
            while($hb != null){
                // 当前结点地址完全一样
                if($ha === $hb) {
                    return $ha;
                }
                $hb = $hb->next;
            }
            $ha = $ha->next;
            $hb = $headB; // 再次遍历B
        }
        return null;
    }

    // 哈希表法
    // 执行用时：1760 ms, 在所有 PHP 提交中击败了6.31%的用户
    // 内存消耗：23.1 MB, 在所有 PHP 提交中击败了69.51%的用户
    function getIntersectionNode2($headA, $headB) {
        // 为空判断
        if($headA === null || $headB === null) {
            return null;
        }
        $ha = $headA;
        $hb = $headB;
        $map = []; // 哈希表
        // 先遍历链表A，然后存成map
        while($ha != null) {
            array_push($map, $ha);// 追加到map中
            $ha = $ha->next;
        }
        // 遍历链表B，看看有没有节点在map里面
        while($hb != null) {
            // true严格匹配 map中有无包含hb
            if(in_array($hb, $map, true)) {
                return $hb;
            }
            $hb = $hb->next;
        }
        return null; // 最后都没有找到
    }

    // 双指针法
    // 执行用时：76 ms, 在所有 PHP 提交中击败了55.86%的用户
    // 内存消耗：23 MB, 在所有 PHP 提交中击败了84.15%的用户
    function getIntersectionNode($headA, $headB) {
        // 为空判断
        if($headA === null || $headB === null) {
            return null;
        }
        $ha = $headA;
        $hb = $headB;
        // 采用全等比较，比较地址
        while($ha !== $hb) {
            $ha = ($ha === null) ? $headB : $ha->next;
            $hb = ($hb === null) ? $headA : $hb->next;
        }
        // 相交节点
        return $ha;
    }
}
// @lc code=end



// @tacks think=start
/*
 

题意 meaning 

   如果两个链表相交，第一个相交的节点就是起始节点，后面相交的长度是一样的，也就是值都相等。

关键 key 

    "链表的遍历"
    "都是线性的，没有循环的部分"
    "相交的两个链表，结尾结点一定相同"
    "哈希表法"
    "双指针"

想法 idea

【1】暴力法
    这个就是遍历A链表，里面遍历B链表，然后看每一个A链表结点，是否存在B链表结点是一样的。
    这个方法太暴力，时间复杂度也高，直接放弃。！

时间复杂度：O(MN) （M,N为两个链表长度）
空间复杂度：O(1) (只需要遍历两个链表，没有用到额外的变量)

【2】哈希表法
    这样，我们遍历链表A，然后把每一个节点的值都存储在哈希表中
    然后遍历链表B，对每一个节点进行判断是否在哈希表中
    存在则直接返回当前节点，遍历结束还没有则返回null
时间复杂度：O(M+N) （两次遍历链表）
空间复杂度：O(M)或者O(N) (采用哈希表存储节点)

【3】双指针法(最佳！！！)
    创建两个指针pA pB ，然后都指向两个链表的头节点，然后逐步向后遍历。
    当pA链表到达A链表的尾部的时候，重定向到B的头节点，
    当pB链表到达B链表的尾部的时候，重定向到A的头节点，
    如果在某一时刻相遇 pA == pB ，那么当前就是相交节点。
原因：
    例如链表A = {1，6，9，7} \ 链表B = {5,7}
    根据题意，如果两个链表相交，那么他们的尾节点一定是一样的
    这么看来，其实相交的节点，到结尾节点的距离是一样的。
    我们的目的：通过重定向消除两个链表的长度差。
时间复杂度：O(M+N) （两次遍历链表）
空间复杂度：O(1) （存储两个指针）

*/
// @tacks think=end