设计实现双端队列。

你的实现需要支持以下操作：
MyCircularDeque(k)：构造函数,双端队列的大小为k。
insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
isEmpty()：检查双端队列是否为空。
isFull()：检查双端队列是否满了。

示例：
MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
circularDeque.insertLast(1);			        // 返回 true
circularDeque.insertLast(2);			        // 返回 true
circularDeque.insertFront(3);			        // 返回 true
circularDeque.insertFront(4);			        // 已经满了，返回 false
circularDeque.getRear();  				// 返回 2
circularDeque.isFull();				        // 返回 true
circularDeque.deleteLast();			        // 返回 true
circularDeque.insertFront(4);			        // 返回 true
circularDeque.getFront();				// 返回 4
 
提示：
所有值的范围为 [1, 1000]
操作次数的范围为 [1, 1000]
请不要使用内置的双端队列库。

1. Clarification:
如何存储：
 	底层如果用数组的话，使用下标标记循环链表头部和尾部的话，向尾部插入，感觉挺怪的。也不是很合理，所以这里最好是使用链表的方式进行存储。

Q： 使用单链表的话，方便插入头尾节点，需要使用 head 和 tail 两个头尾节点
A： 这个时候向尾部插入的话，新的节点的 Next = tail，这个时候怎么和tail 节点的前一个节点连接上呢？这个时候发现连接不上


Q:  单链表好像走不通，为了将tail节点的前一个节点给连接上，我们需要使用双向链表进行连接选择
A： 双向链表纸上画一波，是可以走通的

type ListNode struct {
    Pre *ListNode
    Next *ListNode
    Val int
}

type MyCircularDeque struct {
    Head *ListNode
    Tail *ListNode
    CurrSize int
    Size int
}


func Constructor(k int) MyCircularDeque {
    head := &ListNode{}
    tail := &ListNode{}
    head.Next = tail
    tail.Pre = head
    return MyCircularDeque {
        Head:head,
        Tail:tail,
        CurrSize:0,
        Size:k,
    }
}

// 双向链表四根线，先接涉及到后面的节点的，因为后面的会被断掉
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }
    newNode := &ListNode{
        Val:value,
    }
    this.Head.Next.Pre = newNode
    newNode.Next = this.Head.Next
    this.Head.Next = newNode
    newNode.Pre = this.Head
    this.CurrSize += 1
    return true
}

// 四根线，先管tail前面的节点信息
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }
    newNode := &ListNode{
        Val:value,
    }
    newNode.Pre = this.Tail.Pre
    this.Tail.Pre.Next = newNode
    this.Tail.Pre = newNode
    newNode.Next = this.Tail
    this.CurrSize += 1
    return true
}


func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    this.Head.Next.Next.Pre = this.Head
    this.Head.Next = this.Head.Next.Next
    this.CurrSize -= 1
    return true
}


func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    // 这两句一开始写反了，先写这个的话 会 panic this.Tail.Pre = this.Tail.Pre.Pre
    this.Tail.Pre.Pre.Next = this.Tail
    this.Tail.Pre = this.Tail.Pre.Pre
    this.CurrSize -= 1
    return true
}


func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.Head.Next.Val
}


func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.Tail.Pre.Val
}


func (this *MyCircularDeque) IsEmpty() bool {
    return this.CurrSize == 0
}


func (this *MyCircularDeque) IsFull() bool {
    return this.CurrSize >= this.Size
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

2. 看题解：
Q: 自己当时为什么没有用数组去存储呢？
A: 可能是自己当时没有想清楚下标是怎么移动的，看了题解是 front 和 rear 的变化

为了避免队列为空和队列为满的判别条件有冲突，有意浪费一个位置

浪费一个位置是指：循环数组中任何时刻一定至少有一个位置不存放有效元素
判断队列为空的条件是： front == rear
判断队列为满的条件是： (rear + 1) % capacity == front 

因为有循环的存在，要特别注意下标可能越界的情况
指针后移：下标 + 1，要去模
指针前移：为了循环到数组的末尾，要加上数组的长度，然后再对数组长度取模

3. 复杂度分析：
时间复杂度：O(1)
空间复杂度：O(1)

4. 总结：
4.1: 仔细分析哈，画画图，理解下
