一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？
示例 1：
输入：m = 3, n = 7
输出：28
示例 2：
输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下
示例 3：
输入：m = 7, n = 3
输出：28
示例 4：
输入：m = 3, n = 3
输出：6
提示：
1 <= m, n <= 100
题目数据保证答案小于等于 2 * 109

1. Clearfication:
  uniquePaths:
        Finish这个位置是由它的左边和它的上面到的，我们是求总共有多少条不同的路径？
                dp[m - 1][n - 1] = dp[m - 2][n-1] + dp[m-1][n - 2]
            定义：dp[m][n] 的含义是到 m 和 n 下标位置的不同路径数
            初始化：dp[0][0 -> n-1] = 1,dp[0 -> m-1][0] = 1
            动态规划方程： dp[m][n] = dp[m-1][n] + dp[m][n - 1]
            
            return dp[m - 1][n - 1]
            返回结果 return dp[m - 1][n - 1]
            伪代码：
            初始化
                dp = make([][]int,m)
                for i := 0;i < m;i++ {
                    dp[i] = make([]int, n)
                }
             for i := 0;i < m;i++ {
                 dp[i][0] = 1
             } 
             for j := 0;j < n;j++ {
                 dp[0][j] = 1
             }
            for i := 1;i < m;i++ {
                for j := 1;j < n;j++ {
                    dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
                }
            }
            return dp[m - 1][n - 1]

2. Coding:
func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        dp[i][0] = 1
    }
    for j := 0; j < n; j++ {
        dp[0][j] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[m-1][n-1]

3. 看题解：
https://leetcode.com/problems/unique-paths/discuss/362375/3-Solutions-in-Go-best-beats-100-timespace-(DP)
这个题解里面是从 递归 -> 记忆化数组->dp->节省空间的dp一路过来的，感觉这样的思路才是比较完整的解题思路，和爬楼梯蛮像的, 分解子问题，然后对子问题进行求解
想了一些递归，确实有些地方不是那么清晰了

4. 复杂度分析：
时间复杂度:O(m*n)
空间复杂度:O(m*n)

5. 总结：
5.1: 这道题目自己做过直接可以有dp的思路，如果没有做过呢？是不是可以使用递归去解决,递归是自己没有想到的

5.2: 慢慢明白了为什么要去看题解，你自己写出来了，你觉得就ok了，如果不去看别人的题解的话，你就不会知道还有更好的思路和想法，也就不会进步
