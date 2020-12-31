/*
 * @lc app=leetcode.cn id= lang=php
 *
 *  面试题 02.02. 返回倒数第 k 个节点
 *
 * https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/
 *
 *
 * 实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
 * 
 * 
 * 示例 ：
 *
 * 输入： 1->2->3->4->5 和 k = 2
 * 输出： 4
 * 
 * 
 * 
 * 提示：
 * 
 * 给定的 k 保证是有效的。
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
     * @param Integer $k
     * @return Integer
     */
     // 快慢指针法
     // 执行用时： 0 ms, 在所有 PHP 提交中击败了100.00%的用户
     // 内存消耗： 14.4 MB, 在所有 PHP 提交中击败了77.08%的用户
    function kthToLast($head, $k) {
        // 快慢指针
        $fast = $head; 
        $slow = $head; 
        // 快指针先跑k个节点
        for($i=0; $i<$k; $i++) {
            $fast = $fast->next;
        }
        // 然后快慢指针再一起跑
        while($fast != null) {
            $fast = $fast->next;
            $slow = $slow->next;
        }
        // 最后慢指针就落在倒数第k个节点
        return $slow->val;
    }
}
// @lc code=end


// @tacks think=start
/*
 

题意 meaning 

    题意就是想找倒数第k个节点的值，哈哈哈哈基本上简单的题目，题意都直给的很清楚。

关键 key 

    “链表的遍历”，“链表从前向后遍历next指针指向下一个节点”
  

想法 idea

数组：
    所有元素都连续的存储于一段内存中，且每个元素占用的内存大小相同。具备了通过下标快速访问数据的能力,时间复杂度O(1)。
    连续存储的缺点也很明显，增加容量，增删元素的成本很高，时间复杂度均为 O(n)
链表：
    由若干个结点组成，每个结点包含数据域和指针域。
    链表的存储方式使得它可以高效的在指定位置插入与删除，时间复杂度均为 O(1)。

【1】快慢指针

    非常好用！
    假设整个链表长度为 len
    1、定义两个指针，fast快，slow慢，一开始都是同一个位置。
    2、现在fast从头移动k个长度的节点位置，那么也就是后面还剩余 len-k
    3、然后slow再和fast一起走，直到fast走到结尾部分(len-k长度)，那么slow也就走len-k，因此slow最后的位置就是距离尾部k的位置。
    4、返回slow指向的节点即可。
综上所述
    时间复杂度O(n)
    空间复杂度O(1)

 



*/
// @tacks think=end