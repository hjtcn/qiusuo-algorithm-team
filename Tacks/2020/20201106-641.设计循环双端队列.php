<?php

/*
 * @lc app=leetcode.cn id=641 lang=php
 *
 * [641] 设计循环双端队列
 *
 * https://leetcode-cn.com/problems/design-circular-deque/description/
 *
 * algorithms
 * Medium (53.29%)
 * Likes:    61
 * Dislikes: 0
 * Total Accepted:    12.7K
 * Total Submissions: 23.9K
 * Testcase Example:  '["MyCircularDeque","insertLast","insertLast","insertFront","insertFront","getRear","isFull","deleteLast","insertFront","getFront"]\n' +
  '[[3],[1],[2],[3],[4],[],[],[],[4],[]]'
 *
 * 设计实现双端队列。
 * 你的实现需要支持以下操作：
 * 
 * 
 * MyCircularDeque(k)：构造函数,双端队列的大小为k。
 * insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
 * insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
 * deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
 * deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
 * getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
 * getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
 * isEmpty()：检查双端队列是否为空。
 * isFull()：检查双端队列是否满了。
 * 
 * 
 * 示例：
 * 
 * MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
 * circularDeque.insertLast(1);                    // 返回 true
 * circularDeque.insertLast(2);                    // 返回 true
 * circularDeque.insertFront(3);                    // 返回 true
 * circularDeque.insertFront(4);                    // 已经满了，返回 false
 * circularDeque.getRear();                  // 返回 2
 * circularDeque.isFull();                        // 返回 true
 * circularDeque.deleteLast();                    // 返回 true
 * circularDeque.insertFront(4);                    // 返回 true
 * circularDeque.getFront();                // 返回 4
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 所有值的范围为 [1, 1000]
 * 操作次数的范围为 [1, 1000]
 * 请不要使用内置的双端队列库。
 * 
 * 
 */

// @lc code=start
class MyCircularDeque {

    protected $queue; // 一个数组来存储队列的元素
    protected $capacity;
    protected $head; // 队列头
    protected $tail; // 列尾的下一个位置


    /**
     * Initialize your data structure here. Set the size of the deque to be k.
     * @param Integer $k
     */
    function __construct($k) {
        $this->capacity = $k + 1; // 有意保留多一个元素位置
        $this->queue = array_fill(0, $this->capacity, null); // 初始化
        $this->head  = 0;
        $this->tail  = 0; // 下一个插入元素的位置

    }
  
    /**
     * Adds an item at the front of Deque. Return true if the operation is successful.
     * @param Integer $value
     * @return Boolean
     */
    // 从队列头部插入
    function insertFront($value) {
        // 判断队列是否已经满了
        if ($this->isFull()) {
            return false;
        }

        $this->head = ($this->head + $this->capacity - 1) % $this->capacity;
        $this->queue[$this->head] = $value;

        return true;
    }
  
    /**
     * Adds an item at the rear of Deque. Return true if the operation is successful.
     * @param Integer $value
     * @return Boolean
     */
    // 从队列尾部插入
    function insertLast($value) {
        // 判断队列是否已经满了
        if ($this->isFull()) {
            return false;
        }

        $this->queue[$this->tail] = $value;
        $this->tail = ($this->tail + 1) % $this->capacity; // 尾部

        return true;
    }
  
    /**
     * Deletes an item from the front of Deque. Return true if the operation is successful.
     * @return Boolean
     */
    // 删除队首元素
    function deleteFront() {
        if ($this->isEmpty()) {
            return false;
        }

        $this->head = ($this->head + 1) % $this->capacity; // 向后移动一位

        return true;
    }
  
    /**
     * Deletes an item from the rear of Deque. Return true if the operation is successful.
     * @return Boolean
     */
    // 删除队尾元素
    function deleteLast() {
        if ($this->isEmpty()) {
            return false;
        }

        $this->tail = ($this->tail - 1 + $this->capacity) % $this->capacity;
        return true;
    }
  
    /**
     * Get the front item from the deque.
     * @return Integer
     */
    // 获得队首元素
    function getFront() {
        if($this->isEmpty()) {
            return -1;
        }
        return $this->queue[$this->head];
    }
  
    /**
     * Get the last item from the deque.
     * @return Integer
     */
    // 获取队尾元素
    function getRear() {
        if($this->isEmpty()) {
            return -1;
        }
        return $this->queue[ ($this->tail - 1 + $this->capacity) % $this->capacity];
    }
  
    /**
     * Checks whether the circular deque is empty or not.
     * @return Boolean
     */
    // 是否队列为空
    function isEmpty() {
        return $this->head == $this->tail;
    }
  
    /**
     * Checks whether the circular deque is full or not.
     * @return Boolean
     */
    // 是否队列为满
    function isFull() {
        // 数组的容量为队列容量 + 1，这里有意浪费一个位置
        // 这样可以区分当 head == tail 的时候是空队列
        // 当 ( tail + 1 + capacity ) % capacity = head是满队列
        return ($this->tail + 1 + $this->capacity) % $this->capacity == $this->head;
    }
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * $obj = MyCircularDeque($k);
 * $ret_1 = $obj->insertFront($value);
 * $ret_2 = $obj->insertLast($value);
 * $ret_3 = $obj->deleteFront();
 * $ret_4 = $obj->deleteLast();
 * $ret_5 = $obj->getFront();
 * $ret_6 = $obj->getRear();
 * $ret_7 = $obj->isEmpty();
 * $ret_8 = $obj->isFull();
 */
// @lc code=end

