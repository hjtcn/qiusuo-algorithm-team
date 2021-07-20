给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
示例 1:
输入: prices = [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
示例 2:
输入: prices = [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例 3:
输入: prices = [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 
提示：
1 <= prices.length <= 3 * 104
0 <= prices[i] <= 104

    1. Clarification:
        动态规划5部曲：
       
       1.动态规划定义：
        dp[i][0]: 目前买入的时候的最大收入:有可能是上次买入的最大收入，也有可能是上次卖出，这次买入的收入
        dp[i][1]: 目前卖出的时候的最大收入：有可能是上次卖出的时候最大收入，也有可能是上次买入，这次卖出的时候的最大收入
     
       2. dp方程
          dp[i][0] = max(dp[i - 1][0],dp[i-1][1] - price[i])
          dp[i][1] = max(dp[i - 1][1],dp[i-1][0] + prices[i])

      3. dp初始化
         dp[0][0] = -prices[0]
         dp[0][1] = 0

    4. 遍历顺序
        从小到到大，顺序遍历即可
 
    5. 例子推导
       dp[1][0] = max(dp[0][0],dp[0][1]-prices[i]) = max(-7,-1) = -1
       dp[1][1] = max(dp[0][1],dp[0][0] + prices[i]) = max(0,-6) = 0
       dp[2][0] = max(dp[1][0],dp[1][1] - prices[i]) = max(-1,-5) = -1
       dp[2][1] = max(dp[1][1],dp[1][0] + prices[i]) = max(0,-1 + 5) = 4

coding:
func maxProfit(prices []int) int {
    if len(prices) <= 1{
        return 0
    }
    n := len(prices)
   // fmt.Println(n)
    dp := make([][]int, n)
    for i := 0;i < n;i++ {
        dp[i] = make([]int, 2)
    }
    dp[0][0] = -prices[0]
    dp[0][1] = 0
    for i := 1;i < n;i++ {
         dp[i][0] = max(dp[i - 1][0], dp[i-1][1] - prices[i])
         dp[i][1] = max(dp[i - 1][1], dp[i-1][0] + prices[i])
    }
    return dp[n - 1][1]
}

func max(x,y int)int {
    if x > y {
        return x
    }
    return y
}

debug:
  debug 的时候在 n := len(prices) 初始化的时候有点没有想清楚，初始化定义的时候定义了 n := len(prices) + 1

2. 看题解：
压缩空间：
func maxProfit(prices []int) int {
    n := len(prices)
    dp0, dp1 := 0, -prices[0]
    for i := 1; i < n; i++ {
        dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
    }
    return dp0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4. 总结：
4.1: 和第一道题不一样的地方在于，第一道题只能买卖一次，这道题可以进行多次交易，多次交易的时候必须把上次的给卖出去
4.2: 发现动态规划的题目需要提前把题目分析清楚，然后写代码，这个也是可以让我们遇到问题先分析定义清楚问题以后再下手，会好得多
