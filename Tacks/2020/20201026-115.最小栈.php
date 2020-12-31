<?php

/*
 * @lc app=leetcode.cn id=155 lang=php
 *
 * [155] 最小栈
 */

// @lc code=start
class MinStack {

    private $stack;
    private $minStack;
    /**
     * initialize your data structure here.
     */
    function __construct() {
        // SplStack是php的标准库
        $this->stack = new SplStack();
        $this->minStack = new SplStack();
    }
  
    /**
     * @param Integer $x
     * @return NULL
     */
    // 入栈
    function push($x) {
        $this->stack->push($x); // 入栈
        // 判断最小栈
        if($this->minStack->count()) {
            // 比较栈顶元素跟 当前插入的值进行比较 谁更小谁就插入
            $x = min($this->minStack->top(), $x);
        }
        $this->minStack->push($x);
    }
  
    /**
     * @return NULL
     */
    // 出栈
    function pop() {
        if($this->stack->isEmpty()) {
            return null;
        }
        $this->stack->pop(); // 出栈
        if($this->minStack->count()) {
            // 保持同时出栈
            $this->minStack->pop();
        }
    }
  
    /**
     * @return Integer
     */
    function top() {
        // 栈顶元素
        return $this->stack->top();
    }
  
    /**
     * @return Integer
     */
    function getMin() {
        // 如果最小栈为空 直接返回
        if($this->minStack->isEmpty()) {
            return null;
        }
        // 最小栈 栈顶维护最小值
        return $this->minStack->top();
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * $obj = MinStack();
 * $obj->push($x);
 * $obj->pop();
 * $ret_3 = $obj->top();
 * $ret_4 = $obj->getMin();
 */
// @lc code=end



/*

【标准库SplStack】 辅助栈
额外维护一个最小栈用来以O(1)的时间复杂度返回

使用一个辅助栈，与元素栈同步插入与删除，用于存储与每个元素对应的最小值

push 取当前辅助栈的栈顶存储的最小值，与当前元素比较得出最小值，将这个最小值插入辅助栈中以及栈中
pop  当一个元素要出栈时，我们把辅助栈的栈顶元素也一并弹出
getMin 最小值就存储在辅助栈的栈顶元素中。


时间复杂度O(1)
空间复杂度O(N)
执行用时：32 ms, 在所有 PHP 提交中击败了86.46%的用户
内存消耗：22.6 MB, 在所有 PHP 提交中击败了5.56%的用户
*/