// 746. 使用最小花费爬楼梯
// 数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。

// 每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。

// 请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。

 

// 示例 1：

// 输入：cost = [10, 15, 20]
// 输出：15
// 解释：最低花费是从 cost[1] 开始，然后走两步即可到阶梯顶，一共花费 15 。
//  示例 2：

// 输入：cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
// 输出：6
// 解释：最低花费方式是从 cost[0] 开始，逐个经过那些 1 ，跳过 cost[3] ，一共花费 6 。
 

// 提示：

// cost 的长度范围是 [2, 1000]。
// cost[i] 将会是一个整型数据，范围为 [0, 999] 。

/*
    思路：好吧，确实没思路。
         昨天就在看了，看了官方的题解
         dp[i]=min(dp[i−1]+cost[i−1],dp[i−2]+cost[i−2])
         也没理解，下标和跨楼，画着画着就混乱了。
         今天重拾起来，找了找别的题解。
         dp[i]=Math.min(dp[i-2],dp[i-1])+cost[i]
         嘿，我感觉这个就比较好理解了。但是我又觉得这个有点贪心的感觉。
         1.从前两步直接跳或者从上一步跳一步到目前位置。所以求min(dp[i-2],dp[i-1])
         2.加上当前cost[i]的值。
         3.最后结果再pk一次，是两步前直接跨越还是剩下一步跨越。因此min(dp[len-1],dp[len-2])

*/
var minCostClimbingStairs = function(cost) {
    let dp=[],len=cost.length
    dp[0]=cost[0]
    dp[1]=cost[1]
    for(let i=2;i<len;i++){
        dp[i]=Math.min(dp[i-2],dp[i-1])+cost[i]
    }
    return Math.min(dp[len-1],dp[len-2])
};
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

//优化空间复杂度为O(1)

var minCostClimbingStairs = function(cost) {
    let len=cost.length,cur
    let pre1=cost[1],pre2=cost[0]
    for(let i=2;i<len;i++){
        cur=Math.min(pre2,pre1)+cost[i]
        pre2=pre1
        pre1=cur
    }
    return Math.min(pre1,pre2)
};

