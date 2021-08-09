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
初始化：
  Length = 0
定义 LinkNode {
Val int
Next *LinkNode
}
AddAtHead() {
newNode := &LinkNode{
Val:val
}
head := this.Head
newNode.Next = head.Next 
head.Next = newNode
this.Length++
}
AddAtTail() {
    newNode := &LinkNode{
Val : val
   }  
    head := this.head
for i := 0;i < this.Length;i++ {
head = head.Next
}
head.Next = newNode
}
AddAtIndex() {
if index > this.Len {
return
}
if index == len{
this.AddAtTail
}
if index <= 0 {
this.AddAtHead()
}
}

DeleteAtIndex() {
head := this.head
for i := 0;i < index;i++ {
head = head.Next
}

head.Next = head.Next.Next
}
type MyLinkedList struct {
    Length int
    Head *LinkNode
}

type LinkNode struct {
    Val int
    Next *LinkNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    return MyLinkedList{
        Length:0,
    }
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    if index > this.Length {
        return -1
    }
    head := this.Head
    for i := 0;i <= index;i++{
        head = head.Next
    }
    return head.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
    newNode := &LinkNode {
        Val:val,
    }
    head := this.Head
    newNode.Next = head.Next
    head.Next = newNode
    this.Length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    newNode := &LinkNode {
        Val:val,
    }
    head := this.Head
    for i := 0;i < this.Length;i++{
        head = head.Next
    }
    head.Next = newNode
    this.Length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index > this.Length {
        return
    }
    if index == this.Length {
        this.AddAtTail(val)
        return
    }
    if index <= 0 {
        this.AddAtHead(val)
        return
    }
    newNode := &LinkNode{
        Val:val,
    }
    head := this.Head
    for i := 0;i < index;i++ {
        head = head.Next
    }
    newNode.Next = head.Next.Next
    head.Next = newNode
    this.Length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index > this.Length {
        return
    }
    head := this.Head
    for i := 0;i < index;i++ {
        head = head.Next
    }
    head.Next = head.Next.Next
    this.Length--
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

 疑问点是:
newNode.Next newNode is nil ?
    newNode := &LinkNode {
        Val:val,
    }
    head := this.Head
    newNode.Next = head.Next
debug 了半天：还对比着以前的代码才调通了。。。

type MyLinkedList struct {
    Length int
    Head *LinkNode
}

type LinkNode struct {
    Val int
    Next *LinkNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    return MyLinkedList{
        Length:0,
        Head:&LinkNode{
        },
    }
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.Length {
        return -1
    }
    head := this.Head
    for i := 0;i < index;i++{
        head = head.Next
    }
    return head.Next.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
     head := this.Head
     newNode := &LinkNode {
        Val:val,
        Next:head.Next,
    }    
    head.Next = newNode
    this.Length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    head := this.Head
    for i := 0;i < this.Length;i++{
        head = head.Next
    }
    newNode := &LinkNode {
        Val:val,
        Next:head.Next,
    }
    head.Next = newNode
    this.Length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index > this.Length {
        return
    }
    if index == this.Length {
        this.AddAtTail(val)
        return
    }
    if index <= 0 {
        this.AddAtHead(val)
        return
    }
    newNode := &LinkNode{
        Val:val,
    }
    head := this.Head
    for i := 0;i < index;i++ {
        head = head.Next
    }
    newNode.Next = head.Next
    head.Next = newNode
    this.Length++
  //  this.Print()
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= this.Length {
        return
    }
    head := this.Head
    for i := 0;i < index;i++ {
        head = head.Next
    }
    head.Next = head.Next.Next
    this.Length--
}

func (this *MyLinkedList)Print() {
    head := this.Head
    for i := 0;i < this.Length;i++{
        head = head.Next
        if head != nil {
          fmt.Println(head.Val)
        }
    }
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
比较容易出错的地方在, 精髓，因为你获取 head.Next 的时候你要保证它存在，比较容易出错
if index >= this.Length {
}

2. 看题解：
几个月前提交的

func Constructor() MyLinkedList {
    return MyLinkedList{
        size:0,
        head:&ListNode{},
    }
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1 
    }
    pre := this.head
    for i := 0;i < index;i++ {
        pre = pre.Next
    }
    return pre.Next.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
    this.AddAtIndex(0,val)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    this.AddAtIndex(this.size,val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index < 0 || index > this.size {
        return
    }
    pre := this.head
    for i := 0;i < index;i++ {
        pre = pre.Next
    }
    node := &ListNode{
        Val:val,
        Next:pre.Next,
    }
    pre.Next = node
    this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= this.size {
        return
    }
    pre := this.head
    for i := 0;i < index;i++ {
        pre = pre.Next
    }
    pre.Next = pre.Next.Next
    this.size--
}

3. 复杂度分析：
时间复杂度：向表头插入O(1),get(index):O(index),addAtTail:O(n),deleteAtIndex:O(index)
空间复杂度：O(1)

4. 总结：
4.1: 链表真的都是细节，设计重于实现，说起来容易，做起来真的有点难

4.2: debug的时候不好debug自己去写一个小工具去调试

4.3: 链表这种题，思路大家都知道，但是真正能写出来的很少的，所以动手去实践，不要害怕失败，不要害怕 指针找不到
