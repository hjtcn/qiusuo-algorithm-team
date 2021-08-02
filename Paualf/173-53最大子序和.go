给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 
示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：
输入：nums = [1]
输出：1
示例 3：
输入：nums = [0]
输出：0
示例 4：
输入：nums = [-1]
输出：-1
示例 5：
输入：nums = [-100000]
输出：-100000
 
提示：
1 <= nums.length <= 3 * 104
-105 <= nums[i] <= 105
 
进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。

   1.Clarification:
        贪心：
            定义ret,curretnSum
                currentSum += nums[i]
                if currentSum > ret {
                    ret = currentSum
                }
                if currentSum < 0 {
                    currentSum = 0
                }
        动态规划五部曲：
            1. dp 定义：
                dp[i] 表示 下标为i的位置，[0-i]位置的连续最大和
            2. dp方程
                if dp[i-1] + nums[i] > 0 {
                    dp[i] = dp[i-1] + nums[i]
                }else {
                    dp[i] = 0
                }
            3.初始化
                dp[0] = nums[i]
            4. 遍历顺序 从前往后
            5. case by case
             返回 max(dp)

func maxSubArray(nums []int) int {
    m := len(nums)
    
    if m <= 0 {
        return 0
    }
    dp := make([]int, m)
    ret := nums[0]
    dp[0] = nums[0]
    for i := 1;i < m;i++ {
        if dp[i-1] + nums[i] > 0 {
            dp[i] = dp[i - 1] + nums[i]
        }else {
            dp[i] = 0
        }
        if dp[i] > ret {
            ret = dp[i]
        }
    }
    return ret
}

上面的代码在执行 [-2,1] 这个case的时候凉了
说明自己没有分析清楚
为什么要与0 比较呢？如果 > 0 才加上去？为什么要这样写？
不是应该去它本身去比较么，如果>我本身我就加上去，如果不大于我本身就等于我本身，有负数的情况哈，-99,-1 这样
修改后：

func maxSubArray(nums []int) int {
    m := len(nums)
    
    if m <= 0 {
        return 0
    }
    dp := make([]int, m)
    ret := nums[0]
    dp[0] = nums[0]
    for i := 1;i < m;i++ {
        if dp[i-1] + nums[i] > nums[i] {
            dp[i] = dp[i - 1] + nums[i]
        }else {
            dp[i] = nums[i]
        }
        if dp[i] > ret {
            ret = dp[i]
        }
    }
    return ret
}

func maxSubArray(nums []int) int {
    m := len(nums)
    if m <= 0 {
        return 0
    }
    ret,currentSum := nums[0],nums[0]
    for i := 1;i < m;i++ {
        if currentSum + nums[i] > nums[i] {
            currentSum = currentSum + nums[i]
        }else {
            currentSum = nums[i]
        }
        if currentSum > ret {
            ret = currentSum
        }
    }
    
    return ret
}

2. 看题解：
看到了线段树
func maxSubArray(nums []int) int {
    return get(nums, 0, len(nums) - 1).mSum;
}
func pushUp(l, r Status) Status {
    iSum := l.iSum + r.iSum
    lSum := max(l.lSum, l.iSum + r.lSum)
    rSum := max(r.rSum, r.iSum + l.rSum)
    mSum := max(max(l.mSum, r.mSum), l.rSum + r.lSum)
    return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
    if (l == r) {
        return Status{nums[l], nums[l], nums[l], nums[l]}
    }
    m := (l + r) >> 1
    lSub := get(nums, l, m)
    rSub := get(nums, m + 1, r)
    return pushUp(lSub, rSub)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

type Status struct {
    lSum, rSum, mSum, iSum int
}

看题解介绍，这种做法在本道题里面没有用，当时想了下，如果像题解中说的那样，将所有子区间的信息用堆栈存储，记忆话下来

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) or O(1)

4. 总结：
4.1: 分析题目的时候动态规划方程没有分歧清楚

4.2: 越来越感觉写代码是要有施工图的，就像技术方案一样，施工地图画的越详细，后续改动越少，结构就越稳定
