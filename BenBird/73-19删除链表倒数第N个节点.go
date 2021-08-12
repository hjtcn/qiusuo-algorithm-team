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

import (
	"fmt"
	"strconv"
)

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
type ListNode struct {
	Val int
	Next *ListNode
}

func initLinkList(arr []int) *ListNode {
	link_list := &ListNode{0, nil}
	temp := link_list

	for _, val := range arr {
		temp.Next = &ListNode{val, nil}
		temp = temp.Next
	}

	return link_list.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := head

	length := 0
	for temp != nil {
		temp = temp.Next
		length++
	}

	if n > length || n < 1 {
		return nil
	}

	temp_list := head
	if length == n {
		return head.Next
	} else {
		for i := (length - n); i > 0; i-- {
			if i == 1 {
				p := temp_list.Next
				temp_list.Next = p.Next
				p.Next = nil
				break
			}

			temp_list = temp_list.Next
		}
	}

	return head
}

func (this *ListNode) printLinkList() {
	str := ""
	for this != nil {
		if this.Next != nil {
			str = str + strconv.Itoa(this.Val) + "->"
		} else {
			str = str + strconv.Itoa(this.Val)
		}
		
		this = this.Next
	}

	fmt.Println(str)
}

func main()  {
	arr := []int{1, 2, 3}
	obj := initLinkList(arr)
	//obj.printLinkList()

	hah := removeNthFromEnd(obj, 2)
	hah.printLinkList()
}
//leetcode submit region end(Prohibit modification and deletion)
