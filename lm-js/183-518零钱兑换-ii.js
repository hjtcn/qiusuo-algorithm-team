/*
518. 零钱兑换 II
给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。 

 

示例 1:

输入: amount = 5, coins = [1, 2, 5]
输出: 4
解释: 有四种方式可以凑成总金额:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
示例 2:

输入: amount = 3, coins = [2]
输出: 0
解释: 只用面额2的硬币不能凑成总金额3。
示例 3:

输入: amount = 10, coins = [10] 
输出: 1
 

注意:

你可以假设：

0 <= amount (总金额) <= 5000
1 <= coin (硬币面额) <= 5000
硬币种类不超过 500 种
结果符合 32 位符号整数
*/


/*
    思路：寻找能组合金额数的组合数。上模版
        for(let i=0;i<len;i++){
            for(let j=amount;j>=coins[i];j--){
                dp[j]+=dp[j-coins[i]]
            }
        }
        一执行，出错了，嘿，不应该呀。
        在对照五步法则：
        1.确认dp数组及下标的含义
        2.确定状态转移方程
        dp[j]=dp[j]+dp[j-coins[i]]
        3.dp初始化
        dp=Array(amount+1).fill(0)
        dp[0]=1
        4.确认遍历方向
        这个位置和之前的474题(一和零)相比，引发一些思考
        一和零，每个字符只能有一次，是有限背包，需要倒序
        而本题，一种面额的硬币有无数个，为完全背包，可正序
        5.举例推导dp数组.画出二维图

*/



var change = function(amount, coins) {
    let len=coins.length;
    let dp=Array(amount+1).fill(0)
    dp[0]=1
    for(let i=0;i<len;i++){
        for(let j=coins[i];j<=amount;j++){
            dp[j]+=dp[j-coins[i]]
        }
    }
    return dp[amount]
};

/*
    时间复杂度：O(n*amount)
    空间复杂度：O(amount)
*/

/*
   
    01背包：数组中的元素不可重复使用，nums放在外循环，target在内循环，且内循环倒序
    完全背包：数组中的元素可重复使用，nums放在外循环，target在内循环。且内循环正序。

    如果组合问题需考虑元素之间的顺序，需将target放在外循环，将nums放在内循环

*/