//给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//
//
// 进阶：
//
// 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
//
//
//
// 示例：
//
// 输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 8 -> 0 -> 7
//
// Related Topics 链表
// 👍 291 👎 0
package main

type ListNode struct {
	Val int
	Next *ListNode
}

type Stack struct {
	nums []int
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
	list_one, list_two := l1, l2
	l3 := &ListNode{0, nil}
	s_one := initStack()
	s_two := initStack()

	//将链表l1入栈
	for list_one != nil {
		s_one.push(list_one.Val)
		list_one = list_one.Next
	}

	//将链表l2入栈
	for list_two != nil {
		s_two.push(list_two.Val)
		list_two = list_two.Next
	}

	carry := 0
	curr_num := 0
	curr_stack := initStack()
	for !s_one.isEmpty() || !s_two.isEmpty() || carry != 0 {
		num_one := s_one.pop()
		num_two := s_two.pop()

		sum := num_one + num_two + carry

		//控制进位数
		if sum > 9 {
			carry = 1
		} else {
			carry = 0
		}
		curr_num = sum % 10

		//将获取求和、进位处理后的数字入栈
		curr_stack.push(curr_num)
	}

	//将求和结果出栈存入链表l3
	curr_list := l3
	for !curr_stack.isEmpty() {
		curr_list.Next = &ListNode{curr_stack.pop(), nil}
		curr_list = curr_list.Next
	}

	return l3.Next
}

/**
初始化空栈
 */
func initStack() *Stack {
	return &Stack{
		nums: []int{},
	}
}

/**
入栈
 */
func (s *Stack) push(n int) {
	s.nums = append(s.nums, n)
}

/**
出栈
 */
func (s *Stack) pop() int {
	if s.isEmpty() {
		return 0
	}

	res := s.nums[len(s.nums) - 1]
	s.nums = s.nums[:len(s.nums) - 1]

	return res
}

/**
判断栈是否为空
 */
func (s *Stack) isEmpty() bool {
	if len(s.nums) == 0 {
		return true
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}

/**
题解：两数相加II

耗时：20ms
内存：6.2M

因本题的链表是从高位到低位一次排列，所以不能直接进行遍历链表依次进行求和，
所以将两链表从第一个节点开始分别存储到两个栈s_one, s_two中，
然后再同时出栈，将出栈两元素进行求和，同时控制好进位数，
但是这样求和结果是逆序的，所以先求和结果存入栈curr_stack
然后将curr_stack中的数一次出栈，存入链表中
这样一正一反，求和结果就变成了正序

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(n)
 */
