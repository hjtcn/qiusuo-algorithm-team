// 剑指 Offer 09. 用两个栈实现队列
// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

 

// 示例 1：

// 输入：
// ["CQueue","appendTail","deleteHead","deleteHead"]
// [[],[3],[],[]]
// 输出：[null,null,3,-1]
// 示例 2：

// 输入：
// ["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
// [[],[],[5],[2],[],[]]
// 输出：[null,-1,null,null,5,2]
// 提示：

// 1 <= values <= 10000
// 最多会对 appendTail、deleteHead 进行 10000 次调用




var CQueue = function() {
    this.stack1=[]
    this.stack2=[]
};

/** 
 * @param {number} value
 * @return {void}
 */
CQueue.prototype.appendTail = function(value) {
    //栈1控制元素入队
    this.stack1.push(value)
};

/**
 * @return {number}
 */
CQueue.prototype.deleteHead = function() {
    //栈二控制元素出队
    if(this.stack2.length){
        //stack2不为空，直接弹出栈顶元素
       return this.stack2.pop()
    }
    //此时stack2为空，将stack1中的栈顶元素弹出的同时push到stack2中
   while(this.stack1.length){
       this.stack2.push(this.stack1.pop())
   }
   //返回stack2弹出的栈顶元素,不存在即为-1
   return this.stack2.pop()||-1
};

/**
 * Your CQueue object will be instantiated and called as such:
 * var obj = new CQueue()
 * obj.appendTail(value)
 * var param_2 = obj.deleteHead()
 */


 /** 题解
  //没过10分钟，我就去看题解了，为什么呢？题意我是懂了，但是我没懂输入输出，以及js定义的对象和原型函数(js要补补了)
  维护双栈方法。原理也很好理解，两个栈意思是只能用push,pop
  就是第二个栈来push第一个栈pop的数据，这下栈顶元素就跑到栈底了。

  解决方法：
  一个stack1用来添加，一个stack2用来删除。
  删除的时候判断stack2是否为空，不为空，直接返回stack2的栈顶元素。
  为空，则开始不断往stack2里面传stack1弹出的值了。
  最后弹出stack2的栈顶元素并返回，不存在栈顶元素，就返回-1了
复杂度分析：
 时间复杂度：O(1)
 每个元素最多只会有一次插入和弹出的一次
 空间复杂度：O(n)
 定义两个栈
*/