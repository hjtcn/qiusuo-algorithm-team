/*
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
circularDeque.insertLast(1);           // 返回 true
circularDeque.insertLast(2);           // 返回 true
circularDeque.insertFront(3);           // 返回 true
circularDeque.insertFront(4);           // 已经满了，返回 false
circularDeque.getRear();      // 返回 2
circularDeque.isFull();            // 返回 true
circularDeque.deleteLast();           // 返回 true
circularDeque.insertFront(4);           // 返回 true
circularDeque.getFront();    // 返回 4
*/
 
分析：
循环双端队列与循环队列区别就是双端：
循环队列还是只能尾部进，头部出
循环双端队列可以尾部进尾部出，头部进头部出
判断队列为空的条件：
rear == front
判断队列满的条件：
(rear + 1) % size == front
所以结构体的设计为
type MyCircularDeque struct {
    data[] int,  // 存储数据
    front,rear int,  // 头指针和尾部指针
    size int,  // 队列大小
}
插入操作不同：
如果使用数组实现的话，向头部插入的时候，需要移动其余的数组元素，感觉不是很合适，时间复杂度会比较高
比较适合用链表来实现，如果用链表来实现的话，存储数据如何设计呢？
尝试设计一下呢
节点结构：
typedef struct QNode 
{
    
}
发现到这里思维又发散了，想不出来了，这里是对链表还是不熟悉？
看了题解发现，使用数组实现的话，向头部插入的话，并不需要移动后对面的元素，只需要改变front 的位置就可以了

type MyCircularDeque struct {
    data[] int
    front,rear int
    size int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        data:make([]int,k + 1),
        front:0,
        rear:0,
        size:k+1,
    }
}
/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }
    this.front = (this.front - 1 + this.size) % this.size
    this.data[this.front] = value
    return true
}
/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }
    this.data[this.rear] = value
    this.rear = (this.rear + 1) % this.size
    return true
}
/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    this.front = (this.front + 1) % this.size
    return true
}
/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    this.rear = (this.rear - 1 + this.size) % this.size
    return true
}
/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.data[this.front]
}
/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    index := (this.rear - 1 + this.size) % this.size
    return this.data[index]
}
/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    if this.front == this.rear {
        return true
    }
    return false
}
/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    if (this.rear + 1) % this.size == this.front {
        return true
    }
    return false
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

自己出错的地方在于向头部插入数据的时候
this.front = (this.front - 1 + this.size) % this.size
如果当前位置下标索引为0,0-1=-1,我们不想得到-1，想得到数组最后下标的位置，比如k=3,那么 size=4, (-1 + 4)%4 = 3，这个时候谁队列满的时候，当然也有可能是可以插入的，这个时候需要判断 rear 的位置是在那里

复杂度分析：
时间复杂度：O(1): 数组可以实现下标随机访问，所以时间复杂度为O(1)
空间复杂度：O(n): 使用数组存储元素O(n)

如果用链表实现的话，怎么实现呢？怎么设计呢？不会哈，去参考了题解：https://leetcode-cn.com/problems/design-circular-deque/solution/golang-lian-biao-by-sunhongyue4500/

链表思路：使用head 指针指向第一个节点，最后一个节点后是尾节点
自己原来疑惑的地方，链表实现的话怎么限制长度呢？ -> 我们可以使用一个变量 count 记录呀，如果当前实现不了，我们可以找方法，比如找个变量记录哈

type MyCircularDeque struct {
    head *Node
    tail *Node
    len int
    count int
}
type Node struct {
    Next *Node
    Pre *Node
    Val int
}
/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    head := Node {
        Next:nil,
        Pre:nil,
        Val:-1,
    }
    tail := Node {
        Next:nil,
        Pre:nil,
        Val:-1,
    }
    head.Next = &tail
    tail.Pre = &head
    deque := MyCircularDeque{
        head:&head,
        tail:&tail,
        len:k,
        count:0,
    }
    return deque
}
/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }
    temp := this.head.Next
    tempNode := Node {
        Next:temp,
        Pre:this.head,
        Val:value,
    }
    this.head.Next = &tempNode
    temp.Pre = &tempNode
    this.count++
    return true
}
/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }
    temp := this.tail.Pre
    tempNode := Node {
        Next: this.tail,
        Pre: temp,
        Val: value,
    }
    this.tail.Pre = &tempNode
    temp.Next = &tempNode
    this.count++
    return true
}
/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    deleteTemp := this.head.Next
    this.head.Next = deleteTemp.Next
    deleteTemp.Next.Pre = this.head
    deleteTemp.Next,deleteTemp.Pre = nil,nil
    this.count--
    return true
}
/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return  false
    }
    deleteTemp := this.tail.Pre
    deleteTemp.Pre.Next = this.tail
    this.tail.Pre = deleteTemp.Pre
    deleteTemp.Next,deleteTemp.Pre = nil,nil
    this.count--
    return true
}
/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.head.Next.Val
}
/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.tail.Pre.Val
}
/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    return this.head.Next == this.tail
}
/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    return this.len == this.count
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

复杂度分析：
时间复杂度：O(1)
空间复杂度：O(n), 比数组占用空间会多一点，因为head 和tail 节点占用空间多一点，基本可以忽略

总结：
1. 思路不是很难，自己做的时候还是很多千奇百怪的问题，说明自己还是没有真正理解，也说明自己动手实践，如果没有动手实践的东西和得到反馈的东西是真的会吗？我想不然，应该还是不会的

2. 自己害怕链表实现，为什么害怕呢？-》深层次的原因还是自己不会，逃避是没有用的，只有自己去面对它，学会了，下次遇到就不害怕了

3. Easy Said Than Done, 去一点点做吧，定期总结，反馈，不要害怕失败，不要怕题目不会写
