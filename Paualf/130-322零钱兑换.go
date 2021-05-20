给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
你可以认为每种硬币的数量是无限的。
 
示例 1：
输入：coins = [1, 2, 5], amount = 11
输出：3 
解释：11 = 5 + 5 + 1
示例 2：
输入：coins = [2], amount = 3
输出：-1
示例 3：
输入：coins = [1], amount = 0
输出：0
示例 4：
输入：coins = [1], amount = 1
输出：1
示例 5：
输入：coins = [1], amount = 2
输出：2
 
提示：
1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104

1. Clearfication:
coins[i] * n ++ == amount
能不能组成，自己困惑的点在于可以无限选择同一面额的金币，贪心的话选最大数量的最大面额，然后找小的补上去

怎么实现呢？没想出来。。。

2. Coding:

3. 看题解：
思路类似于青蛙跳楼梯
dp{n} = mind{dp(n - coins[0],dp(n-coins[1]))}
func coinChange(coins []int, amount int)int {
    dp := make([]int, amount + 1)
    
    dp[0] = 0
    
    for i := 1;i <= amount;i++ {
        dp[i] = amount + 1
        for _,coin := range coins {
            if i >= coin {
                dp[i] = min(dp[i], dp[i - coin] + 1)
            } 
        }
    }
    
    if dp[amount] == amount + 1 {
        return -1
    }
    
    return dp[amount]
}
func min(x,y int)int {
    if x < y {
        return x
    }
    return y
}
Go search with memo：代码写的好漂亮
func coinChange(coins []int, amount int) int {  
    dp := make([]int, amount + 1)
    return helper(coins, amount, dp)
}
func helper(coins []int, amount int, dp []int) int {
    if amount == 0 || dp[amount] != 0 {
        return dp[amount]
    }
    
    res := math.MaxInt64
    for _, v := range coins {
        if v == amount {
            res = 1
        } else if v < amount {
            temp := 1 + helper(coins, amount - v, dp)
            if temp != 0 && temp < res {
                res = temp
            }
        }
    }
    
    if res == math.MaxInt64 {
        dp[amount] = -1
        return -1
    }
    
    dp[amount] = res
    return res
}

4. 复杂度分析：
时间复杂度：O(Sn) : S是金额，n是面额数，需要计算O(S)个状态，S为题目给的总金额，每次需要枚举n个面额来转移状态
空间复杂度：O(n)

5. 总结：
5.1: 分解子问题的能力需要培养

5.2: 从后往前分析问题

5.3: take it easy
