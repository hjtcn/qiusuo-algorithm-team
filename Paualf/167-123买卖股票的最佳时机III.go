给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
示例 1:
输入：prices = [3,3,5,0,0,3,1,4]
输出：6
解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
示例 2：
输入：prices = [1,2,3,4,5]
输出：4
解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。   
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。   
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例 3：
输入：prices = [7,6,4,3,1] 
输出：0 
解释：在这个情况下, 没有交易完成, 所以最大利润为 0。
示例 4：
输入：prices = [1]
输出：0
 
提示：
1 <= prices.length <= 105
0 <= prices[i] <= 105

    1. Clarfication:
        动态规划五部曲：
      1. dp方程含义
        最多可以完成两笔交易
        d[i][j][k]: 第i天，j今天买入or卖出，k是交易次数
        语意：第i天，满足交易次数的情况下，买入or卖出的最大利润
        dp[i][0][0] = max(dp[i-1][0][0])
        dp[i][0][1] = max(dp[])
        dp[i][1][0]
        dp[i][1][1]
    我觉得难的地方还是在于 k 的限制，k的地方没有想清楚
      2. 动态方程
      3. 初始化
      4. 遍历顺序
      5. 推导例子

2. 看题解：
1.  确定dp数组以及下标的含义：
一天一共就有五个状态：
0 没有操作
1 第一次买入
2 第一次卖出
3 第二次买入
4 第二次卖出
dp[i][j] 中i表示第i天，j为[0-4] 5个 状态，dp[i][j] 表示第i天状态j所剩最大现金

2. 确定递推公式：
dp[i][1] 表示的是第i天，买入股票的状态，并不是说一定要第i天买入股票
dp[i][1] 有两个具体操作：
操作一：第i天买入股票了，dp[i][1] = dp[i-1][0] - prices[i]
操作二：第i天没有操作，而是沿用前一天买入的状态，即dp[i][1] = dp[i-1][1]
同理dp[i][2] 也有两个操作：
操作一：第i天卖出股票了，dp[i][2] = dp[i-1][1] + prices[i]
      操作二：第i天没有操作，沿用前一天卖出股票的状态，即：dp[i][2] = dp[i-2][2]
同理
dp[i][3] = max(dp[i-1][3], dp[i-1][2] - prices[i])
dp[i][4] = max(dp[i-1][4], dp[i-1][3] + prices[i])

3. dp数组如何初始化：
第0天没有操作，这个最容易想到，就是0，即dp[0][0] = 0
第0天第一次买入，dp[0][1] = -prices[i]
第0天第一次卖出，dp[0][2] = 0
第0天第二次买入，dp[0][3] = -prices[0]
第0天第二次卖出：dp[0][4] = 0

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    dp := make([][]int,len(prices))
    for i := 0;i < len(dp);i++ {
        dp[i] = make([]int,5)
    }
    dp[0][1] = -prices[0]
    dp[0][3] = -prices[0]
    for i := 1;i < len(prices);i++ {
        dp[i][0] = dp[i - 1][0]
        dp[i][1] = max(dp[i-1][1],dp[i-1][0] - prices[i])
        dp[i][2] = max(dp[i-1][2],dp[i-1][1]+prices[i])
        dp[i][3] = max(dp[i-1][3],dp[i-1][2] - prices[i])
        dp[i][4] = max(dp[i-1][4],dp[i-1][3] + prices[i])
    }
    return dp[len(prices) - 1][4]
}

func max(x,y int)int {
    if x > y {
        return x
    }
    return y
}

官方优化过空间的代码：
func maxProfit(prices []int) int {
    buy1, sell1 := -prices[0], 0
    buy2, sell2 := -prices[0], 0
    for i := 1; i < len(prices); i++ {
        buy1 = max(buy1, -prices[i])
        sell1 = max(sell1, buy1+prices[i])
        buy2 = max(buy2, sell1-prices[i])
        sell2 = max(sell2, buy2+prices[i])
    }
    return sell2
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4. 总结：
4.1: 陷入思维定式，印象中需要加k去处理，其实使用j将所有5种状态表示清楚就行了

4.2: 状态集还是要分析清楚的，从那里会变成什么？
