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
1.Clarification:
   昨天的题目是最多可以完成2笔交易
    今天的限制是 最多可以完成k笔交易
    k：怎么去限制和使用呢？
    对比昨天的进行分析哈：
    1. 确定dp数组以及下标的含义：
        一天有多少个状态呢？
        k 满足情况 没有操作
        k 满足情况，买入 k 值进行变化
        k 满足情况，卖出去
        k 不满足情况，无法买入
        k 不满足情况，无法卖出
        第i天，k满足，买入
        第i天，k满足，卖出
        第i天，k不满足，无法交易
    2. 确定递推公式：
        dp[i][j][k]: 第i天，第j个状态，k的数量 k : 0 -> k-1
    感觉有点难哈。。。
自己还是限在了k上，为什么要三维数组呢？诶诶诶
2. 看题解：
1. 确定 dp 数组以及下标的含义
使用二维数组dp[i][j]:第i天的状态为j，所剩下的最大现金是dp[i][j]
j的状态表示为：
    0: 表示不买入
    1:第一次买入
    2:第一次卖出
    3:第二次买入
    4:第二次卖出
    ......
除了0以外，偶数是卖出，奇数是买入
最多有k笔交易，所以j的范围可以定义为 2 * k  + 1
2. 确定递推公式：

dp[i][1], 表示的是第i天，买入股票的状态，并不是说一定要第i天买入股票
dp[i][1]的状态，有两个具体操作：
操作一：第i天买入股票了，那么dp[i][1] = dp[i-1][0] - prices[i]
操作二：第i天没有操作，沿用前一天的状态,即 dp[i][1] = dp[i-1][1]
选最大的，所以dp[i][1] = max(dp[i-1][0] - prices[i], dp[i-1][1])
同理：dp[i][2] = max(dp[i-1][i] + prices[i], dp[i][2])
j为奇数是买，偶数是卖剩的状态
3. dp初始化
第0天没有操作，就是0，即：dp[0][0] = 0
第0天做第一次买入的操作，dp[0][1] = -prices[0]
第0天做第一次卖出的操作，这个初始值是多少呢？
dp[0][2] = 0
所以：
for j := 1;j < 2 * k;j += 2 {
    dp[0][j] = -prices[0]
}
func maxProfit(k int, prices []int) int {
    n := len(prices)
    if n == 0 {
        return 0
    }
    dp := make([][]int,n)
    for i := 0;i < n;i++ {
        dp[i] = make([]int, 2 * k + 1)
    }
    
    for j := 1;j < 2 * k;j += 2{
        dp[0][j] = -prices[0]
    }
    
    for i := 1;i < n;i++ {
        for j := 0;j < 2 * k - 1;j += 2 {
            dp[i][j+1] = max(dp[i - 1][j + 1], dp[i - 1][j] - prices[i])
            dp[i][j+2] = max(dp[i - 1][j + 2], dp[i - 1][j + 1] + prices[i])
        }
    }
    
    return dp[n-1][2*k]
}

func max(x,y int) int {
    if x > y {
        return x
    }
    return y
}

https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/solution/dai-ma-sui-xiang-lu-188-mai-mai-gu-piao-w01v7/

dp数组表示第i天最多交易k次后手上有(1)或者无(0)股票时的最大利益
因为一次交易买入和卖出至少需要两天，所以有效的限制maxK应该不超过 n/2，超过后就没有约束作用了，相当于maxK为无限，也就是第122题的解法，直接拿过来用
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/solution/go-dong-tai-gui-hua-by-tsubasa-8u-2/

func maxProfit(k int, prices []int) int {
    n := len(prices)
    if n == 0 {
        return 0
    }
    dp := make([][]int,n)
    for i := 0;i < n;i++ {
        dp[i] = make([]int, 2 * k + 1)
    }
    
    for j := 1;j < 2 * k;j += 2{
        dp[0][j] = -prices[0]
    }
    
    for i := 1;i < n;i++ {
        for j := 0;j < 2 * k - 1;j += 2 {
            dp[i][j+1] = max(dp[i - 1][j + 1], dp[i - 1][j] - prices[i])
            dp[i][j+2] = max(dp[i - 1][j + 2], dp[i - 1][j + 1] + prices[i])
        }
    }
    
    return dp[n-1][2*k]
}

func max(x,y int) int {
    if x > y {
        return x
    }
    return y
}

使用二维数组模拟除了所有的状态
也可以定义一个三维数组 dp[i][j][k] : 第i天，第j次买卖，k表示买入还是卖的状态，从定义上看比较直观

func maxProfit(maxK int, prices []int) int {
    n := len(prices)
    if maxK > n / 2 {
        return maxProfitKInf(prices)
    }
    dp := make([][][]int, n)
    for i := 0; i < n; i ++ {
        dp[i] = make([][]int, maxK+1)
        for k := 0; k < maxK+1; k++ {
            dp[i][k] = make([]int, 2)
        }
    }
    for i := 0; i < n; i++ {
        for k := 1; k <= maxK; k++ {
            if i == 0 {
                dp[i][k][0] = 0
                dp[i][k][1] = -prices[i]
                continue
            }
            dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i]);
            dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i]);
        } 
    }
    return dp[n-1][maxK][0]
}

// 第122题的解法
func maxProfitKInf(prices []int) int {
    n := len(prices)
    dp_i_0, dp_i_1 := 0, -1<<31
    for i := 0; i < n; i++ {
        temp := dp_i_0
        dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
        dp_i_1 = max(dp_i_1, temp-prices[i])
    }
    return dp_i_0
}
func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
感觉到了智慧呀

3. 复杂度分析：
时间复杂度：O(n*n)
空间复杂度：O(n)

4. 总结：
4.1: 感觉练习和理论知识都不太够，需要有更多的理论支撑，同时也要动手去做东西，去实践中寻找反馈，避免纸上谈兵
