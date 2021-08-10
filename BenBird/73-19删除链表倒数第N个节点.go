//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 进阶：你能尝试使用一趟扫描实现吗？
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1], n = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
// Related Topics 链表 双指针
// 👍 1499 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
思路：
两次扫描法：
第一次遍历全链表，获取长度length
第二次遍历 (len - n) 个节点，到达目标节点的前一个节点
删除当前节点的后一个节点
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := head

	length := 0
	for temp != nil {
		temp = temp.Next
		length++
	}

	if length > n {
		return nil
	}

	if length == n {
		head = head.Next
		return head
	}

	temp_list := head
	for i := (length - n); i > 0; i-- {
		if i == 1 {
			p := temp_list.Next
			temp_list.Next = p.Next
			p.Next = nil
			break
		}
		temp = temp.Next
	}

	return head
}

func main()  {

}
//leetcode submit region end(Prohibit modification and deletion)
