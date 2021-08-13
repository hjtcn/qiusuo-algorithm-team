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

/**
思路：
两次扫描法：
第一次遍历全链表，获取长度length
第二次遍历 (len - n) 个节点，到达目标节点的前一个节点
删除当前节点的后一个节点
*/
func removeNthFromEndTwoError(head *ListNode, n int) *ListNode {
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
/**
两次扫描优化
把获取链表长度拆出一个方法就少会少定义变量
 */
func getLinkListLength(head *ListNode) int {
	length := 0
	for ; head != nil ; head = head.Next {
		length++
	}

	return length
}

func removeNthFromEndTwo(head *ListNode, n int) *ListNode {
	length := getLinkListLength(head)
	dummy := &ListNode{0, head}
	temp := dummy

	for i := 1; i <= (length - n); i++ {
		temp = temp.Next
	}

	temp.Next = temp.Next.Next

	return dummy.Next
}

/**
双指针：
定义fast low指针，
fast 先向前移动n, low再和fast一起想前移动
直到fast移动到最后一个节点，停止移动
此时移除low后面的节点
 */
func removeNthFromEndError(head *ListNode, n int) *ListNode {
	fast := head
	low := head

	for i := 1; i <= n; i++ {
		if fast == nil {
			return nil
		}

		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		if fast != nil {
			low = low.Next
		}
	}

	if low == head {
		head = head.Next
	} else {
		p := low.Next
		low.Next = p.Next
		p.Next = nil
	}

	return head
}

/**
上面的代码是有问题的，在删除第二个的时候 和删除第一个的结果是一样的
主要是没有引入头结点，导致的删除第一个和第二个 不好处理导致的
没有引入头结点的原因是，LeetCode验证时初始化是没有头结点的，导致思想被局限
看了题解后，才知道在移除方法中强行给原链表加头结点，方便处理
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, low := head, dummy

	for i := 1; i <= n; i++ {
		fast = fast.Next
	}

	for ; fast != nil; fast = fast.Next {
		low = low.Next
	}

	low.Next = low.Next.Next

	return dummy.Next
}
/**
还是使用了头结点方便了多，
处理细节：
	low初始位置是头结点，fast初始位置和原链表第一节点位置
	fast先向前移动n个位置后，low和fast再一起移动，直到fast的位置是nil
	此时删除low的后一个节点即可

注意到的问题：
删除某一个节点之前的做法是
	p = temp.Next
	temp.Next = p.Next
	p.Next = nil
看题解是
	temp.Next = temp.Next.Next
区别，将删除的节点赋值给新的变量p，然后将该节点p与它的下个节点解除关系，也就是p.Next = nil
现在是直接忽略将要删除节点的指向，直接将它与它的前一个节点解除关系 temp.Next = temp.Next.Next
不知道是不是因为Go的GC直接会把未引用的节点给回收了，有时间好好研究一下Go的GC
 */

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
	arr := []int{1}
	obj := initLinkList(arr)
	//obj.printLinkList()

	//hah := removeNthFromEnd(obj, 3)
	hah := removeNthFromEnd(obj, 1)
	hah.printLinkList()
}

/**
总结：
复杂度分析：
	两遍扫描法：
		时间复杂度：O(L) L是链表的长度
		空间复杂度：O(1) 未使用多余的空间
	双指针法：
		时间复杂度：O(L) L是链表长度
		空间复杂度：O((1)

从上道题看到鹏飞使用了Debug去排查问题，这次也尝试使用一下排查问题，嗯，还真不错，程序运行的轨迹都很清晰的展示给你
这次遇到的好多问题，都是使用debug排查到的，但是用了很长时间，debug熟悉当中。。。
 */
