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
// 至多调用 ping 方法 10^4 次


// ============================================================
// ===                  学习的优秀代码思想                   ===
// ===      状态：通过,执行用时: 252 ms,内存消耗: 47.9 MB     ===
// ============================================================

class RecentCounter {
    constructor() {

    }

    queue:number[] = []

    ping(t: number): number {
        // 先入队当前的值，再判断若第一个值比最小值还要小就直接出队（剔除）这个值。（题目保证了数的递增性）
        // 最后保留下来的一定是符合条件的
        this.enQueue(t);

        while(this.queue[0] < t - 3000){
            this.deQueue();
        }
        
        return this.queue.length;
    }

    deQueue():number {
        return this.queue.shift() || 0;
    }

    enQueue(num: number):number {
        return this.queue.push(num);
    }
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * var obj = new RecentCounter()
 * var param_1 = obj.ping(t)
 */



// ============================================================
// ===                     自己想的代码                      ===
// ===      状态：通过,执行用时: 700 ms,内存消耗: 48.6 MB     ===
// ============================================================

class RecentCounter_1 {
    constructor() {

    }

    queue:number[] = []

    ping(t: number): number {
        const maxRange = t;
        const minRange = t - 3000;

        let sum = 0;
        const queueLen = this.queue.length;
        let index = 0;
        
        while(index < queueLen){
            // 出队
            const queueNum = this.queue[index];
            // 从头开始判断，若不符合范围就+1，符合了就直接结束，因为后续的数都一定是符合的了。
            if(queueNum >= minRange && queueNum <= maxRange){
                break;
            }else{
                sum++;
            }
            index++;
        }
        this.enQueue(t);

        return this.queue.length - sum;
    }

    queueTop():number {
        return this.queue[0] || 0;
    }

    enQueue(num: number):number {
        return this.queue.push(num);
    }
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * var obj = new RecentCounter()
 * var param_1 = obj.ping(t)
 */


// ============================================================
// ===                      纯队列思想                       ===
// ===                        超时                           ===
// ============================================================


// 思想：从队列中一个一个的出队，判断这个出队的数字是不是在范围内，判断完毕将数字重新入队到队列中
//       最后入队用于判断的值
class RecentCounter_OLD {
    constructor() {

    }

    queue:number[] = []

    ping(t: number): number {
        const maxRange = t;
        const minRange = t - 3000;

        let sum = 0;
        let queueLen = this.queue.length;

        while(queueLen--){
            // 出队
            const queueNum = this.deQueue();
            // 超出范围就直接结束循环则会破坏原始顺序
            if(queueNum >= minRange && queueNum <= maxRange){
                // 计算个数
                sum++;
            }
            // 入队到队列中
            this.enQueue(queueNum);
        }
        // 本身一定存在于范围内
        sum++;
        this.enQueue(t);

        return sum;
    }

    deQueue():number {
        return this.queue.pop() || 0;
    }

    enQueue(num: number):number {
        return this.queue.unshift(num);
    }
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * var obj = new RecentCounter()
 * var param_1 = obj.ping(t)
 */