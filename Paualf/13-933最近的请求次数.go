/*
写一个 RecentCounter 类来计算特定时间范围内最近的请求。
请你实现 RecentCounter 类：
RecentCounter() 初始化计数器，请求数为 0 。
int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去 3000 毫秒内发生的所有请求数（包括新请求）。确切地说，返回在 [t-3000, t] 内发生的请求数。
保证每次对 ping 的调用都使用比之前更大的 t 值。
示例 1：
输入：
["RecentCounter", "ping", "ping", "ping", "ping"]
[[], [1], [100], [3001], [3002]]
输出：
[null, 1, 2, 3, 3]
解释：
RecentCounter recentCounter = new RecentCounter();
recentCounter.ping(1);     // requests = [1]，范围是 [-2999,1]，返回 1
recentCounter.ping(100);   // requests = [<u>1</u>, <u>100</u>]，范围是 [-2900,100]，返回 2
recentCounter.ping(3001);  // requests = [<u>1</u>, <u>100</u>, <u>3001</u>]，范围是 [1,3001]，返回 3
recentCounter.ping(3002);  // requests = [1, <u>100</u>, <u>3001</u>, <u>3002</u>]，范围是 [2,3002]，返回 3
分析：
如何设计每次请求的时间t,然后返回在[t-3000,t] 内发生的请求次数
使用数组存储数据
每次请求t的，时候，循环数组判断在 [t-3000,t] 内的元素个数：
*/

type RecentCounter struct {
    data[] int
    num int
}
func Constructor() RecentCounter {
    return RecentCounter{
        data:make([]int, 0),
        num:0,
    }
}
func (this *RecentCounter) Ping(t int) int {
    this.data = append(this.data,t)
    ret := 0
    // 比较开始的时候只需要从上一次记录的 this.num 个数开始查找
    start := len(this.data) - 1 - this.num
    randStart := t - 3000
    for i := start;i < len(this.data);i++ {
        if this.data[i] >= randStart && this.data[i] <= t {
            ret++
        }
    }
    this.num = ret
    return ret
}
/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

如何进行优化呢？如果是一个常驻进程的程序，你开启一个切片或者数组来存储元素，如果元素很多的话，不进行释放的话会内存溢出的，可不可以节省空间无须存储前面无用的数据呢？
看到了题解里面这样写还是比较优秀的，我遍历数组的时候就直到它是否在有效范围内也就是[t-3000,t] 中，如果不在有效范围内的话，则将该元素移除数组就可以了
里面的break 用的也是很巧妙的，自己写的时候要记得写break 哈，找到第一个不在范围内的数，然后将它之前的截掉就可以了

type RecentCounter struct {
    queue []int
}
func Constructor() RecentCounter {
    return RecentCounter{
        queue: make([]int, 0),
    }
}
func (this *RecentCounter) Ping(t int) int {
    this.queue = append(this.queue, t)
    l,ans := len(this.queue),1
    for l -= 2;l >= 0;l-- {
        if t - this.queue[l] <= 3000 {
            ans++
        }else {
            this.queue = this.queue[l+1:]
            break
        }
    }
    return ans
}
/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

复杂度分析：
时间复杂度：O(n) :每次进来遍历数组进行统计个数
空间复杂度：O(n): 使用数组存储队列
代码还可以更简洁，也是很厉害了

type RecentCounter struct {
    queue []int
}
func Constructor() RecentCounter {
    return RecentCounter{}
}
func (this *RecentCounter) Ping(t int) int {
    this.queue = append(this.queue, t)
    for this.queue[0]+3000 < t {
        this.queue = this.queue[1:]
    }
    return len(this.queue)
}

看到别人的代码还可以用二分查找来优化时间复杂度也是醉醉的，想到这样还是挺厉害的，因为每次 ping 调用都使用比之前更大的t值, 找到t - 3000的下标，然后 len(this.a) - index, 这里我也发现数组下标还是蛮有意思的
[1,2,3,4,5]
下标
[0,1,2,3,4]
5个元素，返回下标1，5-1=4，下标1不仅仅可以代表第2个元素的下标，还可以代表它前面有多少个元素

type RecentCounter struct {
    a []int
}
func Constructor() RecentCounter {
    return RecentCounter{[]int{}}
}
func (this *RecentCounter) Ping(t int) int {
    this.a = append(this.a, t)
    
    idx := binarySearch(this.a, t - 3000)
    
    return len(this.a) - idx
}
func binarySearch(arr []int, target int) int {
    lo, hi := 0, len(arr)-1
    
    for lo <= hi {
        mid := (lo + hi) / 2
        if arr[mid] > target {
            hi = mid - 1
        } else if arr[mid] < target {
            lo = mid + 1
        } else {
            return mid
        }
    }
    
    return lo
}

复杂度分析：
   时间复杂度：O(logN)：使用二分查找减少时间复杂度
   空间复杂度：O(n) : 使用数组存储队列

总结：

1. 题目不是很难，但是自己写的代码目前并不是很好

2. 同样是一种语言，一道题目，可以把逻辑写的这么简洁是很厉害的，还有有使用到二分查找来降低查询时间复杂度还是挺好的
