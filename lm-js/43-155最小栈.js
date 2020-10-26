// 155. 最小栈
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。
 

// 示例:

// 输入：
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]

// 输出：
// [null,null,null,null,-3,null,0,-2]

// 解释：
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.
 

// 提示：

// pop、top 和 getMin 操作总是在 非空栈 上调用。


/*
 * @lc app=leetcode.cn id=155 lang=javascript
 *
 * [155] 最小栈
 */

// @lc code=start
/**
 * initialize your data structure here.
 */
var MinStack = function() {
    this.stack1=[]
    this.stack2=[]
};

/** 
 * @param {number} x
 * @return {void}
 */
MinStack.prototype.push = function(x) {
    this.stack1.push(x)
    //如果stack2为空或x小于等于stack2的栈顶元素时，push元素x
    if((this.stack2.length&&x<=this.stack2[this.stack2.length-1])||!this.stack2.length){
        this.stack2.push(x)
    }
    
};

/**
 * @return {void}
 */
MinStack.prototype.pop = function() {
    let k=this.stack1.pop()
    //如果pop掉的元素等于stack2的栈顶元素,stack2也pop
    if(k==this.stack2[this.stack2.length-1]){
        this.stack2.pop()
    }
};

/**
 * @return {number}
 */
MinStack.prototype.top = function() {
    return this.stack1[this.stack1.length-1]
};

/**
 * @return {number}
 */
MinStack.prototype.getMin = function() {
    return this.stack2[this.stack2.length-1]
};

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(x)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */
// @lc code=end

/** 题解
 * 第二次看到这样的输入输出结构，就没有第一次那么慌张了。
 * 这个题做的还是满顺利的，唯一的难度点应该就是获取最小值。
    除了getMin其它的方法，直接定义一个数组来模拟栈的push，pop，top(即数组的尾部元素)操作即可。
    考虑到getMin方法，一个栈是解决不了的，需要再定义一个栈去存取连续递减的最小值。
    为什么连续递减栈就可以了？？因为直到栈底元素(即数组的第一个元素)也被pop掉的过程中，getMin()也不可能是大于栈底元素的那些值
    因此反向思考，递减栈先存储第一个元素作为栈顶元素，然后遇到小于等于(？？为什么要等于呢？？)栈顶元素的就push入递减栈stack2中。最后每一次获取stack2的栈顶元素，即为getMin()
    ！！！考虑等于栈顶元素这种情况，是为了防止pop掉stack2栈顶元素(即最小值)的时候，其实还有同样的这个值，但是由于一开始没被push到stack2，这样getMin就会出错。


复杂度分析：
  时间复杂度：O(1)
  每一次操作都有一次入栈出栈获取的处理过程

  空间复杂度：O(N)
  变量声明数组，模拟栈的操作。
*/



/**
 * initialize your data structure here.
 */
var MinStack = function() {
    this.stack1=[]//存储的是原元素减去最小值的差
    this.minVal=Infinity//记录最小元素
};

/** 
 * @param {number} x
 * @return {void}
 */
MinStack.prototype.push = function(x) {
   this.stack1.push(x-this.minVal)//将x与最小值的差值入栈
   this.minVal=x<this.minVal?x:this.minVal//更新当前最小值  
};

/**
 * @return {void}
 */
MinStack.prototype.pop = function() {
    let k=this.stack1.pop()
    let minVal=this.minVal
    if(k<0){
        //栈顶元素(差)小于0,则栈对应的真实元素为当前最小值并返回，且更新当前最小值为原最小值减去k则为上一个最小值
        this.minVal=minVal-k
        return minVal
    }
    return k+minVal
};

/**
 * @return {number}
 */
MinStack.prototype.top = function() {
    let k=this.stack1[this.stack1.length-1]
    let minVal=this.minVal
    if(k<0){
        //这个地方写的时候还犯了个错，只有push和pop会影响最小值
        //对比pop方法，一样的思路，只是不必更新当前元素。
        return minVal
    }
    return k+minVal
};

/**
 * @return {number}
 */
MinStack.prototype.getMin = function() { 
    return this.minVal
};


/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(x)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */


 /** 题解
    第二个方法是看题解get的，其实看了几遍还是比较模糊。
    这种不记元素本身，记录差，且pop的时候还能通过差值再找回来的过程，自己硬想真的太难了。
    后来把整个题目代码看了一遍，正反思路都走了走流程，才走通的。
    思路的关键点：
    push方法记录的是x(新元素)与minVal(最小值)的差，且更新最小值(x<minVal才更新，this.minVal=x)
    pop方法则是倒推了，首先pop掉栈顶元素，且记录下来为k。
    如果k<0，则当时push的时候，最小值肯定更新了，倒推(x=this.minVal)，x也是想要返回的目标元素，不要忘了往回找更新最小值(this.minVal=this.minVal-k);
       k>=0,则当时push的时候，最小值未更新，倒推(x=this.minVal+k(差值))，x就是返回的目标元素。
    top方法和pop方法倒推获取方式基本一致，只是不会影响栈和最小值而已。。。。。
    getMin方法，直接返回this.minVal即可
    

复杂度分析：
  时间复杂度：O(1)
  每一次操作都有一次入栈出栈获取的处理过程

  空间复杂度：O(1)
  这里有些疑惑，题解上说空间复杂度已经为O(1)了，但是我考虑到依然定义栈去存储差值了，去处理一些push和pop操作。
  不过它相对于第一种方案是肯定优化了空间复杂度，没有再借助另一个栈去存取最小值的过程。
*/