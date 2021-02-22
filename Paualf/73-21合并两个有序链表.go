/*
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
*/

1. Clearfication:
初始化一个链表，然后比较 l1和l2 的节点元素大小，将元素小的连接到初始化链表的后面
伪代码：
newList := &ListNode{Val:0,Next:nil}
head := newList
for (l1 && l2) {
    if l1.Val <= l2.Val {
        head.Next = l1
        head = head.Next
        l1 = l1.Next
    }else {
        head.Next = l2
        head = head.Next
        l2 = l2.Next
    }
}
if (l1) {
    head.Next = l1
}
if (l2) {
    head.Next = l2
}
return newList.Next

2. Coding
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
       newList := &ListNode{
            Val:0,
            Next:nil,
        }
    
        head := newList
        for l1 != nil && l2 != nil {
            if l1.Val <= l2.Val {
                head.Next = l1
                head = head.Next
                l1 = l1.Next
            }else {
                head.Next = l2
                head = head.Next
                l2 = l2.Next
            }
        }
        if l1 != nil {
            head.Next = l1
        }
        if l2 != nil {
            head.Next = l2
        }
        return newList.Next
}

3. 看题解
看了题解以后看到了递归的写法：
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }else if l2 == nil {
        return l1
    }else if l1.Val < l2.Val {
        l1.Next = mergeTwoLists(l1.Next,l2)
        return l1
    }else {
        l2.Next = mergeTwoLists(l1,l2.Next)
        return l2
    }
}

4. 复杂度分析
时间复杂度：O(n+m) 遍历两个链表进行比较
空间复杂度：迭代法：O(1) 的空间复杂度，递归：O(n+m) 的栈空间

5. 总结
5.1: 合并元素的时候想到了归并排序，对归并排序的细节点还是不是很熟悉，拆分子数组，然后去合并
