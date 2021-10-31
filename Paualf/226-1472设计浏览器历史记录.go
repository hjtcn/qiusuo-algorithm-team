你有一个只支持单个标签页的 浏览器 ，最开始你浏览的网页是 homepage ，你可以访问其他的网站 url ，也可以在浏览历史中后退 steps 步或前进 steps 步。

请你实现 BrowserHistory 类：

BrowserHistory(string homepage) ，用 homepage 初始化浏览器类。
void visit(string url) 从当前页跳转访问 url 对应的页面  。执行此操作会把浏览历史前进的记录全部删除。
string back(int steps) 在浏览历史中后退 steps 步。如果你只能在浏览历史中后退至多 x 步且 steps > x ，那么你只后退 x 步。请返回后退 至多 steps 步以后的 url 。
string forward(int steps) 在浏览历史中前进 steps 步。如果你只能在浏览历史中前进至多 x 步且 steps > x ，那么你只前进 x 步。请返回前进 至多 steps步以后的 url 。
 

示例：
输入：
["BrowserHistory","visit","visit","visit","back","back","forward","visit","forward","back","back"]
[["leetcode.com"],["google.com"],["facebook.com"],["youtube.com"],[1],[1],[1],["linkedin.com"],[2],[2],[7]]
输出：
[null,null,null,null,"facebook.com","google.com","facebook.com",null,"linkedin.com","google.com","leetcode.com"]

解释：
BrowserHistory browserHistory = new BrowserHistory("leetcode.com");
browserHistory.visit("google.com");       // 你原本在浏览 "leetcode.com" 。访问 "google.com"
browserHistory.visit("facebook.com");     // 你原本在浏览 "google.com" 。访问 "facebook.com"
browserHistory.visit("youtube.com");      // 你原本在浏览 "facebook.com" 。访问 "youtube.com"
browserHistory.back(1);                   // 你原本在浏览 "youtube.com" ，后退到 "facebook.com" 并返回 "facebook.com"
browserHistory.back(1);                   // 你原本在浏览 "facebook.com" ，后退到 "google.com" 并返回 "google.com"
browserHistory.forward(1);                // 你原本在浏览 "google.com" ，前进到 "facebook.com" 并返回 "facebook.com"
browserHistory.visit("linkedin.com");     // 你原本在浏览 "facebook.com" 。 访问 "linkedin.com"
browserHistory.forward(2);                // 你原本在浏览 "linkedin.com" ，你无法前进任何步数。
browserHistory.back(2);                   // 你原本在浏览 "linkedin.com" ，后退两步依次先到 "facebook.com" ，然后到 "google.com" ，并返回 "google.com"
browserHistory.back(7);                   // 你原本在浏览 "google.com"， 你只能后退一步到 "leetcode.com" ，并返回 "leetcode.com"

提示：
1 <= homepage.length <= 20
1 <= url.length <= 20
1 <= steps <= 100
homepage 和 url 都只包含 '.' 或者小写英文字母。
最多调用 5000 次 visit， back 和 forward 函数。

1. Clarification:

Q：如何存储浏览历史？
A:  使用 stack 进行存储

Q: 如何标记当前位置？
A：使用 index 标记当前所在的位置信息

Q: visit 的时候如何清空元素信息？
A: 判断元素数量？ len(stack) : 当前栈中元素数量个数 与 当前位置进行比较

如果 栈中元素数量 > 当前位置 + 1 {
将栈后面元素清空
stack = stack[0:current + 1]
}

是不是还可以判断 stack[current + 1] 元素是否存在呢？
if len(stack[current+1]) > 0 {
则进行清空
}
如果不大于0，则不用进行清空

type BrowserHistory struct {
    stack []string
    curr int
}


func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{
        stack: make([]string,0),
        curr: 0,
    }
}


func (this *BrowserHistory) Visit(url string)  {
    // 如果当前位置下表后还有元素，则需要把元素干掉
    if len(this.stack) - 1 > this.curr {
        this.stack = this.stack[0:this.curr + 1]
    }
    this.stack = append(this.stack,url)
    this.curr = len(this.stack) - 1
    fmt.Println("visit", this.stack)
}


func (this *BrowserHistory) Back(steps int) string {
    // 如果后退步数 大于 剩余元素数量
    var ret string
    fmt.Println(this.stack)
    if steps > len(this.stack) - 1 {
        ret = this.stack[0]
        this.curr = 0
        return ret
    }
    stepsThenCurr := this.curr - steps
    ret = this.stack[stepsThenCurr]
    this.curr = stepsThenCurr
    return ret
}


func (this *BrowserHistory) Forward(steps int) string {
    var ret string
    if this.curr + steps > len(this.stack) - 1 {
        ret = this.stack[len(this.stack) - 1]
        this.curr = len(this.stack) - 1
        return ret
    }

    stepsForwardCurr := this.curr + steps
    ret = this.stack[stepsForwardCurr]
    this.curr = stepsForwardCurr
    return ret
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
一开始初始化的时候初始化没有将值放进去改了一下
type BrowserHistory struct {
    stack []string
    curr int
}


func Constructor(homepage string) BrowserHistory {
    a := BrowserHistory{
        stack: make([]string,0),
        curr: 0,
    }
    a.stack = append(a.stack,homepage)
    return a
}


func (this *BrowserHistory) Visit(url string)  {
    // 如果当前位置下表后还有元素，则需要把元素干掉
    if len(this.stack) - 1 > this.curr {
        this.stack = this.stack[0:this.curr + 1]
    }
    this.stack = append(this.stack,url)
    this.curr = len(this.stack) - 1
    // fmt.Println("visit", this.stack)
}


func (this *BrowserHistory) Back(steps int) string {
    // 如果后退步数 大于 剩余元素数量
    var ret string
    //fmt.Println(this.stack)
    // 这里判断条件注意下
    if this.curr - steps < 0 {
        ret = this.stack[0]
        this.curr = 0
        return ret
    }
    stepsThenCurr := this.curr - steps
    ret = this.stack[stepsThenCurr]
    this.curr = stepsThenCurr
    return ret
}


func (this *BrowserHistory) Forward(steps int) string {
    var ret string
    if this.curr + steps > len(this.stack) - 1 {
        ret = this.stack[len(this.stack) - 1]
        this.curr = len(this.stack) - 1
        return ret
    }

    stepsForwardCurr := this.curr + steps
    ret = this.stack[stepsForwardCurr]
    this.curr = stepsForwardCurr
    return ret
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
    if this.curr - steps < 0 { }
这个判断条件感觉还是蛮有意思的，因为下面用到了所以上面把不合法的给过滤掉了，思想 => fail fast


2. 看题解：
题解的back 还是写的比我简洁：
if steps > this.index{
		this.index = 0
		return this.cur[0]
}
this.index -= steps
return this.cur[this.index]
看到这个 forward 也是很简洁了
func (this *BrowserHistory) Forward(steps int) string {
	this.temp=min(this.temp+steps,len(this.hist)-1)
	return this.hist[this.temp]
}

3. 复杂度分析：
时间复杂度：O(1）
空间复杂度：O(n) 使用栈进行存储

4. 总结：
4.1: 早上把大流程写完了，晚上回来调试通过了，还是蛮有成就感的
4.2: 越来越觉得写代码不是在写代码是在梳理思路，然后去实践思路
4.3: 多看题解，多对比，才能有进步
