/*
746. 使用最小花费爬楼梯
数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。

每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。

请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。

 

示例 1：

输入：cost = [10, 15, 20]
输出：15
解释：最低花费是从 cost[1] 开始，然后走两步即可到阶梯顶，一共花费 15 。
 示例 2：

输入：cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
输出：6
解释：最低花费方式是从 cost[0] 开始，逐个经过那些 1 ，跳过 cost[3] ，一共花费 6 。
 

提示：

cost 的长度范围是 [2, 1000]。
cost[i] 将会是一个整型数据，范围为 [0, 999] 。

*/

/*
    思路：依旧没有做出来，画了状态转移，但是写出的状态转移方程死活都跟结果不同。
    dp[i]=min(dp[i-1],dp[i-2])+cost[i]

*/

//错误代码
var minCostClimbingStairs = function (cost) {
    let dp = [], len = cost.length
    dp[0] = cost[0]
    dp[1] = Math.min(cost[0],cost[1])//最主要就是这个代码写错了，脑子快了一步，提前对比，后续遍历时的对比就不准确了。
    for (let i = 2; i < len; i++) {
        dp[i] = Math.min(dp[i - 1], dp[i - 2]) + cost[i]
    }
    //第二就是没想到Math.min(dp[len - 1], dp[len - 2])
    return dp[len - 1]
};

// 正确代码
var minCostClimbingStairs = function (cost) {
    let dp = [], len = cost.length
    dp[0] = cost[0]
    dp[1] = cost[1]
    for (let i = 2; i < len; i++) {
        dp[i] = Math.min(dp[i - 1], dp[i - 2]) + cost[i]
    }
    return Math.min(dp[len - 1], dp[len - 2])
};
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

//自己没写出来，只能优化一下空间复杂度吧。
var minCostClimbingStairs = function (cost) {
    let len = cost.length
    let pre2 = cost[0], pre1 = cost[1]
    for (let i = 2; i < len; i++) {
        let cur = Math.min(pre1, pre2) + cost[i]
        pre2 = pre1
        pre1 = cur
    }
    return Math.min(pre1, pre2)
};

/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

/*
    思考：算法题一定要先考虑通用性，然后特殊情况特殊处理。
        先后顺序一旦弄乱，规律性就不好把握了，一边觉着自己思路没错，一边又觉得一些情况对不上，是不是自己思路错了。整个思考过程就会很矛盾。
*/