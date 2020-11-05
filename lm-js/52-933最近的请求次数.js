// 933. 最近的请求次数
// 写一个 RecentCounter 类来计算特定时间范围内最近的请求。

// 请你实现 RecentCounter 类：

// RecentCounter() 初始化计数器，请求数为 0 。
// int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去 3000 毫秒内发生的所有请求数（包括新请求）。确切地说，返回在 [t-3000, t] 内发生的请求数。
// 保证每次对 ping 的调用都使用比之前更大的 t 值。

 

// 示例 1：

// 输入：
// ["RecentCounter", "ping", "ping", "ping", "ping"]
// [[], [1], [100], [3001], [3002]]
// 输出：
// [null, 1, 2, 3, 3]

// 解释：
// RecentCounter recentCounter = new RecentCounter();
// recentCounter.ping(1);     // requests = [1]，范围是 [-2999,1]，返回 1
// recentCounter.ping(100);   // requests = [<u>1</u>, <u>100</u>]，范围是 [-2900,100]，返回 2
// recentCounter.ping(3001);  // requests = [<u>1</u>, <u>100</u>, <u>3001</u>]，范围是 [1,3001]，返回 3
// recentCounter.ping(3002);  // requests = [1, <u>100</u>, <u>3001</u>, <u>3002</u>]，范围是 [2,3002]，返回 3
 

// 提示：

// 1 <= t <= 104
// 保证每次对 ping 的调用都使用比之前更大的 t 值
// 至多调用 ping 方法 104 次

/*
 * @lc app=leetcode.cn id=933 lang=javascript
 *
 * [933] 最近的请求次数
 */

// @lc code=start

var RecentCounter = function() {
    this.queue=[]
};

/** 
 * @param {number} t
 * @return {number}
 */
RecentCounter.prototype.ping = function(t) {
    while(this.queue[0]<t-3000){
        this.queue.shift()
    }
    this.queue.push(t)
    return this.queue&&this.queue.length
};

/**
 * Your RecentCounter object will be instantiated and called as such:
 * var obj = new RecentCounter()
 * var param_1 = obj.ping(t)
 */
// @lc code=end

    /**
    暴力。如果新增的数减去3000大于队列头，则需要从头部开始shift队列。
    将当前元素入队。返回队列长度即可

    复杂度分析：
    时间复杂度：O(N)
    得知当前元素-3000的差，则开始循环判断出队操作。

    空间复杂度：O(N)
    利用数组模拟队列
   */