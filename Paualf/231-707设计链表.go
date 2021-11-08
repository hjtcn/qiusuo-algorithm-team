设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
 
示例：
MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3
 
提示：

所有val值都在 [1, 1000] 之内。
操作次数将在  [1, 1000] 之内。
请不要使用内置的 LinkedList 库。

1. Clarification:
Q: 链表的结构是咋样的呢？
A: 里面包含 一个值 val int 和 下一个链表结构的指向

type MyLinkedList struct {
      Val int
Next *MyLinkList
}

Q: 如何初始化呢？
A: 没有想很清楚,返回一个带有头节点的链表

Q: 如何获取链表中第index 个结点的值？如果索引无效，返回-1
A: simple 实现，从头开始遍历，遍历到 index 个结点的时候返回，更好的实现，有一个 hash 的结果记录这个位置的地方

Q: 如何加入到链表头节点中？
A: newNode := &MyLinkedList, newNode.Next = head.Next,head.Next = newNode

Q: 如何加入到链表的末尾？
A: 没有想清楚的地方是如何快速定位到链表的最后一个元素呢？

Q: 如何删除某个位置的元素信息？
A: 找到它的前置结点，然后 node.Next = node.Next.Next


2. 看题解：
哨兵结点: 用来简化插入和删除的操作

因为伪头的关系 addAtHead 和 addAtTail 可以使用addAtIndex 来完成

使用哨兵节点，统一头节点的行为
使用prev 迭代
type ListNode struct {
    Val int
    Next *ListNode
}

type MyLinkedList struct {
    size int
    head *ListNode
}


func Constructor() MyLinkedList {
    return MyLinkedList{
        0,
        &ListNode{},
    }
}


func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }
    prev := this.head
    for i := 0;i < index;i++{
        prev = prev.Next
    }

    return prev.Next.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
    this.AddAtIndex(0,val)
}


func (this *MyLinkedList) AddAtTail(val int)  {
    this.AddAtIndex(this.size, val)
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index < 0 || index > this.size {
        return
    }

    prev := this.head
    for i := 0;i < index;i++{
        prev = prev.Next
    }

    node := &ListNode{Val:val}
    node.Next = prev.Next
    prev.Next = node
    this.size++
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= this.size {
        return
    }
    prev := this.head
    for i := 0;i < index;i++ {
        prev = prev.Next
    }
    prev.Next = prev.Next.Next
    this.size--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
双链表：
使用两个哑节点来作为链表的头尾节点，这两个哑节点可以方便链表中节点的增删

type ListNode struct {
    Val int
    Next,Pre *ListNode
}

type MyLinkedList struct {
    head,tail *ListNode
    size int
}


func Constructor() MyLinkedList {
    head := &ListNode{Val:0,}
    tail := &ListNode{Val:0,}
    head.Next = tail
    tail.Pre = head
    return MyLinkedList{
        head:head,
        tail:tail,
        size:0,
    }
}


func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }
    node := this.head.Next
    for i := 0;i < index;i++ {
        node = node.Next
    }
    return node.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
    newHead := &ListNode{
        Val:val,
        Pre:this.head,
        Next:this.head.Next,
    }
    this.head.Next.Pre = newHead
    this.head.Next = newHead
    this.size++
}


func (this *MyLinkedList) AddAtTail(val int)  {
    newTail := &ListNode{
        Val:val,
        Pre:this.tail.Pre,
        Next:this.tail,
    }
    this.tail.Pre.Next = newTail
    this.tail.Pre = newTail
    this.size++
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index > this.size {
        return
    }
    if index < 0 {
        this.AddAtHead(val)
        return
    }
    node := this.head
    // 注意这里是 <= 
    for i := 0;i <= index;i++ {
        node = node.Next
    }
    tmp := &ListNode {
        Val:val,
        Pre:node.Pre,
        Next:node,
    }
    node.Pre.Next = tmp
    node.Pre = tmp
    this.size++
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index >= this.size || index < 0 {
        return
    }
    node := this.head.Next
    for i := 0;i < index;i++{
        node = node.Next
    }
    node.Pre.Next = node.Next
    node.Next.Pre = node.Pre
    this.size--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4. 总结：
4.1: 理论 + 实践，不要空想。 链表想不清楚画画图，里面的边界 case, 会让你头疼的
