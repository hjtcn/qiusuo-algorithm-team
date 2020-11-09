/*
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
示例:
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7] 
解释: 
  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
*/

第一种可以想到的是暴力法：
遍历i，然后j 从 [i + 1, i + k - 1] 中找到最大值

func maxSlidingWindow(nums []int, k int) []int {
    length := len(nums)
    ret := make([]int, 0)
    if length == 0 {
        return ret
    }
    num := length - k + 1
    for i := 0;i < num;i++ {
        max := nums[i] 
        for j := i + 1;j < i + k;j++ {
            if nums[j] > max {
                max = nums[j]
            }
        }
        ret = append(ret, max)
    }
    return ret
}

复杂度分析：
时间复杂度：O(n * n): 两层for循环
空间复杂度：O(1)：单独变量 O(n)：用于存储结果数组

然后看了 labuladong 的分析：https://leetcode-cn.com/problems/sliding-window-maximum/solution/dan-diao-dui-lie-by-labuladong/
对这道题目有了初步的印象：

典型的用数据结构解决问题的思想：
单调递减队列实现
维护一个单调递减队列，如果元素大于队尾元素，将队尾元素移掉 push 操作的时候判断
pop 操作，如果队头元素是最大值，则移除
获取最大值，因为是单调递减队列，最大值在队列头部

import "container/list"
type Deque struct {
    data *list.List
}
func NewDeque() *Deque {
    return &Deque{data:list.New()}
}
func (q *Deque)Push(val int) {
    for q.data.Len() != 0 && val > q.data.Back().Value.(int) {
        q.data.Remove(q.data.Back())
    }
    q.data.PushBack(val) 
}
func (q *Deque)Pop(val int) {
    if q.data.Len() > 0 && q.Max() == val {
        q.data.Remove(q.data.Front())
    }
}
func (q *Deque)Max() int {
    return q.data.Front().Value.(int)
}
func maxSlidingWindow(nums []int, k int) []int {
    res := []int{}
    window := NewDeque()
    for i,v := range nums {
        if i < k - 1 {
            window.Push(v)
        }else {
            window.Push(v)
            res = append(res, window.Max())
            window.Pop(nums[i - k + 1])
        }
    }
    return res
}

时间复杂度分析：O(n): 遍历一遍 nums, 获取结果
空间复杂度分析：O(n): 维护了一个队列 以及存储结果数组

第三种是用数组模拟了单调队列的实现，对边界条件以及加入结果数组进行判断，感觉第二种方法是通用的解决方法，理解了第二种，第三种方法是它的另外一种表现形式

func maxSlidingWindow(nums []int, k int) []int {
    if nums == nil {
        return []int{}
    }
    window,res := []int{},[]int{}
    for i,x := range nums {
        if i >= k && window[0] <= i - k {
            window = window[1:]
        }
        for len(window) > 0 && nums[window[len(window) - 1]] <= x {
            window = window[:len(window) - 1]
        }
        window = append(window, i)
        if i >= k - 1 {
            res = append(res, nums[window[0]])
        }
    }
    return res
}

复杂度分析：
复杂度和第二种方法一样

总结：
1. 刚做过单调栈，遇到单调队列的时候又没有思路了。。。

2. 好的数据结构可以减少程序运行的时间，题目中用了单调队列，还是挺巧妙的

3. 先把最常见的方法掌握和熟练，然后不断再去优化和演化,不要贪心，一步一步来
