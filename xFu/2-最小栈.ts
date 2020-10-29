// 155. 最小栈
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。
 

class MinStack {
    constructor() {

    }

    stackList:number[] = [];

    // 0,1,2
    minNumList:number[] = [];

    push(x: number): void {
        this.stackList.unshift(x);
        // 将小于等于最小值的值从前插入数组中
        if(this.minNumList[0] === undefined || this.minNumList[0] >= x){
            this.minNumList.unshift(x);
        }
    }

    pop(): void {
        let popValue = this.stackList.shift();
        // 如果原栈中移除的是当前的最小值，就同步移除这个最小值
        if(popValue === this.minNumList[0]){
            this.minNumList.shift();
        }
    }

    top(): number {
        return this.stackList[0];
    }

    getMin(): number {
        return this.minNumList[0];
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(x)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */