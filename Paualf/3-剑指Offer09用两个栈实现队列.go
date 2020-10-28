## 用两个栈实现队列：
type CQueue struct {
    stackA []int
    stackB []int
}
func Constructor() CQueue {
    return CQueue{}
}
func (this *CQueue) AppendTail(value int)  {
    this.stackA = append(this.stackA, value)
}
func (this *CQueue) DeleteHead() int {
    if len(this.stackB) == 0 {
        if len(this.stackA) == 0 {
            return -1
        }
        for len(this.stackA) > 0 {
            index := len(this.stackA) - 1
            value := this.stackA[index]
            this.stackB = append(this.stackB, value)
            this.stackA = this.stackA[:index]
        }
    }
    
    index := len(this.stackB) - 1
    value := this.stackB[index]
    this.stackB = this.stackB[:index]
    return value
}

总结：
1. 官方题解是这样定义 CQueue  的，stack 的类型是 *list.List 目前还没有看懂
type CQueue struct {
    stack1, stack2 *list.List
}
func Constructor() CQueue {
    return CQueue{
        stack1: list.New(),
        stack2: list.New(),
    }
}

2. 一开始自己也没有理解题意，在纸上画了一下元素变化的过程，才理解了其中的变化
使用两个栈，一个栈A进行入栈操作，另一个栈B作为出栈操作，出栈的时候，判断B 栈中是否有元素，如果没有则将A栈中的元素，放入到B栈中，放入的过程就是将A栈中的元素逆序操作了，如 A: 1 2 3 4 5,放入到B中则为 5 4 3 2 1， 将5弹出即可。当获取peek，最上面的元素的时候，判断B中是否有元素，如果有则获取B栈中栈顶元素，如果B栈中没有元素，则判断A中是否有元素，如果A中有元素, 将A栈中元素放入到B栈中元素，然后获取栈顶元素返回

3. 画图将题目意思理解清楚以后，代码结构才能更加清晰的理解，如果题目都没有题解清楚，上来就看代码可能就云里雾里了
