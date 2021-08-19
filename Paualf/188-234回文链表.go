给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。


示例 1：


输入：head = [1,2,2,1]
输出：true
示例 2：


输入：head = [1,2]
输出：false
 

提示：

链表中节点数目在范围[1, 105] 内
0 <= Node.val <= 9
 

进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

Clarification:
本来以为这道题挺简单的，但是自己没有想出来思路

想到快慢指针，快指针每次走2步，慢指针每次走1步，快指针走到末尾的时候，慢指针走到中间，然后再使用一个指针从头开始遍历，慢指针从中间开始走，没有想通奇数和偶数结点的时候的区别

看题解:
2.1: 把值放到数组中去比较：
func isPalindrome(head *ListNode) bool {
    vals := []int{}
    for ; head != nil; head = head.Next {
        vals = append(vals, head.Val)
    }
    n := len(vals)
    for i, v := range vals[:n/2] {
        if v != vals[n-1-i] {
            return false
        }
    }
    return true
}
2.2: 递归，虽然可能不是很喜欢和容易想到，但是也是一种方法

2.3: 快慢指针反转后半段链表


复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

总结：
4.1: 以为这道题挺简单的，看下来是自己挺简单的。。。

4.2: 感觉对单链表不是很熟悉哈。。。
