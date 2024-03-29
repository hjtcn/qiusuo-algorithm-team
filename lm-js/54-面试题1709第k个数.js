// 面试题 17.09. 第 k 个数
// 有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。注意，不是必须有这些素因子，而是必须不包含其他的素因子。例如，前几个数按顺序应该是 1，3，5，7，9，15，21。

// 示例 1:

// 输入: k = 5

// 输出: 9


/**
 * @param {number} k
 * @return {number}
 */
var getKthMagicNumber = function (k) {
    let dp = []
    let p3 = 0, p5 = 0, p7 = 0
    dp[0] = 1
    for (let i = 1; i < k; i++) {
        dp[i] = Math.min(dp[p3] * 3, dp[p5] * 5, dp[p7] * 7)
        //更新对应指针
        if (dp[i] == dp[p3] * 3) p3++;
        if (dp[i] == dp[p5] * 5) p5++;
        if (dp[i] == dp[p7] * 7) p7++;
    }
    return dp[k - 1]
};

/*
  讲真，动态规划真是个神奇的算法。看到这道题的时候，竟然想暴力打表
  但是一旦了解动规算法以及拆分题目需求，题目中还有7个隐藏提示呢。
  一旦捋清思路，用的就是个巧劲，但是我太不巧了。
    min                        p3   p5   p7
    []         dp[0]=1         0    0    0
    [3,5,7]    dp[1]=3         1    0    0
    [9,5,7]    dp[2]=5         1    1    0
    [9,15,7]   dp[3]=7         1    1    1
    [9,15,21]  dp[4]=9         2    1    1 
    。。。
    1.首先初始化三个指针
    2.选择三个指针dp[p3] * 3, dp[p5] * 5, dp[p7] * 7中的最小值作为当前最小乘积
    3.所选的最小值对应指针后移。

    时间复杂度：O(N)
    一层for循环
    空间复杂度：O(N)
    借用数组进行数据存储更新
*/