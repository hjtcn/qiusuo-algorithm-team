//输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。例如，一个链表有6个节点，从头节点开始，
//它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。
//
//
//
// 示例：
//
// 给定一个链表: 1->2->3->4->5, 和 k = 2.
//
//返回链表 4->5.
// Related Topics 链表 双指针
// 👍 85 👎 0
package main

import (
	"fmt"
	"math/rand"
	"os"
)

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
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, low := head, head

	for k > 0 {
		fast = fast.Next
		k--
	}

	for fast != nil {
		fast, low = fast.Next, low.Next
	}

	return low
}
//leetcode submit region end(Prohibit modification and deletion)
//测试代码  start

/**
	创建一个非空链表
 */
func initList() *ListNode {
	head := &ListNode{0, nil}
	tail := head

	for i := 1; i <= 10; i++ {
		tail.Next = &ListNode{Val:rand.Intn(299)}
		tail = tail.Next
		fmt.Println(tail.Val, "初始化")
	}

	return head
}

func main() {

	head := initList()
	k := 1
	res := getKthFromEnd(head, k)

	fmt.Println(*res)
	os.Exit(33)

	/**
	temp := head.Next

	for temp != nil {
		fmt.Println(temp.Val, "遍历")
		temp = temp.Next
	}
	*/
}

//测试代码 end

/**
题解：链表中倒数第k个节点

耗时：0ms
内存：2.2M

双指针(快慢指针):
	使用双指针：fast, low, 首先将fast、low指针指向head指针，然后现将fast指针向前移动k个位置后，
	接着讲fast, low每次向后移动 1 个位置，知道fast指针，移动到链表之外，此时low位置上为倒数第k个
	然后将low指针返回即可

时间复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(1)

总结：
	Go对链表的实现，用到了结构体，今天在利用结构体构建一个链表时，在对Next指针赋值时报错，may lead to nil dereference(可能是空指针取消引用)
	这是在定义一个结构体时，未对其初始化，导致的报错

 */
