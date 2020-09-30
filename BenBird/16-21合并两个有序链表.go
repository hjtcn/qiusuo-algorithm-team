//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例：
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
// Related Topics 链表
// 👍 1290 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{0, nil}
	curr := l3
	num1, num2 := 0, 0
	temp := 0

	for l1 != nil || l2 != nil {
		if l1 == nil {
			curr.Next = l2
			break
		}

		if l2 == nil {
			curr.Next = l1
			break
		}

		num1, num2 = l1.Val, l2.Val

		if num1 < num2 {
			temp = num1
			l1 = l1.Next
		} else {
			temp = num2
			l2 = l2.Next
		}

		curr.Next = &ListNode{temp, nil}
		curr = curr.Next
	}

	return l3.Next
}
//leetcode submit region end(Prohibit modification and deletion)

/**
题解：合并两个有序链表

耗时：0ms
内存：2.6M

复杂度分析：
	时间复杂度：O(MAX(M, N))
	空间复杂度：O(N)
 */
