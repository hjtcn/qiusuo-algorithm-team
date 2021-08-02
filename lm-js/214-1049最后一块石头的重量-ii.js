/*
1049. 最后一块石头的重量 II
有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。

每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。

 

示例 1：

输入：stones = [2,7,4,1,8,1]
输出：1
解释：
组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。
示例 2：

输入：stones = [31,26,33,21,40]
输出：5
示例 3：

输入：stones = [1,2]
输出：1
 

提示：

1 <= stones.length <= 30
1 <= stones[i] <= 100
 */

/*
    思路：这个题一特别容易往博弈和贪心上想。太难去构建背包的思路了。
    题解：碰撞剩下的石头重量最小？
        想象中，将石头分为两份，是不是两份的重量越接近，碰撞后剩下的石头重量就越小了呢？

        怎么让两份重量接近？一堆石头的总重量sum/2=target,这样target作为背包去放足够多(重量多价值多)的石头。比如一边为target,另一边是sum-target,是不是就是最接近的情况。

        转化为01背包.以target作为背包容量去放足够多的石头。
        
        同时动规五部曲：
        1.dp数组的意义
        dp[j]代表最多可以背j重量的石头
        2.确定递推公式:这里第一个stones[i]代表重量，第二个stones[i]代表价值
        dp[j]=Max(dp[j],dp[j-stones[i]]+stones[i])
        3.dp数组初始化
        let dp=Array(len+1).fill(0)
        4.确认遍历顺序。
          1）先重量，再背包?
          2）先背包，再重量？
        这里都可以，还是按照习惯先遍历石头重量，再遍历背包。
        5.举例推导。

        最后回归题本身，找到最大背包dp[target]了，那么剩下最小的石头是多重？
        sum/2是向下去整的。一份石头是dp[sum/2],另一份石头是sum-dp[sum/2]
        sum-dp[sum/2]>dp[sum/2]
        最后一次碰撞也就是这样===>sum-dp[sum/2]-dp[sum/2]====>sum-2*dp[sum/2]
        我们的目标即为sum-2*dp[sum/2]

*/

var lastStoneWeightII = function(stones) {
    let len=stones.length,sum=0
    for(let i=0;i<len;i++){
       sum+=stones[i]
    }
    let packValue=parseInt(sum/2)
    let dp=Array(packValue+1).fill(0)
    for(let i=0;i<len;i++){
        for(let j=packValue;j>=stones[i];j--){
            dp[j]=Math.max(dp[j],dp[j-stones[i]]+stones[i])
        }
    }
    return sum-2*dp[packValue]
};
/*
    时间复杂度：O(n*sum)
    空间复杂度：O(sum)
*/

/*
    思考：找到最大背包，目前接触的都是sum/2?
    1.是经验问题。2.以大见小的勇气
    一堆石头，选任意两个石头====》变为两堆石头，相互碰撞。
*/