输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
要求时间复杂度为O(n)。
示例1:
输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 
提示：
1 <= arr.length <= 10^5
-100 <= arr[i] <= 100

1. Clearfication:
1.1. 贪心，初始化最大值ret 为 nums[0],循环遍历进行累计求和记为currentSum，如果currentSum小于0则将 currentSum 置为0，比较 ret 和 currentSum的大小
1.2. 动态规划,dp[i] 代表累加到当前下标连续子数组和最大,dp[i] = max(dp[i - 1],dp[i - 1] + nums[i]),dp[0] = nums[0]

2. Coding:
func maxSubArray(nums []int) int {
    ret := nums[0]
    currentSum := nums[0]
    for i := 1;i < len(nums);i++ {
        if currentSum < 0 {
            currentSum = 0
        }
        currentSum += nums[i]
        if currentSum > ret {
            ret = currentSum
        }
    }
    return ret
}

注意 对 currentSum 的判断，一开始的时候自己将 currentSum 放到了下面判断，上一次结果如果是负数，则应置为0，所以应该先判断 currentSum 的正负情况，所以这个地方还是没有分析清楚
func maxSubArray(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    dp[0] = nums[0]
    for i := 1;i < n;i++ {
        dp[i] = max(dp[i - 1] + nums[i],dp[i - 1])
    }
    return dp[n - 1]
}

func max(x,y int)int {
    if x > y {
        return x
    }
    return y
}
上面的动态规划方程没有分析正确

3. 看题解：
func maxSubArray(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    dp[0] = nums[0]
    ret := nums[0]
    for i := 1;i < n;i++ {
        dp[i] = max(dp[i - 1] + nums[i], nums[i])
        if dp[i] > ret {
            ret = dp[i]
        }
    }
    return ret
}

func max(x,y int)int {
    if x > y {
        return x
    }
    return y
}

动态规划方程分析错了，是 dp[i] = max(dp[i - 1] + nums[i] ,nums[i]) ,这个不是错的关键，这个自己在调试的时候分析出来了，最没有分析清楚的时候是返回值，自己直接返回了dp[n - 1], 说明自己还是没有理解清楚 dp[i] 的意思，dp[i] 是累加到 下标i的时候的和，如果我们想要找到最大值，需要在dp[i] 数组中找到max(dp[i])，然后进行返回
比较清新的代码：

func maxSubArray(nums []int) int {
    max, sum := nums[0],nums[0]
    
    for _,v := range nums[1:] {
        if sum < 0 {
            sum = v
        }else {
            sum += v
        }
        
        if max < sum {
            max = sum
        }
    }
    
    return max
}

func maxSubArray(nums []int) int {
    currSum,maxSum := nums[0],nums[0]
    
    for i := 1;i < len(nums);i++ {
        currSum = max(nums[i], currSum + nums[i])
        maxSum = max(currSum,maxSum)
    }
    
    return maxSum
}

func max(x,y int) int {
    if x > y {
        return x
    }
    
    return y
}

4. 复杂度分析：
时间复杂度：O(n) : 遍历整个数组
空间复杂度：O(n)： 使用dp数组进行存储结果

5. 总结：
5.1：分析题目的时候写出伪代码的时候，可以run一下，看一些逻辑是否正确，不要着急写代码

5.2: 不要思维定式，要去认真分析，定义状态的含义，然后最后想要返回的含义
