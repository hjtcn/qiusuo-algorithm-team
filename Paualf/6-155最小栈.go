题目：
/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。
*/
思路：
使用辅助栈minStack存储和查找最小的元素，当一个元素append进入stack的时候,将它和 minStack 最顶端的元素对比，取两者之间的较小值，append进入到minStack中

type MinStack struct {
    stack []int
    minStack []int
}
/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stack: []int{},
        minStack: []int{math.MaxInt64},
    }
}
func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack,x)
    top := this.minStack[len(this.minStack) - 1]
    this.minStack = append(this.minStack,min(x,top))
}
func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack) - 1]
    this.minStack = this.minStack[:len(this.minStack) - 1]
}
func (this *MinStack) Top() int {
    return this.stack[len(this.stack) - 1]
}
// 如何检索栈中的最小
func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack) - 1]
}
func min(x,y int) int{
    if x < y {
        return x
    }
    return y
}
/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */



// 这个地方的返回值和定义自己没有写出来，应该是没有理解为什么是这样写,后来看了下资料 stack:[]int{},这种是结构体初始化的一种方式，说基本的语言语法没有掌握
func Constructor() MinStack {
    return MinStack{
        stack: []int{},
        minStack: []int{math.MaxInt64},
    }
}
从老黑哪里学到了Go标准库里面实现了 container/list 的标准库，list包实现了一个双向链表（doubly linked list)
type MinStack struct {
    stack,minStack *list.List
}
/** initialize your data structure here. */
func Constructor() MinStack {
    a := list.New()
    b := list.New()
    b.PushBack(math.MaxInt64)
    return MinStack{
        stack:a,
        minStack:b,
    }     
}
func (this *MinStack) Push(x int)  {
    this.stack.PushBack(x)
    currentMin := this.minStack.Back().Value.(int)
 
    minNum := min(x, currentMin)
    this.minStack.PushBack(minNum)
}
func (this *MinStack) Pop()  {
    this.stack.Remove(this.stack.Back())
    this.minStack.Remove(this.minStack.Back())
}
func (this *MinStack) Top() int {
    return this.stack.Back().Value.(int)
}
// 如何检索栈中的最小
func (this *MinStack) GetMin() int {
    return this.minStack.Back().Value.(int)
}
func min(x,y int)int {
    if x < y {
        return x
    }
    return y
}
/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

一开始的时候一直出现这个报错：
Line 34: Char 29: cannot use this.stack.Back().Value (type interface {}) as type int in return argument: need type assertion (solution.go)
不能使用接口类型转换为int,自己在编辑器里面跑是正常的，在leetcode 中运行就会出现这个问题，猜测应该是标准库使用的版本可能不一样，后来在题解里面看到大神可以使用 .(int)，然后就解决了这个问题

总结：
1. 未掌握结构体初始化语法（学习的时候觉得这些很简单，到实际用的时候就傻眼了。。。）

2. 第一次接触和了解Go标准库，一开始心里面是有点害怕和觉得难的，后来动手把demo跑起来改了改，发现还是可以理解的，没有想象中那么难，所以动手做，不要一直停留在想的层面上

3. 空间换时间的思想，没有完美的结构和设计，只有选择合适的数据结构解决当下的问题才是好的设计和代码

时间复杂度和空间复杂度分析：
时间复杂度：append,pop,top,getMin 时间复杂度都为O(1)
空间复杂度为： 因为需要额外维护一个栈，所以空间复杂度为O(n)
