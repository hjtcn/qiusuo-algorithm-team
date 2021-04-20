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
一上来就想写动态规划，因为动态规划方程自己是会写的，想了下动态规划的本质是找问题的最优子结构。也就是问题有多个子问题。
那么递归和分治也是可以解决这类问题的，我们在爬楼梯那道题里面就看到了，先使用递归，然后使用记忆话数组，最后使用动态规划，然后优化使用空间。
其中优化空间哪一步，不知道是不是自己害怕，还是怎么样，感觉那样的代码蛮难懂的，写出来以后工程上别人很难理解，所以自己在优化空间上面花的时间比较少。

2. Coding:
/*
    分解子问题：
        [10,15,20]
*/
func helper(cost []int,n int) int {
    // terminator
    if n <= -1 {
        return 0
    }
    var tmp1,tmp2 int
    // process current logic
    if n - 1 >= 0 {
        tmp1 = cost[n - 1] + helper(cost,n-1)
    }
    
    if n - 2 >= 0 {
        tmp2 = cost[n - 2] + helper(cost, n - 2)
    }
    return min(tmp1,tmp2)
}

func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    return helper(cost, n)
}

func min(x,y int)int {
    if x < y {
        return x
    }
    return y
}

递归：不出意料的超时了
递归模版
func recur(level int,param int) {
    // terminator
    if level > MAX {
        return
    }
    
    // process current logic
    process(level,param)
    
    // drill down
    recur(level + 1,newParam)
    
    // restore current status
} 

记忆话数组
/*
    分解子问题：
        [10,15,20]
*/
func helper(cost []int,n int,memo[] int) int {
    // terminator
    if n <= -1 {
        return 0
    }
    if memo[n] > 0 {
        return memo[n]
    }
    var tmp1,tmp2 int
    // process current logic
    if n - 1 >= 0 {
        tmp1 = cost[n - 1] + helper(cost,n-1, memo)
    }
    
    if n - 2 >= 0 {
        tmp2 = cost[n - 2] + helper(cost, n - 2, memo)
    }
    memo[n] = min(tmp1,tmp2)
    return memo[n]
}

func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    memo := make([]int, n + 1)
    return helper(cost, n,memo)
}
func min(x,y int)int {
    if x < y {
        return x
    }
    return y
}

动态规划
/*
    dp[n]:到达n层需要的花费
    
    动态规划方程：
         dp[n] = mind(dp[n - 1] + cost[n - 1],dp[n - 2] + cost[n - 2])
    初始化：
        n := len(cost)
        ]
        if len(cost) == 0 {
            return 0
        }
        if len(cost) == 1{
            return cost[0]
        }
        if len(cost) == 2 {
            return min(cost[0], cost[1])
        }
    
        for i := 2;i <= n;i++ {
            dp[i] = min(cost[i - 1] + cost[i - 1],cost[i - 2] + cost[i - 2])
        }
        return dp[n]
*/
func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return cost[0]
    }
    if n == 2 {
        return min(cost[0],cost[1])
    }
    dp := make([]int, n + 1)
    for i := 2;i <= n;i++ {
        dp[i] = min(cost[i - 1] + dp[i - 1], cost[i - 2] + dp[i - 2])
    }
    return dp[n]
}
func min(x,y int)int {
    if x < y {
        return x
    }
    return y
}

3. 看题解
看到题解里面都是直接从2开始的，看了下题目，cost的长度是从[2,1000] => 没有仔细看题目。。。

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) 记忆话数组使用空间 or 动态规划使用空间

5. 总结：
5.1: 本质上还是要找到问题的子问题进行解决 =》 分解子问题
