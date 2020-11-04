<?php

/*
 * @lc app=leetcode.cn id=933 lang=php
 *
 * [933] 最近的请求次数
 *
 * https://leetcode-cn.com/problems/number-of-recent-calls/description/
 *
 * algorithms
 * Easy (71.97%)
 * Likes:    68
 * Dislikes: 0
 * Total Accepted:    19.3K
 * Total Submissions: 26.9K
 * Testcase Example:  '["RecentCounter","ping","ping","ping","ping"]\n[[],[1],[100],[3001],[3002]]'
 *
 * 写一个 RecentCounter 类来计算特定时间范围内最近的请求。
 * 
 * 请你实现 RecentCounter 类：
 * 
 * 
 * RecentCounter() 初始化计数器，请求数为 0 。
 * int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去 3000
 * 毫秒内发生的所有请求数（包括新请求）。确切地说，返回在 [t-3000, t] 内发生的请求数。
 * 
 * 
 * 保证每次对 ping 的调用都使用比之前更大的 t 值。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：
 * ["RecentCounter", "ping", "ping", "ping", "ping"]
 * [[], [1], [100], [3001], [3002]]
 * 输出：
 * [null, 1, 2, 3, 3]
 * 
 * 解释：
 * RecentCounter recentCounter = new RecentCounter();
 * recentCounter.ping(1);     // requests = [1]，范围是 [-2999,1]，返回 1
 * recentCounter.ping(100);   // requests = [<u>1</u>, <u>100</u>]，范围是
 * [-2900,100]，返回 2
 * recentCounter.ping(3001);  // requests = [<u>1</u>, <u>100</u>,
 * <u>3001</u>]，范围是 [1,3001]，返回 3
 * recentCounter.ping(3002);  // requests = [1, <u>100</u>, <u>3001</u>,
 * <u>3002</u>]，范围是 [2,3002]，返回 3
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= t <= 10^4
 * 保证每次对 ping 的调用都使用比之前更大的 t 值
 * 至多调用 ping 方法 10^4 次
 * 
 * 
 */

// @lc code=start
class RecentCounter {

    protected $queue; // 队列

    function __construct() {
        $this->queue = []; // 初始化
    }

    /**
     * @param Integer $t
     * @return Integer
     */
    function ping($t) {
        // 每次进入队列
        array_push($this->queue, $t);
        // 如果超出3000的时间队列就弹出。
        while( ($t - current($this->queue) > 3000) ) {
            array_shift($this->queue); // 弹出
        }
        return count($this->queue);
    }
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * $obj = RecentCounter();
 * $ret_1 = $obj->ping($t);
 */
// @lc code=end


/*
题目一开始没看懂，然后去看了题解，然后发现其实就是一个阅读理解，没什么算法，也就是用到队列。我这里使用数组模拟队列。

准确的说，就是返回在3000内的请求数 [t-3000, t]


每次，程序都会调用一此ping,传入一个t是时间值，标识进行一次ping请求，多次请求下来，才是一个数组

示例
 * ["RecentCounter", "ping", "ping", "ping", "ping"]
 * [[], [1], [100], [3001], [3002]]
 * 输出：
 * [null, 1, 2, 3, 3]
 * 
[“RecentCouter”] 与 [null] 不需要考虑
第一个 ping声 出现的时间为 1毫秒，
第二个 ping声 出现的时间为 100毫秒，
第三个 ping声 出现的时间为 300毫秒，
第四个 ping声 出现的时间为 3002毫秒。

那么就要计算的是在当前ping声出现的时间及其前3000毫秒出现了多少次 ping 声

如果把当前ping声出现的时间记做 t ，那么就要找出(t-3000)毫秒至t毫秒之间的 ping声有多少次。

也就是当
t = 1     (0~1)    此时符合时间范围的ping 就是1 一个
t = 100   (0~100)  此时符合时间范围的ping 就是1，100 两个
t = 3001  (1~3001) 此时符合时间范围的ping 就是1，100，3001，三个
t = 3002  (2~3002) 此时符合时间范围的ping 就是2，100，3002，三个

就有点类似队列，一个长度为3000的队列，

如果超出这个长度，那么前面的值就需要出队。

每次求的值，也就是队列中符合条件的值的个数。 
*/