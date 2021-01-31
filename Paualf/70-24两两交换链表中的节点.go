给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
示例 1：
输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：
输入：head = []
输出：[]
示例 3：
输入：head = [1]
输出：[1]
 
提示：
链表中节点的数目在范围 [0, 100] 内
0 <= Node.val <= 100
 
进阶：你能在不修改链表节点值的情况下解决这个问题吗?（也就是说，仅修改节点本身。）

1. Clearfication:

两两交换其中相邻的节点
偶数节点可以这样交换，如果是奇数节点呢？它会怎么交换节点呢
看着题目很简单，发现自己不知道怎么思考和解决了，不知道如何下手去解决问题了
没有掌握递归的精髓，你只需要保证当前的逻辑正确，n=1 => 正确，n=2 => 正确，那么 根据数学归纳法的话，那么 n= n-1 和 n =n 也是会正确的

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := head.Next
    head.Next = swapPairs(newHead.Next)
    newHead.Next = head
    return newHead
}
迭代法，通过加入一个头节点进行循环遍历:
func swapPairs(head *ListNode) *ListNode {
    dummyHead := &ListNode{0, head}
    temp := dummyHead
    
    for temp.Next != nil && temp.Next.Next != nil {
        node1 := temp.Next
        node2 := temp.Next.Next
        temp.Next = node2
        node1.Next = node2.Next
        node2.Next = node1
        temp = node1
    }
    
    return dummyHead.Next
}

复杂度分析：
时间复杂度：O(n) ： 遍历链表中的每个节点
空间复杂度：递归：O(n) ，迭代的空间复杂度是O(1)

总结：
1. 对递归还是不熟悉，分解子问题的能力有点弱

2. 有意识的锻炼自己分解子问题和解决问题的能力
