/*
188. 买卖股票的最佳时机 IV
给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

 

示例 1：

输入：k = 2, prices = [2,4,1]
输出：2
解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
示例 2：

输入：k = 2, prices = [3,2,6,5,0,3]
输出：7
解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。
 

提示：

0 <= k <= 100
0 <= prices.length <= 1000
0 <= prices[i] <= 1000

*/

/*
    思路：1.没有想到拆分为买卖的两种状态。盲目的先假设
            dp[i][j]代表第i天完成j笔交易的最大利润
         2.也没有推出来状态转移方程
    好吧，停止，看看题解。。。
    题解：1.拆分两种买卖状态
         buy[i][j]代表第i天完成j笔买入时的最大利润
         sell[i][j]代表第i天完成j笔收入卖出的最大利润
         2.状态转移方程
         1）保证手里有股票
         买入的最大利润=max(前一天买入的最大利润(无操作),前一天卖出的最大利润-price[i](真实买入)）
         buy[i][j]=Math.max(buy[i-1][j],sell[i-1][j]-price[i])
         2）手里无股票
         卖出的最大利润=max(前一天卖出的最大利润(无操作),前一天买入的最大利润(j-1)+price[i](真实卖出)）
         sell[i][j]=Math.max(sell[i-1][j],buy[i-1][j-1]-price[i])

         3.buy和sell初始化以及边界情况
         一开始我只初始化了为0，
         然后发现buy[0][j],sell[0][j]需要确立边界，赋值为-Infinity.
         刚开始有点想不通，为什么不通过全部赋值为-Infinity来初始化呢？
         后来仔细研究了一下，边界之所以作为边界，是意味着超出边界，是无法进行操作的。
         比如：buy[0][j]，j代表[0,j]笔交易。
         只有一天的价格。prices[0]，其实一笔交易都完不成。
         也就是除了buy[0][0]是有意义的，其他j>0的操作都是没有意义的。因此，将其赋值-Infinity，无论它加多大的股票价格，都无法出来利润。
         其它有意义(或者说可出现)的情况初始化需要为0，这样各种操作产生的利润值也不会被初始化影响。
         同时buy[0][0]=-prices[0]
            sell[0][0]=0
        是最初始的状态。
        4.遍历顺序
        先遍历股票价格，再控制交易次数
        5.举例

    看了看代码随想录的题解，感觉更有了从2～k拓展的感觉。
    找规律的过程，真的考验勇气。
    1.dp[i][j] 第i天最多完成j/2次交易
        j的状态列举：
        0:无操作
        1:第一天买入股票的状态
        2:第一天卖出。。。
        3:第二天买入。。。
        4:第二天卖出。。。
        。。。
        0无操作,奇数买入，偶数卖出
        j的数据范围为[0.2k]
    2.状态转移方程：
        //买入的可能状态
        dp[i][j+1]=Math.max(dp[i-1][j+1],dp[i-1][j]-price[i])
        //卖出的可能状态
        dp[i][j+2]=Math.max(dp[i-1][j+2],dp[i-1][j+1]+price[i])

        for循环遍历j时，由于j+1,j+2的存在,且同时处理两个。
        则for(j=0;j<2*k-1;j+=2)
    3.dp初始化
        dp=Array.from(Array(len),()=>Array(2*k+1).fill(0))
        dp[0][1] dp[0][3]...dp[2*k-1]都为-price[0]
        其余都为0，不做处理
    4.遍历顺序
        先遍历股票，再交易状态
        且从小到达递推
    5.举例


*/
var maxProfit = function(k, prices) {
    let len=prices.length
    if(!len){
        return 0
    }
    k = Math.min(k, Math.floor(len / 2))
    let buy=Array.from(Array(len),()=>Array(k+1).fill(0))
    let sell=Array.from(Array(len),()=>Array(k+1).fill(0))
      for (let i = 1; i <= k; ++i) {
        buy[0][i] = sell[0][i] = -Number.MAX_VALUE;
    }
    buy[0][0]=-prices[0],sell[0][0]=0
    for(let i=1;i<len;i++){
        buy[i][0]=Math.max(buy[i-1][0],sell[i-1][0]-prices[i])
        for(let j=1;j<=k;j++){
            buy[i][j]=Math.max(buy[i-1][j],sell[i-1][j]-prices[i])
            sell[i][j]=Math.max(sell[i-1][j],buy[i-1][j-1]+prices[i])
        }
    }
    return Math.max(...sell[len-1])
};

/*
    时间复杂度：O(len*min(len/2,k))
    空间复杂度：O(len*min(len/2,l))
*/


var maxProfit = function(k, prices) {
    let len=prices.length
    if(!len){
        return 0
    }
    let dp=Array.from(Array(len),()=>Array(2*k+1).fill(0))
      for (let i = 1; i < 2*k;i+=2) {
        dp[0][i] = -prices[0]
    }
    for(let i=1;i<len;i++){
        for(let j=0;j<2*k-1;j+=2){
            dp[i][j+1]=Math.max(dp[i-1][j+1],dp[i-1][j]-prices[i])
            dp[i][j+2]=Math.max(dp[i-1][j+2],dp[i-1][j+1]+prices[i])
        }
    }
    // console.log(sell[len-1])
    return dp[len-1][2*k]
};
/*
    时间复杂度：O(len*k)
    空间复杂度：O(len*k)
*/

var maxProfit = function(k, prices) {
    let len=prices.length
    if(!len){
        return 0
    }
    let dp=Array(2*k+1).fill(0)
      for (let i = 1; i < 2*k;i+=2) {
        dp[i] = -prices[0]
    }
    for(let i=1;i<len;i++){
        for(let j=0;j<2*k-1;j+=2){
            dp[j+1]=Math.max(dp[j+1],dp[j]-prices[i])
            dp[j+2]=Math.max(dp[j+2],dp[j+1]+prices[i])
        }
    }
    // console.log(sell[len-1])
    return dp[2*k]
};

/*
    时间复杂度：O(len*k)
    空间复杂度：O(k)
*/

/*
    限制从2～k,很多题目的扩展都是这样开始的，要学会自己往深处思考。
*/