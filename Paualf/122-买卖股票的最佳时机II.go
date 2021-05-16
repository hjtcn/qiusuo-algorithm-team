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
• 1 <= prices.length <= 3 * 104
• 0 <= prices[i] <= 104

1. Clearfication:
想了一下，没有思路，不知道从那里下手
2. Coding:

3. 看题解：
贪心算法：
只要可以盈利就交易,没有想清楚为什么这样盈利是最大的
func maxProfit(prices []int) int {
    ans := 0
    for i := 1;i < len(prices);i++ {
        ans += max(0, prices[i] - prices[i - 1])
    }
    return ans
}
func max(x,y int)int {
    if x > y {
        return x
    }
    return y 
}

看完动态规划后，发现动态规划比贪心更加容易理解，更加清晰一点
func maxProfit(prices []int)int {
    n := len(prices)
    dp := make([][2]int, n)
    dp[0][1] = -prices[0]
    
    for i := 1;i < n;i++ {
        dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i])
        dp[i][1] = max(dp[i-1][1],dp[i-1][0] - prices[i])
    }
    
    return dp[n - 1][0]
}   

func max(x,y int)int {
    if x > y {
        return x
    }
    return y
}

只存储前面两个状态：
func maxProfit(prices []int)int {
    n := len(prices)
    dp0,dp1 := 0,-prices[0]
    for i := 1;i < n;i++ {
        dp0,dp1 = max(dp0,dp1 + prices[i]),max(dp1,dp0 - prices[i])
    }
    return dp0
}
func max(x,y int)int {
    if x > y {
        return x
    }
    return y
}

4. 复杂度分析：
时间复杂度：O(n): 遍历数组
空间复杂度：O(n) ：开辟数组存储空间

5. 总结：
5.1： 感觉动态规划比贪心更明确一点
5.2:  动态规划是不是解决问题的时候需要记录它的状态而已，是一种解决问题的方式和方法，也不需要害怕它
