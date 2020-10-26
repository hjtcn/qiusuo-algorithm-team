//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//
// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。
//
//
//
//
// 示例:
//
// 输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//
//
//
//
// 提示：
//
//
// pop、top 和 getMin 操作总是在 非空栈 上调用。
//
// Related Topics 栈 设计
// 👍 706 👎 0
package main

import (
	"fmt"
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	stack []int
	min_stack []int
}


/** initialize your data structure here. */
func ConstructorNew() MinStack {
	return MinStack{
		stack:     []int{},
		min_stack: []int{math.MaxInt64},
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	top := this.min_stack[len(this.min_stack) - 1]
	this.min_stack = append(this.min_stack, min(x, top))
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack) - 1]
	this.min_stack = this.min_stack[:len(this.min_stack) - 1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}


func (this *MinStack) GetMin() int {
	return this.min_stack[len(this.min_stack) - 1]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := ConstructorNew();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)

/**
测试代码
 */
func main()  {
	temp_arr := []int{6, 4, 2, -2, 0, 7, 8}
	obj := ConstructorNew();
	for key, val := range temp_arr {
		obj.Push(val)
		if key % 3 == 0 {
			obj.Pop()
		}
	}
	param_3 := obj.Top();
	param_4 := obj.GetMin();
	fmt.Println(param_3, param_4)
}

/**
题解：最小栈

##辅助栈##
本道题，开始想着使用一个栈stack和一个变量min构建一个MinStack的结构体，当时对栈是Push操作时，min = Min(min, x)的变更还很好处理
但是当Pop操作时，min就不好处理

题解提示使用辅助栈，也就是使用两个栈stack, min_stack，
stack栈正常存储出栈和入栈元素，
min_stack每次在stack入栈时，将min(x ,top(min_stack))入栈min_stack，
	所以min_stack栈顶元素代表当前栈stack中的最小值
	每次进行出栈操作时，将stack，min_stack同时将栈顶元素弹出即可

复杂度分析：
	时间复杂度：Push、Pop、GetMin、Top 都为O(1)
	空间复杂度：O(n)
 */
