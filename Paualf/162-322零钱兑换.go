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
      最少的硬币个数
        dp[amount] = min(dp[amount - coins[i]])
        
        动态规划五部曲
        1. dp[j], 组成j的总金额使用的硬币数
        
        2. dp 公式
          dp[amount] = min(dp[amount - conis[i]])
        3. dp数组如何初始化
            dp[0] = 0
        4. 遍历顺序

Coding:
没有想清楚的思路。。。
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
    dp[0] = 1
    ret := amount + 1
    
    for _,coin := range coins {
        dp[coin] = 1
    }
    fmt.Println(dp)
    for  i := 1;i <= amount;i++ {
        for _,coin := range coins {
         if i - coin >= 0 {
                dp[i] = dp[i] + dp[i - coin]
                if dp[i] < ret {
                    ret = dp[i]
                }
            }
        }
    }
    
    fmt.Println(dp)
    return ret
}

2. 看题解：
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
   
    for i := 1;i < len(dp);i++ {
        dp[i] = amount + 1
    }
    
    dp[0] = 0
    for  i := 1;i <= amount;i++ {
        for _,coin := range coins {
            if coin <= i {
                dp[i] = min(dp[i],dp[i - coin] + 1)
            }
        }
    }
    
    if dp[amount] <= amount {
        return dp[amount]   
    }
    
    return -1
}

func min(x,y int)int {
    if x < y {
        return x
    }
    return y
}

3. 复杂度分析：
时间复杂度：O(amount * O(len(coins))
空间复杂度：O(n)

4. 总结：
4.1： 没有想出来的是里面的 for 循环，虽然思路是想出来了，里面的 coin <= i，然后去比较min，没有想出来

4.2:  初始化也是蛮重要的一步，初始化 amount + 1, dp[0] =0 ，哪一步错了都执行不正确的

4.3:  代码还是要多写，有思路没有用的，思路转换不成代码，转换不成产品，则都是纸上谈兵
