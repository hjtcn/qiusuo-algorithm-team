/*
123. 买卖股票的最佳时机 III
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
*/

/*
    思路：一开始思路就有点偏，想着不断的更新保留两个最高的利润值。后来发现如果将偏小的利润值替换掉，剩余的利润值并不是那段时期最大的利润值了。
    case:[1,2,4,2,5,7,2,4,9,0]
    我的错误代码借用了栈，就不放上来了。

    题解：题解感觉既巧妙，又勇敢。
    1.dp数组的意义
    直接列举一个五个状态
    dp[i][0]无操作
    dp[i][1]第一次买入
    dp[i][2]第一次卖出
    dp[i][3]第二次买入
    dp[i][4]第二次卖出
    2.状态转移方程
    (1)无操作
    dp[i][0]=dp[i-1][0]
    (2)第一次买入。当然需要和前一天第一次买入的状态比较，去进行更新
    dp[i][1]=Math.max(dp[i-1][1],dp[i-1][0]-prices[i])
    (3)第一次卖出。当然需要和前一天第一次卖出的状态比较，去进行更新
    dp[i][2]=Math.max(dp[i][1]+prices[i],dp[i-1][2])
    (4)第二次买入
    dp[i][3]=Math.max(dp[i-1][3],dp[i-1][2]-prices[i])
    (4)第二次卖出
    dp[i][4]=Math.max(dp[i-1][4],dp[i-1][3]+prices[i])
    （和前一天的同一种状态比较进行更新的过程中踩坑了，理解偏了，以为应该前一天可能无操作的利润进行对比更新）


*/

var maxProfit = function(prices) {
    let len=prices.length
    let dp=Array.from(Array(len),()=>Array(5).fill(0))
    dp[0][0]=0
    dp[0][1]=-prices[0]
    dp[0][2]=0
    dp[0][3]=-prices[0]
    dp[0][4]=0
    for(let i=1;i<len;i++){
        dp[i][0]=dp[i-1][0]
        dp[i][1]=Math.max(dp[i-1][1],dp[i-1][0]-prices[i])
        dp[i][2]=Math.max(dp[i][1]+prices[i],dp[i-1][2])
        dp[i][3]=Math.max(dp[i-1][3],dp[i-1][2]-prices[i])
        dp[i][4]=Math.max(dp[i-1][4],dp[i-1][3]+prices[i])
    }
    return Math.max(dp[len-1][2],dp[len-1][4])
};
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    优化降维过程
*/
var maxProfit = function(prices) {
    let len=prices.length
    let buy1=-prices[0],sell1=0,buy2=-prices[0],sell2=0
    for(let i=1;i<len;i++){
        buy1=Math.max(buy1,-prices[i])
        sell1=Math.max(buy1+prices[i],sell1)
        buy2=Math.max(buy2,sell1-prices[i])
        sell2=Math.max(sell2,buy2+prices[i])
    }
    return sell2
};
/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/
/*
    玩股票总是想手动模拟，我觉得自己膨胀了，哈哈哈哈
*/