给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
示例 1：
输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
示例 2：
输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
 
提示：
1 <= prices.length <= 105
0 <= prices[i] <= 104

1. Clarification:
 定义 min = prices[0]
    ret := 0 
    从 i := 1;i < len(prices);i++ {
        if prices[i] < min {
            min = prices[i]
        }
        if prices[i] - min > ret {
            ret = prices[i] - min
        }
    }
    核心关键：
        更新min，然后去和min去比较

func maxProfit(prices []int) int {
    min := prices[0]
    ret := 0
    for i := 1;i < len(prices);i++ {
        if prices[i] < min {
            min = prices[i]
        }
        if prices[i] - min > ret {
            ret = prices[i] - min
        }
    }
    return ret
}

2. 看题解：

动态规划：
1. 确定dp数组（dp  table) 以及下标的含义
dp[i][0] 表示第 i 天持有股票所得最多现金
dp[i][1] 表示第 i 天不持有股票所得最多现金

2. 确定递推公式
如果第 i 天持有股票即 dp[i][0] :
第 i - 1天就持有股票，dp[i-1][0]
第i天买入股票，所得现金就是买入今天的股票后所得现金即：-prices[i]
dp[i][0] = max(dp[i-1][0], - prices[i])
如果第i天不持有股票即dp[i][1]:
第 i - 1 天就不持有股票，dp[i-1][1]
第i天卖出骨片，所得现金就是按照今天股票价格卖出后所得现金即：prices[i] + dp[i-1][0]
dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])

3. dp 初始化
由递推公式可以从dp[0][0] 和 dp[0][1] 推导出来

4. 确定遍历顺序
dp[i] 从 dp[i-1]推导出来，所以是从前向后遍历

5. 举例推导dp数组

func maxProfit(prices []int) int {
    length:=len(prices)
    if length==0{return 0}
    dp:=make([][]int,length)
    for i:=0;i<length;i++{
        dp[i]=make([]int,2)
    }
    
    dp[0][0]=-prices[0]
    dp[0][1]=0
    for i:=1;i<length;i++{
        dp[i][0]=max(dp[i-1][0],-prices[i])
        dp[i][1]=max(dp[i-1][1],dp[i-1][0]+prices[i])
    }
    return dp[length-1][1]
}
func max(a,b int)int {
    if a>b{
        return a 
    }
    return b 
}
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/dai-ma-sui-xiang-lu-121-mai-mai-gu-piao-odhle/

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4. 总结：
4.1: dp 定义蛮有意思的，把一个问题转化成计算机可以理解的状态了 =》 数组去记录一个大问题的一个个小问题
