/*
309. 最佳买卖股票时机含冷冻期
给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
示例:

输入: [1,2,3,0,2]
输出: 3 
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
*/
/*
    思路：终于学会了区分状态。
        1.dp的含义
          1）dp[i][0]第i天买入态
          2）dp[i][1]第i天卖出态
          3）dp[i][2]第i天冷冻态
        2.状态转移方程
        dp[i][0]=Math.max(dp[i-1][0],dp[i-1][2]-prices[i])
        dp[i][1]=Math.max(dp[i-1][1],dp[i-1][0]+prices[i])
        dp[i][2]=Math.max(dp[i-1][2],dp[i-1][1])
        返回max(dp[len-1][1],dp[len-1][2])
        3.dp初始化
        dp=Array.from(Array(len),()=>Array(3).fill(0))
        dp[0][0]=-prices[0]
        4.遍历顺序
        按天数从小到达，遍历股票价格
        5.举例

    
*/
var maxProfit = function(prices) {
    let len=prices.length
    let dp=Array.from(Array(len),()=>Array(3).fill(0))
    dp[0][0]=-prices[0]
    for(let i=1;i<len;i++){
        dp[i][0]=Math.max(dp[i-1][0],dp[i-1][2]-prices[i])
        dp[i][1]=Math.max(dp[i-1][1],dp[i-1][0]+prices[i])
        dp[i][2]=Math.max(dp[i-1][2],dp[i-1][1])
    }
    return Math.max(dp[len-1][1],dp[len-1][2])
};

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/
//降维
//踩坑：计算冷冻期的状态，需要对比的是更新前的dp1.
var maxProfit = function(prices) {
    let len=prices.length
    let dp0=-prices[0],dp1=0,dp2=0
    for(let i=1;i<len;i++){
        dp0=Math.max(dp0,dp2-prices[i])
        let lastdp1=dp1
        dp1=Math.max(dp1,dp0+prices[i])
        dp2=Math.max(dp2,lastdp1)
    }
    return Math.max(dp1,dp2)
};
/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

/*
    五天的买卖股票题，坚持下去，我要当股神，哈哈哈
*/