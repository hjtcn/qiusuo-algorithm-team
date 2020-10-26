//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的
//功能。(若队列中没有元素，deleteHead 操作返回 -1 )
//
//
//
// 示例 1：
//
// 输入：
//["CQueue","appendTail","deleteHead","deleteHead"]
//[[],[3],[],[]]
//输出：[null,null,3,-1]
//
//
// 示例 2：
//
// 输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]
//
//
// 提示：
//
//
// 1 <= values <= 10000
// 最多会对 appendTail、deleteHead 进行 10000 次调用
//
// Related Topics 栈 设计
// 👍 130 👎 0
package main

import (
	"container/list"
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
type CQueue struct {
	stack1, stack2 *list.List
}


func Constructor() CQueue {
	return CQueue{
		stack1:list.New(),
		stack2:list.New(),
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.stack1.PushBack(value)
}


func (this *CQueue) DeleteHead() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}

	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}

	return -1
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	obj := Constructor();
	fmt.Println(obj)
	obj.AppendTail(3);
	fmt.Println(obj)
	param_2 := obj.DeleteHead();
	fmt.Println(param_2)
}

/**
题解：用两个栈实现题解

使用包"container/list" (双向链表)，现在还不太熟悉，后面需要研究一下

当然本题把双向链表当栈来使用，主要使用container/list封装好的方法
初始化两个空栈，stack1, stack2，
元素插入：直接入栈stack1，
元素删除：使用stack2，删除时先判断stack2是否为空，
	不为空直接删除，
	若为空将stack1元素依次出栈，然后入栈stack2，
	若stack1, 和stack2都为空则返回 -1

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(1)
 */
