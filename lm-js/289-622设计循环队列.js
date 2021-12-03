/*
622. 设计循环队列
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


*/

/*
    思路：其实感觉自己思考的不是循环队列，因为把链表弄成环实在是心里有恐惧，我就先试着去实现它，用链表的形式。更新头尾节点，记录当前链表长度，搞清楚各种边界值
    踩坑：头尾节点的初始化踩坑了，我直接给它们赋值为空节点，最后判断是否含有头节点的时候，又根据是否有next节点判断。因此不会返回-1了。

    题解：看题解，很多也都没有将this.tail.next=this.head,队尾连在队首
*/

/**
 * @param {number} k
 */
 function NodeList(val){
    this.val=val
    this.next=null
}
var MyCircularQueue = function(k) {
    this.size=k
    this.count=0
    this.head=null
    this.tail=null
};

/** 
 * @param {number} value
 * @return {boolean}
 */
MyCircularQueue.prototype.enQueue = function(value) {
    if(this.count>=this.size) return false
    this.count++;
    if(this.count==1) 
    {
        this.head=this.tail=new NodeList(value)
        return true
    }
    let prev=this.tail
    this.tail=new NodeList(value)
    prev.next=this.tail
    return true
};

/**
 * @return {boolean}
 */
MyCircularQueue.prototype.deQueue = function() {
    if(this.count<=0) return false 
    let head=this.head
    if(head.next){
        this.head=head.next
    }
    else{
        this.head=null
        this.tail=null
    }
    this.count--
    return true
};

/**
 * @return {number}
 */
MyCircularQueue.prototype.Front = function() {
    if(this.head) return this.head.val
    return -1
};

/**
 * @return {number}
 */
MyCircularQueue.prototype.Rear = function() {
    if(this.tail) return this.tail.val
    return -1
};

/**
 * @return {boolean}
 */
MyCircularQueue.prototype.isEmpty = function() {
    if(this.count==0) return true
    return false
};

/**
 * @return {boolean}
 */
MyCircularQueue.prototype.isFull = function() {
    if(this.count==this.size) return true
    return false
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

/*
    时间复杂度：O(1)
    空间复杂度：O(N)
*/

/*
    思考：先解决问题，再思考优化
*/