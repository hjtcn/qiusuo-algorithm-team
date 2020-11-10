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
    this.head=-1
    this.tail=-1
};

/**
 * Insert an element into the circular queue. Return true if the operation is successful. 
 * @param {number} value
 * @return {boolean}
 */
//插入操作
MyCircularQueue.prototype.enQueue = function(value) {
    //队列已满无法插入
    if(this.isFull()){
        return false
    }
    //队列为空，头指针为0
    if(this.isEmpty()){
        this.head=0
    }
    //尾指针后移，因为是循环队列，则+1对k求余
    this.tail=(this.tail+1)%this.len
    //队列赋值
    this.queue[this.tail]=value
    return true
};

/**
 * Delete an element from the circular queue. Return true if the operation is successful.
 * @return {boolean}
 */
//从头部删除
MyCircularQueue.prototype.deQueue = function() {
    //队列为空无法删除
    if(this.isEmpty()){
        return false
    }
    //头尾指针相同，说明原队列只有一个元素，删除后，即队列为空，将头尾指针置为-1
    if(this.head==this.tail){
        this.head=this.tail=-1
    }
    else{
    //从头部删除，头指针后移，循环队列，故头指针加一后对k求余
        this.head=(this.head+1)%this.len
    }
    return true
};

/**
 * Get the front item from the queue.
 * @return {number}
 */
//获取队首
MyCircularQueue.prototype.Front = function() {
    //队列不为空，返回头指针指向的元素
    return this.isEmpty()?-1:this.queue[this.head]
};

/**
 * Get the last item from the queue.
 * @return {number}
 */
//获取队尾
MyCircularQueue.prototype.Rear = function() {
    //队列不为空，返回尾指针指向的元素
    return this.isEmpty()?-1:this.queue[this.tail]
};

/**
 * Checks whether the circular queue is empty or not.
 * @return {boolean}
 */
//队列是否为空
MyCircularQueue.prototype.isEmpty = function() {
    //当从头部删除deQueue的时候，如果队列被清空，会将头指针置为-1。
    return this.head==-1
};

/**
 * Checks whether the circular queue is full or not.
 * @return {boolean}
 */
//队列是否满了
MyCircularQueue.prototype.isFull = function() {
    //如果尾指针+1对长度求余追上头指针，则队列已满
    return this.head==(this.tail+1)%this.len
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
    第一次没看清题是用数组的api写的，后来看到不让用内置api的时候有点懵了，有点产生思维习惯了。
    后来认真看了题解，借用头尾指针来获取队列的状态。注意循环队列，要对队列长度k求余
    
    复杂度分析：
	时间复杂度：O(1)
    调用方法，处理头尾指针，并对数组进行赋值模拟队列操作
	空间复杂度：O(N)
	提前声明数组模拟队列的操作
 */
