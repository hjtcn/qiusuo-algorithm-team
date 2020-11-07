// 622. 设计循环队列
// 设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。

// 循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。

// 你的实现应该支持如下操作：

// MyCircularQueue(k): 构造器，设置队列长度为 k 。
// Front: 从队首获取元素。如果队列为空，返回 -1 。
// Rear: 获取队尾元素。如果队列为空，返回 -1 。
// enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
// deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
// isEmpty(): 检查循环队列是否为空。
// isFull(): 检查循环队列是否已满。
 

// 示例：

// MyCircularQueue circularQueue = new MyCircularQueue(3); // 设置长度为 3
// circularQueue.enQueue(1);  // 返回 true
// circularQueue.enQueue(2);  // 返回 true
// circularQueue.enQueue(3);  // 返回 true
// circularQueue.enQueue(4);  // 返回 false，队列已满
// circularQueue.Rear();  // 返回 3
// circularQueue.isFull();  // 返回 true
// circularQueue.deQueue();  // 返回 true
// circularQueue.enQueue(4);  // 返回 true
// circularQueue.Rear();  // 返回 4
 

// 提示：

// 所有的值都在 0 至 1000 的范围内；
// 操作数将在 1 至 1000 的范围内；
// 请不要使用内置的队列库。

// ============================================================
// ===                     自己想的代码                      ===
// ===      状态：通过,执行用时: 124 ms,内存消耗: 46.3 MB     ===
// ============================================================

class MyCircularQueue {
    constructor(k: number) {
        this.queue = new Array(k+1);
    }

    front = 0
    rear = 0
    queue:number[] = []

    enQueue(value: number): boolean {
        if(this.isFull()){
            return false;
        }
        this.queue[this.rear] = value;
        this.rear = this.getNext(this.rear);
        return true;
    }

    deQueue(): boolean {
        if(this.isEmpty()){
            return false;
        }
        this.front = this.getNext(this.front);
        return true;
    }

    Front(): number {
        return this.isEmpty() ? -1 : this.queue[this.front];
    }

    Rear(): number {
        return this.isEmpty() ? -1 : this.queue[this.getPrev(this.rear)];
    }

    isEmpty(): boolean {
        return this.front ===  this.rear;
    }

    isFull(): boolean {
        return this.front === this.getNext(this.rear);
    }

    /**
     * 寻找当前下标的下一个下标位置
     * @param current 当前下标位置
     */
    getNext(current: number): number {
        let nextIndex = current + 1;
        // 若判断的是最后一个元素，则他的下一个元素应该是第一个元素
        if(nextIndex >= this.queue.length){
            nextIndex = 0;
        }
        return nextIndex;
    }

    /**
     * 寻找当前下标的上一个下标位置
     * @param current 当前下标位置
     */
    getPrev(current: number): number {
        let prevIndex = current - 1;
        // 若判断的是第一个元素，则他的上一个元素应该是列表的最后一个元素
        if(prevIndex < 0){
            prevIndex = this.queue.length - 1;
        }
        return prevIndex;
    }
}

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