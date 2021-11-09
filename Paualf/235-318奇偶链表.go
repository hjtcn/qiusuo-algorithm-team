给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:
输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL

示例 2:
输入: 2->1->3->5->6->4->7->NULL 
输出: 2->3->6->7->1->5->4->NULL

说明:
应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。

1. Clarification:

原地的方法没有想出来，那么就使用非原地的方法

用两个链表，oddPre 和 evenPre 
1.1: 最后遍历将后面的加到前面链表的后面

根据自己思路写的，panic了
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    curr := head
    oddPre := &ListNode{}
    evenPre := &ListNode{}

    var oddPreHead,evenPreHead *ListNode
    oddPreHead = oddPre
    evenPreHead = evenPre
    
    for curr != nil {
        oddPreHead.Next = curr
        evenPreHead.Next = curr.Next

        if curr.Next == nil {
            break
        }

        oddPreHead = oddPreHead.Next
        evenPreHead = evenPreHead.Next
        curr = curr.Next.Next
    }

    curr = oddPre
    for curr.Next != nil {
        curr = curr.Next
    }

    curr.Next = evenPre.Next

    return oddPre.Next
}

2. 看题解：
分离节点后合并

将奇数节点和偶数节点分离成奇数链表和偶数链表，然后将偶数链表连接在奇数链表之后，合并后的链表即为结果链表
偶数链表的头节点：evenHead = head.Next,evenHead 是偶数链表的头节点

odd和even分别指向奇数节点和偶数节点，初始化 odd = head,even = evenHead
每一步先更新奇数节点，然后更新偶数节点
更新奇数节点时：奇数节点的后一个节点需要指向偶数节点的后一个节点，odd.Next = even.Next,然后令odd = odd.Next,此时odd变成even 的后一个节点
更新偶数节点的时候：偶数节点的后一个节点需要指向奇数节点的后一个节点，even.Next = odd.Next,even = even.Nxt
分离完毕的条件是：even 为空或者even.Next 为空的时候

func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
    	return head
    }
    
    evenHead := head.Next
    odd := head
    even := evenHead
    
    for even != nil && even.Next != nil {
     	odd.Next = even.Next
        odd = odd.Next
        even.Next = odd.Next
        even = even.Next
    }
    odd.Next = evenHead
    return head
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1: 看完题解感觉真的厉害，思路也是棒棒棒的，学习，享受编程和写代码带来的快乐
