给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
示例 1：
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
示例 2：
输入：grid = [[1,2,3],[4,5,6]]
输出：12
 
提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 100
1. Clearfication:
    m,n 是代表下标，也就是i，j,有助于分析写成了 m 和 n
    打表法：
        dp[m][n] = min(dp[m - 1][n],dp[m][n - 1]) + nums[m][n]
   
    动态规划方程定义：
        dp[m][n] : 到达m,n 位置的最小路径总和
        动态规划方程：
            dp[m][n] = min(dp[m - 1][n], dp[m][n - 1]) + nums[m][n]
       
        初始化：
            for i := 0;i < n;i++ 
                dp[0][i] = nums[0][ii]
            for j := 0;j < m;j==
                dp[j][0] = nums[j][0]
        返回值
            return dp[m - 1][n - 1]

2. Coding:
func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dp := make([][]int, m)
    for i := 0;i < m;i++ {
        dp[i] = make([]int, n)
    }
    for i := 1;i < m;i++ {
         dp[0][i] = grid[0][i] + grid[0][i - 1] 
    }
    for j := 1;j < n;j++ {
        dp[j][0] = grid[j][0] + grid[j - 1][0]
    }
    for i := 1;i < m;i++ {
        for j := 1;j < n;j++ {
            dp[i][j] = min(dp[i - 1][j],dp[i][j - 1]) + grid[i][j] 
        }
    }
    return dp[m - 1][n - 1]
}
func min(x,y int) int {
    if x < y {
        return x
    }
    return y
}
上面的代码输出结果不对，自己排查了一会，没想出来就去打印了 dp数组的结果，发现dp[0][2] 和 dp[2][0] 的值有问题
后来发现自己初始化写的有问题，初始化的时候应该加上 dp[0][i - 1] 自己写成了 加上了 grid[0][i - 1]，说明还是没有分析清楚，所以将问题分析清楚，是代码写正确的很关键的一步哈,调试后的代码
func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dp := make([][]int, m)
    for i := 0;i < m;i++ {
        dp[i] = make([]int, n)
    }
    dp[0][0] = grid[0][0]
    for i := 1;i < n;i++ {
         dp[0][i] = grid[0][i] + dp[0][i - 1] 
    }
    for j := 1;j < m;j++ {
        dp[j][0] = grid[j][0] + dp[j - 1][0]
    }
    for i := 1;i < m;i++ {
        for j := 1;j < n;j++ {
            dp[i][j] = min(dp[i - 1][j],dp[i][j - 1]) + grid[i][j] 
        }
    }
    return dp[m - 1][n - 1]
}
func min(x,y int) int {
    if x < y {
        return x
    }
    return y
}

3. 看题解：
看到官方题解看到
当 i > 0 且j = 0 时，dp[i][0] = dp[i - 1][0] + grid[i][0]
当 i = 0 且 j > 0时，dp[0][j] = dp[0][j-1] + grid[0][j]
当 i > 0 且 j > 0 时，dp[i][j] = min(dp[i - 1][j],dp[i][j - 1]) + grid[i][j]
最后返回 dp[m - 1][n - 1]
感觉这是数学问题哈，如果公式列出来了，代码也就出来了

4. 复杂度分析：
时间复杂度：O(m*n)
空间复杂度：O(m * n)

5. 总结：
5.1:思路是正确的，写出来的代码也不一定正确，因为有边界条件要处理以及可能还有自己没有考虑到的情况

5.2: 找出最优子结构
