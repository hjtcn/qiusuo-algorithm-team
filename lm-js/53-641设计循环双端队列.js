// 641. 设计循环双端队列
// 设计实现双端队列。
// 你的实现需要支持以下操作：

// MyCircularDeque(k)：构造函数,双端队列的大小为k。
// insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
// insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
// deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
// deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
// getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
// getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
// isEmpty()：检查双端队列是否为空。
// isFull()：检查双端队列是否满了。
// 示例：

// MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
// circularDeque.insertLast(1);			        // 返回 true
// circularDeque.insertLast(2);			        // 返回 true
// circularDeque.insertFront(3);			        // 返回 true
// circularDeque.insertFront(4);			        // 已经满了，返回 false
// circularDeque.getRear();  				// 返回 2
// circularDeque.isFull();				        // 返回 true
// circularDeque.deleteLast();			        // 返回 true
// circularDeque.insertFront(4);			        // 返回 true
// circularDeque.getFront();				// 返回 4
 
 

// 提示：

// 所有值的范围为 [1, 1000]
// 操作次数的范围为 [1, 1000]
// 请不要使用内置的双端队列库。

/*
 * @lc app=leetcode.cn id=641 lang=javascript
 *
 * [641] 设计循环双端队列
 */

// @lc code=start
/**
 * Initialize your data structure here. Set the size of the deque to be k.
 * @param {number} k
 */
var MyCircularDeque = function(k) {
    this.queue=[]
    this.len=k+1
    //头部指向第一个指向的位置，从头部增加：this.head-1(求余先减k)，头部删除：this.head+1(求余)
    this.head=0
    //尾部下一个指向的位置，从尾部增加：this.tail+1(求余),尾部删除：this.tail-1(求余先加k)
    this.tail=0
};

/**
 * Adds an item at the front of Deque. Return true if the operation is successful. 
 * @param {number} value
 * @return {boolean}
 */
MyCircularDeque.prototype.insertFront = function(value) {
    if(this.isFull()){
        return false
    }
    this.head=(this.head-1+this.len)%this.len
    this.queue[this.head]=value
    return true
};

/**
 * Adds an item at the rear of Deque. Return true if the operation is successful. 
 * @param {number} value
 * @return {boolean}
 */
MyCircularDeque.prototype.insertLast = function(value) {
    if(this.isFull()){
        return false
    }
    this.queue[this.tail]=value
    this.tail=(this.tail+1)%this.len
    return true
};

/**
 * Deletes an item from the front of Deque. Return true if the operation is successful.
 * @return {boolean}
 */
MyCircularDeque.prototype.deleteFront = function() {
    if(this.isEmpty()){
        return false
    }
    this.head=(this.head+1)%this.len
    return true
};

/**
 * Deletes an item from the rear of Deque. Return true if the operation is successful.
 * @return {boolean}
 */
MyCircularDeque.prototype.deleteLast = function() {
    if(this.isEmpty()){
        return false
    }
    this.tail=(this.tail-1+this.len)%this.len
    return true
};

/**
 * Get the front item from the deque.
 * @return {number}
 */
MyCircularDeque.prototype.getFront = function() {
    //这里踩了下坑，下意识的获取队列头，忘记了头指针，而是直接返回this.queue[0]了
    return this.isEmpty()?-1:this.queue[this.head]
};

/**
 * Get the last item from the deque.
 * @return {number}
 */
MyCircularDeque.prototype.getRear = function() {
    return this.isEmpty()?-1:this.queue[(this.tail-1+this.len)%this.len]
};

/**
 * Checks whether the circular deque is empty or not.
 * @return {boolean}
 */
MyCircularDeque.prototype.isEmpty = function() {
    //头尾指针相等，队列为空
    return this.head==this.tail
};

/**
 * Checks whether the circular deque is full or not.
 * @return {boolean}
 */
MyCircularDeque.prototype.isFull = function() {
    //尾指针+1对k求余，转一圈追上头指针
    return this.head==(this.tail+1)%this.len
};

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * var obj = new MyCircularDeque(k)
 * var param_1 = obj.insertFront(value)
 * var param_2 = obj.insertLast(value)
 * var param_3 = obj.deleteFront()
 * var param_4 = obj.deleteLast()
 * var param_5 = obj.getFront()
 * var param_6 = obj.getRear()
 * var param_7 = obj.isEmpty()
 * var param_8 = obj.isFull()
 */
// @lc code=end

    /**
    头尾指针，做这道题的原则就是放手给指针和对应属性方法。忘记queue[0]或者queue.length
    难点在于头尾增删的处理。  
    头部增加：this.head-1(求余先减k)，头部删除：this.head+1(求余)
    尾部增加：this.tail+1(求余),尾部删除：this.tail-1(求余先加k)
 
    复杂度分析：
    时间复杂度：O(1)
    调用相应的方法，直接判断，进行相应的处理，返回相应的结果
    空间复杂度：O(N)
    定义数组模拟队列。
   */