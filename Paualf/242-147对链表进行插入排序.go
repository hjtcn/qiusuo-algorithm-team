对链表进行插入排序。
插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。

插入排序算法：
插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。
 
示例 1：
输入: 4->2->1->3
输出: 1->2->3->4

示例 2：
输入: -1->5->3->4->0
输出: -1->0->3->4->5
通过次数103,451提交次数151,590

1. Clarification:
插入排序：
每次从链表中找到最小的那个节点
将节点连接起来

如果是在数组中怎么排序呢？

如果是在数组中的话，每次找到最小节点的位置，然后和当前位置交换，如果在链表中是不是也可以的哈

也可以的话：
我们要解决这样几个问题：
1.1: 找到最小的节点的时候断开节点
1.2: 将该节点插入到对应的位置
1.3: 对应的位置是哪里？需要有一个指针进行标记

如何断开节点，如果我找到最小节点是1，我怎么断开，我要把1的前一个节点的.Next = pre.Next.Next
我们可以构建一个map去存储，节点的值 => 节点对应的前一个节点

Q: 如何将在链表中找到最小的节点这个操纵循环下去呢？
A：这个地方没有想清楚哈。。。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil {
        return
    }

    // 构建map
    curr := head
    m := make(map[int]*ListNode,0)
    var pre *ListNode
    for curr != nil {
        m[curr.Val] = pre
        pre = curr
        curr = curr.Next
    }

    // findMinNode
    for 
}

2. 看题解：
看完题解以后发现是自己思考的方向错了， 插入排序，每次找到对应的位置进行插入就行了，自己想的好像是冒泡排序了

光看题解是看不懂的，要动手画一下图

题解里面的 具体过程还是蛮详细的值得学习一下

func insertionSortList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    dummyHead := &ListNode{Next: head}
    lastSorted, curr := head, head.Next
    for curr != nil {
        if lastSorted.Val <= curr.Val {
            lastSorted = lastSorted.Next
        } else {
            prev := dummyHead
            for prev.Next.Val <= curr.Val {
                prev = prev.Next
            }
            lastSorted.Next = curr.Next
            curr.Next = prev.Next
            prev.Next = curr
        }
        curr = lastSorted.Next
    }
    return dummyHead.Next
}


3. 复杂度分析：
时间复杂度：O(nxn)
空间复杂度：O(1)

4. 总结：
4.1: 我们都是通过生活和工作，不断的完善自己的认知，以及对事物和事情的理解

4.2: 程序是一种思维和思考方式的表达和展示
