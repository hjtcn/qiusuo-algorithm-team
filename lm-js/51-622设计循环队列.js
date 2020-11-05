// 622. 设计循环队列
// 设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。

// 循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。

// 你的实现应该支持如下操作：

// MyCircularQueue(k): 构造器，设置队列长度为 k 。
// Front: 从队首获取元素。如果队列为空，返回 -1 。
// Rear: 获取队尾元素。如果队列为空，返回 -1 。
// enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
// deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
// isEmpty(): 检查循环队列是否为空。
// isFull(): 检查循环队列是否已满。
 

// 示例：

// MyCircularQueue circularQueue = new MyCircularQueue(3); // 设置长度为 3
// circularQueue.enQueue(1);  // 返回 true
// circularQueue.enQueue(2);  // 返回 true
// circularQueue.enQueue(3);  // 返回 true
// circularQueue.enQueue(4);  // 返回 false，队列已满
// circularQueue.Rear();  // 返回 3
// circularQueue.isFull();  // 返回 true
// circularQueue.deQueue();  // 返回 true
// circularQueue.enQueue(4);  // 返回 true
// circularQueue.Rear();  // 返回 4
 

// 提示：

// 所有的值都在 0 至 1000 的范围内；
// 操作数将在 1 至 1000 的范围内；
// 请不要使用内置的队列库。


/*
 * @lc app=leetcode.cn id=622 lang=javascript
 *
 * [622] 设计循环队列
 */

// @lc code=start
/**
 * Initialize your data structure here. Set the size of the queue to be k.
 * @param {number} k
 */
var MyCircularQueue = function(k) {
    this.queue=[]
    this.len=k
};

/**
 * Insert an element into the circular queue. Return true if the operation is successful. 
 * @param {number} value
 * @return {boolean}
 */
//插入操作
MyCircularQueue.prototype.enQueue = function(value) {
    if(this.queue.length<this.len){
        this.queue.push(value)
        return true
    }
    return false
};

/**
 * Delete an element from the circular queue. Return true if the operation is successful.
 * @return {boolean}
 */
//从头部删除
MyCircularQueue.prototype.deQueue = function() {
    if(this.queue&&this.queue.length){
        this.queue.shift()
        return true
    }
    return false
};

/**
 * Get the front item from the queue.
 * @return {number}
 */
//获取队首
MyCircularQueue.prototype.Front = function() {
    //这里曾经踩过坑
    return this.queue.length?this.queue[0]:-1
};

/**
 * Get the last item from the queue.
 * @return {number}
 */
//获取队尾
MyCircularQueue.prototype.Rear = function() {

    return this.queue.length?this.queue[this.queue.length-1]:-1
};

/**
 * Checks whether the circular queue is empty or not.
 * @return {boolean}
 */
//队列是否为空
MyCircularQueue.prototype.isEmpty = function() {
    return !this.queue.length
};

/**
 * Checks whether the circular queue is full or not.
 * @return {boolean}
 */
//队列是否满了
MyCircularQueue.prototype.isFull = function() {
    return this.queue.length==this.len
};

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * var obj = new MyCircularQueue(k)
 * var param_1 = obj.enQueue(value)
 * var param_2 = obj.deQueue()
 * var param_3 = obj.Front()
 * var param_4 = obj.Rear()
 * var param_5 = obj.isEmpty()
 * var param_6 = obj.isFull()
 */
// @lc code=end


/** 题解
    这种题，理解队列以及js的对象结构，一切还是比较清晰的。
    但是过程中踩了一个小坑。比如返回队首，我一开始用的返回this.queue[0]||-1，
    踩坑了如果this.queue[0]为0的情况。最后还是根据队列长度乖乖判断
    复杂度分析：
	时间复杂度：O(1)

	空间复杂度：O(N)
	提前声明数组模拟队列的操作
 */
