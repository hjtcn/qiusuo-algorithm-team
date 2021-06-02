/*
    322. 零钱兑换
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

你可以认为每种硬币的数量是无限的。

 

示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3 
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0
示例 4：

输入：coins = [1], amount = 1
输出：1
示例 5：

输入：coins = [1], amount = 2
输出：2
 

提示：

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104


*/
/*
    思路：开始的时候我卡卡就把前两天的模块写出来了，寻找总金额为amount的最小组成数
    for(let i=0;i<len;i++){
        for(let j=amount;j>=coins[i];j--){
            dp[j]=dp[j]+dp[j-coins[i]]
        }
    }
    此时就迷路了，dp[j]到底代表什么？内层的遍历顺序还是倒序吗？

    题解：动规：根据别人的答案，总结自己的做题模版
        1.确认dp数组及下标意义
            dp[j]代表组成总金额为j的所需金币数
        2.状态转移方程
             尝试思考状态转移方程(哈哈，没想出来)
            dp[j]=Math.min(dp[j],dp[j-coins[i]]+1)
        3.dp数组如何初始化
            找最小值。Array(amount+1).fill(Infinity)
        4.确定遍历顺序
            内层遍历顺序。dp[j-coins[i]]+1,
            统计应该从小到大。
            边界 j需要大于coins[i]
        5.举例推导dp数组.画出二维图

      递归：还是细节以及记忆化，没有静下心来寻找各种边界值。

    

*/

/*
 * @lc app=leetcode.cn id=322 lang=javascript
 *
 * [322] 零钱兑换
 */


// dp[j]代表能得到amout的个数。
// @lc code=start
/**
 * @param {number[]} coins
 * @param {number} amount
 * @return {number}
 */
 var coinChange = function (coins, amount) {
    let len = coins.length
    let dp=Array(amount+1).fill(Infinity)
    dp[0]=0
    for(let i=0;i<len;i++){
        for(let j=0;j<=amount;j++){
            if(j>=coins[i]){
                dp[j]=Math.min(dp[j],dp[j-coins[i]]+1)
            }
        }
    }
    return dp[amount]==Infinity?-1:dp[amount]
};
// @lc code=end

var coinChange = function(coins, amount) {
    var arr = new Array(amount+1).fill(-2);
    return dfs(coins,amount,arr);
};
function dfs(coins,amount,arr){
    if (amount <= 0) return 0;
    if(arr[amount] != -2)return arr[amount];
    let min = Infinity;
    for (let i=0;i<coins.length;i++) {
        if (amount >=  coins[i]){
            let tmp = dfs(coins, amount - coins[i],arr);
            if (tmp != -1){
                min = Math.min(min, tmp+1);
            }
        }
    }
    return arr[amount] = Number.isFinite(min) ? min : -1;
}

/*
    时间复杂度：O(N*amount)
    空间复杂度：O(amount)
*/

/*
    思考：套模版也要灵活思考，别害怕
 */
