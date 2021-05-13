/*
122. 买卖股票的最佳时机 II
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

*/

/*
    思路：贪心就是摸索着走着，买卖股票就是买入最小金额，卖出最大金额。
        遍历数组.
        当前单价大于前一天的价格，则更新买入价（min(buy,prices[i-1])），卖出价(prices[i]);
        当前单价小于前一天的价格，说明前一天可以卖出了，count+=sell-buy.并且更新买入价，卖出价都为当前单价prices[i].
        跳出循环，最后如果sell-buy不为0，说明最后连续递增，则返回count+sell-buy.
    这个思路是自己试探着思考提交的，确实自己也不太确定有没有什么漏洞。

    题解：(1)发现人家思路太简单了，以前我们好像也遇到过这样的问题。核心思路：只加正的。
    
        (2)难道真的万物皆可动规。不能限制在一个框框里，更何况，现在贪心的思路模版也没有特别清晰。
        不过这个dp二维数组的定义很特别。
        1)dp[i][0]代表交易完手里没股票了。两种情况。一种是卖掉了。一种是当天没交易。
        dp[i][0]=Math.max(dp[i][0],dp[i][1]+prices[i])
        2)dp[i][1]代表当天交易完手里有股票。两种情况。一种是当天湄交易，一种是买入了。
        dp[i][1]=Math.max(dp[i][1],dp[i][0]-price[i])

        js提交二维数组又遇到了溢出，直接开始优化优化。

        最绝望的是，经历了一个假期，把二维数组的初始化给忘掉了。稳住啊啊啊啊啊
*/

// @lc code=start
/**
 * @param {number[]} prices
 * @return {number}
 */
 var maxProfit = function(prices) {
    let buy=prices[0],sell=prices[0],count=0
    for(let i=1;i<prices.length;i++){
        if(prices[i]>prices[i-1]){
            buy=Math.min(prices[i-1],buy)
            sell=prices[i]
        }
        else{
            count+=sell-buy
            buy=prices[i]
            sell=prices[i]
        }
    }
    return count+sell-buy
};
//只加正的
var maxProfit = function(prices) {
    let count=0,len=prices.length
    if(len<2){
        return 0
    }
    for(let i=1;i<len;i++){
       let diff=prices[i]-prices[i-1]
       if(diff>0){
           count+=diff
       }
    }
    return count
};

//动规..二维数组，会溢出。优化优化
var maxProfit = function(prices) {
    let len=prices.length
    if(len<2){
        return 0
    }
    let dp=Array.from(Array(len),()=>Array(len).fill(0))
    dp[0][1]=-prices[0]
    for(let i=1;i<len;i++){
      dp[i][0]=Math.max(dp[i-1][0],dp[i-1][1]+prices[i])
      dp[i][1]=Math.max(dp[i-1][1],dp[i-1][0]-prices[i])
    }
    return dp[len-1][0]
};


//动态规划优化。
var maxProfit = function(prices) {
    let len=prices.length
    if(len<2){
        return 0
    }
    let noHold=0,hold=-prices[0]
    for(let i=1;i<len;i++){
        noHold=Math.max(noHold,hold+prices[i])
        hold=Math.max(hold,noHold-prices[i])
    }
    return noHold
};

/*
    时间复杂度：O(N)
    空间复杂度：O(1)  
*/

/*
    思考：贪心代码都不能，难的是如何找到毫无漏洞的贪心策略。加油，慢慢练习。
        之前做过的也不能做过就完事，太放松遗忘的会特别快，哈哈哈

*/