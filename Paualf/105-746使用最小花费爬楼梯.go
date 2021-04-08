数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。
每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。
请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。
 
示例 1：
输入：cost = [10, 15, 20]
输出：15
解释：最低花费是从 cost[1] 开始，然后走两步即可到阶梯顶，一共花费 15 。
 示例 2：
输入：cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
输出：6
解释：最低花费方式是从 cost[0] 开始，逐个经过那些 1 ，跳过 cost[3] ，一共花费 6 。
 
提示：
• cost 的长度范围是 [2, 1000]。
• cost[i] 将会是一个整型数据，范围为 [0, 999] 。

1. Clearfication:
我是这样分析的，要到达第n阶台阶：它可以从 n - 1 和 n - 2 到达，它们到达的时候要加上 n - 1 台阶的花费和 n - 2 台阶的花费
所以是这样定义的：
1.1 定义：dp[n]：到达第n阶台阶需要的花费
1.2: 方程：dp[n] = min(dp[n - 1] + cost[n - 1],dp[n - 2] + cost[n - 2])
1.3：初始化：
   dp[0] = 0
   dp[1] = 0

2. Coding:
/*
   可以选择从下标为0或1的元素作为初始阶梯 
   dp[n]：到达第n阶台阶需要的花费
   dp[n] = min(dp[n - 1] + cost[n - 1],dp[n - 2] + cost[n - 2])
   dp[2] = min(dp[0] + 10,dp[1] + 15)
   dp[0] = 0
   dp[1] = 0
*/
func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    if n == 1 {
        return cost[0]
    }
    if n == 2 {
        return min(cost[0],cost[1])
    }
    dp := make([]int, n)
    dp[0] = 0
    dp[1] = 0
    for i := 2;i < n;i++ {
        dp[i] = min(dp[i - 1] + nums[i - 1],dp[i - 2] + nums[i - 2])
    } 
    return dp[n - 1]
}

func min(x,y int)int {
    if x < y {
        return x
    }
    return y
}
上面的代码需要进行调试一波：
debug:
调试的时候,到达顶峰，所以要是n，所以下标是可以等于n的
/*
   可以选择从下标为0或1的元素作为初始阶梯 
   dp[n]：到达第n阶台阶需要的花费
   dp[n] = min(dp[n - 1] + cost[n - 1],dp[n - 2] + cost[n - 2])
   dp[2] = min(dp[0] + 10,dp[1] + 15)
   dp[0] = 0
   dp[1] = 0
*/
func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    if n == 1 {
        return cost[0]
    }
    if n == 2 {
        return min(cost[0],cost[1])
    }
    dp := make([]int, n + 1)
    dp[0] = 0
    dp[1] = 0
    for i := 2;i <= n;i++ {
        dp[i] = min(dp[i - 1] + cost[i - 1],dp[i - 2] + cost[i - 2])
    } 
    return dp[n]
}

func min(x,y int)int {
    if x < y {
        return x
    }
    return y
}

3. 看题解：
func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    dp := make([]int, n + 1)
    for i := 2;i <= n;i++ {
        dp[i] = min(dp[i - 1] + cost[i - 1],dp[i - 2] + cost[i - 2])
    } 
    return dp[n]
}

func min(x,y int)int {
    if x < y {
        return x
    }
    return y
}
使用最小花费爬楼梯

还有滚动数组的方法，去节省空间，现阶段个人觉得将动态规划方程推导出来去计算比去节省空间更重要

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度:  O(n)

5. 总结：
5.1: 多多练习锻炼和训练自己的思维方式
