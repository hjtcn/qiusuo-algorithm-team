给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。 
示例 1:
输入: amount = 5, coins = [1, 2, 5]
输出: 4
解释: 有四种方式可以凑成总金额:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
示例 2:
输入: amount = 3, coins = [2]
输出: 0
解释: 只用面额2的硬币不能凑成总金额3。
示例 3:
输入: amount = 10, coins = [10] 
输出: 1
 
注意:
你可以假设：
0 <= amount (总金额) <= 5000
1 <= coin (硬币面额) <= 5000
硬币种类不超过 500 种
结果符合 32 位符号整数
1 Clearfication:
// 求组合数
数学题变形： x * conins[0] + y * coins[1] + z *coins[2] = amount 求 x,y,z 构成的数量信息
 dp 倒着推导 5 = amount - conins[i] 
 dp[amount] = sum(dp[amount - coins[i]])  dp[n] 的含义是 组合成 n 的总数
 什么时候终止，并且给结果 + 1 呢 ? amount - conins[i] == 0 的时候，给结果 + 1

2. Coding:
func change(amount int, coins []int) int {
    ret := 0
    dp := make([]int, amount + 1)
    for amount >= 0 {
        dp[amount] = 
        for _, i := range conins {
             
        }
    }
   
}

3. 看题解：
我的思路和这个比较像，有一点还是没看懂，为什么i 要从 i -> amout + 1 ?
func change(amount int, coins []int) int {
    var dp = make([]int, amount+1)
    dp[0] = 1
    for _, v := range coins {
        for i := 1; i < amount+1; i++ {
            if i-v >= 0 {
                dp[i] += dp[i-v]
            }
        }
    }
    return dp[amount]
}

好像有点看懂这个题目的意思了
dp[n][amount] 是放入n得到amount的总数
func change(amount int, coins []int) int {
    n := len(coins)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, amount+1)
    }
    for i := 0; i <= n; i++ {
        dp[i][0] = 1
    }
    for i := 1; i <= n; i++ {
        for j := 1; j <= amount; j++ {
            if j-coins[i-1] >= 0 {
                dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }
    return dp[n][amount]
}

这个也是很厉害的: 记忆话数组的搜索
type pair [2]int
func change(amount int, coins []int) int {
    memo := make(map[pair]int)
    return findChange(amount, 0, coins, memo)
}
func findChange(amount, index int, coins[]int, memo map[pair]int) int {
    
    if amount == 0 {
        return 1
    }
    
    if amount < 0 || index > len(coins) - 1 {
        return 0
    }
    
    if _, ok := memo[pair{amount, index}]; !ok {
        // for each type of coins, we choose to TAKE or NOT TAKE in the current combo
        memo[pair{amount, index}] = findChange(amount - coins[index], index, coins, memo) + findChange(amount, index + 1, coins, memo)
    }
    return memo[pair{amount, index}]
}

4. 复杂度分析：
时间复杂度：O(n*n)
空间复杂度：O(n*n)

5. 总结
5.1: 代码不是很难，难的是思路， 是想法，是去如果解决这个问题和怎么定义这个问题
5.2: 定义问题 -》 定义动态规划方程
