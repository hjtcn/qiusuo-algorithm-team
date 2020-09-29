//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
// Related Topics 链表 数学
// 👍 4843 👎 0
package main

type ListNode struct {
	Val int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{0, nil}
	curr := l3
	carry := 0
	sum := 0

	for l1 != nil || l2 != nil {
		num1, num2 := 0, 0

		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}

		sum = num1 + num2 + carry

		if sum >= 10 {
			carry = 1
			sum = sum % 10
		} else {
			carry = 0
		}

		curr.Next = &ListNode{sum, nil}
		curr = curr.Next
	}

	//当链表末尾需要进位时，继续创建一个新节点
	if carry != 0 {
		curr.Next = &ListNode{carry, nil}
	}

	return l3.Next
}
//leetcode submit region end(Prohibit modification and deletion)

func main()  {

}

/**
题解：两数相加

耗时：12ms
内存：4.9M

注意：
	进位问题
		链表末尾进位
		非链表末尾进位
	两链表长度不一问题

进位问题可设置一个进位变量 carry，每次判断carry是否为 0，来辨别是否有进位

首先遍历两个链表，只要存在一个链表未到末尾，就继续遍历
	遍历时将链表的两个数值 num1、num2、carry进行相加，得sum
	判断sum 是否大于 9
		大于，则需要求进位数 carry = sum / 10; sum = sum % 10
		小于，进位数 carry = 0; sum = sum

	然后创建一个新节点，val = sum, next = nil

直到最后一个链表完成遍历，即可得到一个新的链表，返回

复杂度分析：
	时间复杂度：O(max(m, n))
	空间复杂度：O(n)
 */
