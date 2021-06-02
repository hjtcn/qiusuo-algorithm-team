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
    思路：找到最小买入值，找到最大卖出值。卖出值位置要在买入值索引位置之后，不断更新最大差值。但是超时了，害，还是要贪心，但是没把判断逻辑想清楚。
    题解：贪心：买入值只会不断变小。
        如果当前值小于买入值，则更新买入值
        如果当前值-买入值大于买卖差值，则更新买卖差值。
        循环结束
        返回买卖差值
        动态规划：dp[i][0]第i天结束，不持股，手里的现金数
                dp[i][1]第i天结束，持股，手里的现金数
                js二维数组依旧会溢出
                进一步优化
*/

var maxProfit = function(prices) {
    let max=0,len=prices.length
    for(let i=0;i<len-1;i++){
        for(let j=i+1;j<len;j++){
            if(prices[j]>prices[i]){
                max=Math.max(max,prices[j]-prices[i])
            }
        }
    }
    return max
};

//超时
// @lc code=end

var maxProfit = function(prices) {
    let len=prices.length,buyVal=Infinity,maxDiff=0;
    for(let i=0;i<len;i++){
        if(prices[i]<buyVal){
            buyVal=prices[i]
        }
        else if(prices[i]-buyVal>maxDiff){
            maxDiff=prices[i]-buyVal
        }
    }
    return maxDiff
}


var maxProfit = function(prices) {
    let len=prices.length;
    let dp=Array.from(Array(len),()=>Array(len).fill(0))
    dp[0][1]=-prices[0]
    for(let i=1;i<len;i++){
        dp[i][0]=Math.max(dp[i-1][0],dp[i-1][1]+prices[i])
        dp[i][1]=Math.max(dp[i-1][1],-prices[i])
    }
    return dp[len-1][0]
}


//二维数组优化
var maxProfit = function(prices) {
    let len=prices.length;
    let hold=-prices[0],noHold=0
    for(let i=1;i<len;i++){
        noHold=Math.max(noHold,hold+prices[i])
        hold=Math.max(hold,-prices[i])
    }
    return noHold
}

/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

/*
    思考：才刚过半个月，我可又忘完了。
         难过。这题的核心就在于dp[i][0]，dp[i][1]的定义。
*/