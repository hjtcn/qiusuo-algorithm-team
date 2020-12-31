<?php

/*
 * @lc app=leetcode.cn id=622 lang=php
 *
 * [622] 设计循环队列
 *
 * https://leetcode-cn.com/problems/design-circular-queue/description/
 *
 * algorithms
 * Medium (41.64%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    41.2K
 * Total Submissions: 98.8K
 * Testcase Example:  '["MyCircularQueue","enQueue","enQueue","enQueue","enQueue","Rear","isFull","deQueue","enQueue","Rear"]\n' +
  '[[3],[1],[2],[3],[4],[],[],[],[4],[]]'
 *
 * 设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于
 * FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。
 * 
 * 
 * 循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。
 * 
 * 你的实现应该支持如下操作：
 * 
 * 
 * MyCircularQueue(k): 构造器，设置队列长度为 k 。
 * Front: 从队首获取元素。如果队列为空，返回 -1 。
 * Rear: 获取队尾元素。如果队列为空，返回 -1 。
 * enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
 * deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
 * isEmpty(): 检查循环队列是否为空。
 * isFull(): 检查循环队列是否已满。
 * 
 * 
 * 
 * 
 * 示例：
 * 
 * MyCircularQueue circularQueue = new MyCircularQueue(3); // 设置长度为 3
 * circularQueue.enQueue(1);  // 返回 true
 * circularQueue.enQueue(2);  // 返回 true
 * circularQueue.enQueue(3);  // 返回 true
 * circularQueue.enQueue(4);  // 返回 false，队列已满
 * circularQueue.Rear();  // 返回 3
 * circularQueue.isFull();  // 返回 true
 * circularQueue.deQueue();  // 返回 true
 * circularQueue.enQueue(4);  // 返回 true
 * circularQueue.Rear();  // 返回 4
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 所有的值都在 0 至 1000 的范围内；
 * 操作数将在 1 至 1000 的范围内；
 * 请不要使用内置的队列库。
 * 
 * 
 */

// @lc code=start
class MyCircularQueue {

    private $front    = 0;   // 队列首部
    private $rear     = 0;   // 队列尾部的下一个位置
    private $capacity = 0;   // 队列长度
    private $data     = [];  // 队列

    /**
     * Initialize your data structure here. Set the size of the queue to be k.
     * @param Integer $k
     */
    function __construct($k) {
        $this->capacity = $k+1; // 多空一个位置
    }
  
    /**
     * Insert an element into the circular queue. Return true if the operation is successful.
     * @param Integer $value
     * @return Boolean
     */
    // 入队
    function enQueue($value) {
        if($this->isFull()) {
            return false; // 如果已经满了就返回false
        }
        // rear指向队尾下一个位置
        $this->data[$this->rear] = $value;
        $this->rear = ($this->rear + 1) % $this->capacity;
        return true;
    }
  
    /**
     * Delete an element from the circular queue. Return true if the operation is successful.
     * @return Boolean
     */
    // 出队
    function deQueue() {
        if($this->isEmpty()) {
            return false;
        }
        // 头指针向前移动一下
        $this->front = ($this->front + 1) % $this->capacity;
        return true;
    }
  
    /**
     * Get the front item from the queue.
     * @return Integer
     */
    // 返回队首
    function Front() {
        if($this->isEmpty()) {
            return -1;
        }else{
            return $this->data[$this->front]; // 队首元素
        }
    }
  
    /**
     * Get the last item from the queue.
     * @return Integer
     */
    function Rear() {
        if($this->isEmpty()) {
            return -1;
        }else{
            return $this->data[ ($this->rear - 1 + $this->capacity) % $this->capacity];
        }
    }
  
    /**
     * Checks whether the circular queue is empty or not.
     * @return Boolean
     */
    function isEmpty() {
        // 队列为空判定条件
        return $this->front == $this->rear;
    }
  
    /**
     * Checks whether the circular queue is full or not.
     * @return Boolean
     */
    function isFull() {
        // 队列已满判定条件
        return ($this->rear + 1) % $this->capacity == $this->front;
    }
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * $obj = MyCircularQueue($k);
 * $ret_1 = $obj->enQueue($value);
 * $ret_2 = $obj->deQueue();
 * $ret_3 = $obj->Front();
 * $ret_4 = $obj->Rear();
 * $ret_5 = $obj->isEmpty();
 * $ret_6 = $obj->isFull();
 */
// @lc code=end

/*
采用数组动态模拟循环队列

front 指向队列头部第 1 个有效数据的位置
rear  指向队列尾部（即最后 1 个有效数据）的下一个位置，即下一个从队尾入队元素的位置。

判别队列为空的条件是 front == rear
判别队列为满的条件是 (rear + 1) % capacity == front;

- 时间复杂度 O(1)
- 空间复杂度 O(N)

参考：https://leetcode-cn.com/problems/design-circular-queue/solution/shu-zu-shi-xian-de-xun-huan-dui-lie-by-liweiwei141/
*/