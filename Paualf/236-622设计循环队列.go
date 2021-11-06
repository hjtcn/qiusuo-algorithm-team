设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。

循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。

你的实现应该支持如下操作：
MyCircularQueue(k): 构造器，设置队列长度为 k 。
Front: 从队首获取元素。如果队列为空，返回 -1 。
Rear: 获取队尾元素。如果队列为空，返回 -1 。
enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
isEmpty(): 检查循环队列是否为空。
isFull(): 检查循环队列是否已满。
 
示例：
MyCircularQueue circularQueue = new MyCircularQueue(3); // 设置长度为 3
circularQueue.enQueue(1);  // 返回 true
circularQueue.enQueue(2);  // 返回 true
circularQueue.enQueue(3);  // 返回 true
circularQueue.enQueue(4);  // 返回 false，队列已满
circularQueue.Rear();  // 返回 3
circularQueue.isFull();  // 返回 true
circularQueue.deQueue();  // 返回 true
circularQueue.enQueue(4);  // 返回 true
circularQueue.Rear();  // 返回 4
 
提示：
所有的值都在 0 至 1000 的范围内；
操作数将在 1 至 1000 的范围内；
请不要使用内置的队列库。

1. Clarification:
Q: 什么时候队列为空？
A: 不知道哈，纸上画了下，没画出来

Q: 什么时候队列已满？
A: xxx

Q:队首队尾怎么表示？
A: xxx

Q: 初始化的时候 front 是否等于 rear?
A: xxx

2. 看题解：
对于一个固定大小的数组
知道队列长度，可以知道队尾
tailIndex = (headIndex + count - 1) mod capacity
题解中还给了多个case

设计数据结构的关键是如何设计属性，好的设计属性数量更少
属性数量少说明属性之间冗余更低
属性冗余度越低，操作逻辑越简单，发生错误的可能性更低。
属性数量少，使用的空间也少，操作性能高。

queue: 一个固定大小的数组，用于保存循环队列的元素
headIndex: 一个整数，保存队首head的索引
count: 循环队列当前的长度，即循环队列中的元素数量
capacity: 循环队列的容量，即队列中最多可以容纳的元素数量

type MyCircularQueue struct {
    Queue []int
    Count int
    Capacity int
    HeadIndex int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        Queue:make([]int,k),
        Count:0,
        Capacity:k,
        HeadIndex:0,
    }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    // 如果满了则推出
    if this.IsFull() {
        return false
    }
    // 这个地方调试了半天
    this.Queue[(this.HeadIndex + this.Count) % this.Capacity] = value
    this.Count++
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
        return false
    }
    this.HeadIndex = (this.HeadIndex + 1) %this.Capacity
    this.Count--
    return true
}


func (this *MyCircularQueue) Front() int {
    if this.IsEmpty(){
        return -1
    }
    return this.Queue[this.HeadIndex]
}


func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
        return -1
    }
    rear := (this.HeadIndex + this.Count - 1) % this.Capacity
    return this.Queue[rear]
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.Count == 0
}


func (this *MyCircularQueue) IsFull() bool {
    return this.Count == this.Capacity 
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
  // 这个地方调试了半天
    this.Queue[(this.HeadIndex + this.Count) % this.Capacity] = value

有的时候以为自己理解了，其实并没有真正的理解

实践出真知

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1: 当你在纸上画不清楚它的实现细节的时候，有可能你对它并没有真正的掌握和理解

4.2: 读书或者学习是为了理解，而不是仅仅是我知道有这个东西，后面还是不会。。。
