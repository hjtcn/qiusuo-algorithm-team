/*
121. 买卖股票的最佳时机
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

*/

/*
    思路：一开始读错题了，以为可以多次买卖，一顿操作，发现只能买卖一次。
        以贪心的思路思考就是：
        以最低的股票买入，不断的更新最大差值，即为最大利润值。
        如果当前价格>买入价，就可以更新最大差值
        如果当前价格<=买入价，就更新买入价。

        刻意使用动规去思考一下
        1.dp数组的意义
        dp[j]第j天的最大利润值
        2.状态转移方程
        不断更新买入价，从而不断更新利润值。
        dp[j]=Math.max(dp[j-1],prices[j]-buy)
        3.dp初始化
        dp=Array(len).fill(0)
        4.遍历顺序。
        只买卖一次，只遍历一次股价，不断更新即可
        5.举例
*/

var maxProfit = function(prices) {
    let buy=prices[0],sell=prices[0]
    let max=0
    let len=prices.length
    for(let i=1;i<len;i++){
        if(prices[i]>buy){
            sell=prices[i]
            max=Math.max(max,sell-buy)
        }
        else{
            buy=prices[i]
        }
    }
    return max
};
/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

var maxProfit = function(prices) {
    let buy=prices[0]
    let len=prices.length
    let dp=Array(len).fill(0)
    for(let i=1;i<len;i++){
        buy=Math.min(prices[i],buy)
        dp[i]=Math.max(dp[i-1],prices[i]-buy)
    }
    return dp[len-1]
}
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：这种题更习惯以贪心的角度思考。
*/