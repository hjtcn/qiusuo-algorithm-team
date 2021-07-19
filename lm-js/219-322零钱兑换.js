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
    思路:1.dp数组的含义
          dp[j]为凑成总金额为j时所需的最少的硬币个数。
        2.状态转移方程
        dp[j]=Math.min(dp[j],dp[j-coins[i]]+1)
        3.dp初始化
        求最小值？初始化都为Infinity
        dp[0]=0,凑0块不需要硬币
        4.遍历顺序
        特意关注了这一点，针对这道题:
        先遍历硬币，后金额
        先遍历金额，后硬币
        两周遍历方式无所谓前后。
        5.举例推导

      坑点难点:1.初始化的过程，是否能想到dp[0]=0
             2.状态转移方程，求最小值？
             当前dp[j]和上一步dp[j-coins[i]]的关系？
             dp[j]=Math.min(dp[j],dp[j-coins[i]]+1)
*/
var coinChange = function(coins, amount) {
    let len=coins.length
    let dp=Array(amount+1).fill(Infinity)
    dp[0]=0
    for(let i=0;i<len;i++){
        for(let j=1;j<=amount;j++){
            if(j-coins[i]>=0){
                dp[j]=Math.min(dp[j],dp[j-coins[i]]+1)
            }
        }
    }
    return dp[amount]==Infinity?-1:dp[amount]
};

/*
    时间复杂度：O(N*amount)
    空间复杂度：O(amount)
*/


/*
    思考：不要熟悉框架的同时，丢失了思考状态转移的过程。
 */