将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：
输入：l1 = [], l2 = []
输出：[]

示例 3：
输入：l1 = [], l2 = [0]
输出：[0]
 
提示：
两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列

1. Clarification:

新建一个链表l，l1,l2 元素相比较，如果l1 <= l2, 将l1元素放到l后面，否则将l2元素放到后面

l1和l2有一个链表为空，则退出循环

最后检查l1和l2是否为空

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    l := &ListNode{}
    curr := l

    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            curr.Next = list1
            list1 = list1.Next
        }else {
            curr.Next = list2
            list2 = list2.Next
        }

        curr = curr.Next
    }

    if list1 != nil {
        curr.Next = list1
    }

    if list2 != nil {
        curr.Next = list2
    }

    return l.Next
}
2. 看题解：

递归：
// terminator
if l1 == nil {
	return l2
}else if l2 == nil {
	return l1
}else if l1.Val < l2.Val {
    l1.Next = mergeTwoLists(l1.Next,l2)
    return l1
}else {
    l2.Next = mergeTwoLists(l1,l2.Next)
    return l1
}

3. 复杂度分析：
时间复杂度：O(n+m)
空间复杂度：O(n+m)

4. 总结：
4.1: just do it 
