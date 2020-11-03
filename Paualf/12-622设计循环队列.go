/*
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
*/

设计循环队列：
循环队列队空标记： rear== front
循环队列队满标记： rear + 1 % size == front
循环队列长度：(rear- front  + size ) % size

type MyCircularQueue struct {
    queue[] int
    front,rear int
    size int
}
/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
    // 结构体初始化
    return MyCircularQueue{
        queue: make([]int, k+1),
        front:0,
        rear:0,
        size:k+1,
    }
}
/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
        return false
    }
    //fmt.Println(value)
    this.queue[this.rear] = value
    //fmt.Println(this.rear,this.queue)
    this.rear = (this.rear + 1) % this.size
    return true
}
/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
        return false
    }
    this.front = (this.front + 1) % this.size
    return true
}
/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
        return -1
    }
    return this.queue[this.front]
}
/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
        return -1
    }
    lastIndex := this.rear - 1
    if lastIndex < 0 {
        lastIndex = this.size - 1
    }
    return this.queue[lastIndex]
}
/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
    if this.rear == this.front {
        return true
    }
    return false
}
/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
    if (this.rear + 1) % this.size == this.front {
        return true
    }
    return false
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

如果了解队满和队空的条件的话代码是不难实现的，如果不了解的话，还是有点懵懵的，以前不了解，写起来感觉像是在死记硬背的感觉
结果我在leetcode 编辑器里面运行结果是可以走到4里面的，return false 以后仍然是可以执行下去的，我的天哈，这个地方让我费解了好久
1
0 [1 0 0 0]
2
1 [1 2 0 0]
3
2 [1 2 3 0]
4
4
3 [1 2 3 4]

然后我同样代码在本地运行发现return false 以后就不走下去了,差点没有被坑死。。。
1
0 [1 0 0 0]
2
1 [1 2 0 0]
3
2 [1 2 3 0]
4

总结：
1. 理论挺重要的，如果不了解循环队列队空和队满的状态，代码还是不容易实现的，理解理论然后动手实践

2. leetcode的编辑器还是挺神奇的，卡了我半天。。。

3. 做一名合格的工程师，把每一行代码尽可能的写好
