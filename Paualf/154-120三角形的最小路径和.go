1. Clearfication:
从上到下：
dp[i][j] 含义是啥呢：从 i:0到当前位置下标和的最大值
dp[0][0] = triangle[0][0]
动态规划方程：
dp[i][j] = max(dp[i-1][j],dp[i][j-1]) + triangle[i][j]
if i - 1 < 0 {
return 0
}
if j - 1 {
return 0
}
return max(dp[len(triangle)-1])
func minimumTotal(triangle [][]int) int {
    m := len(triangle)
    dp := make([][]int,m)
    for i := 0;i < m;i++ {
        dp[i] = make([]int, m)
    }
    dp[0][0] = triangle[0][0]
    for i := 1;i < m;i++ {
        for j := 1;j < len(triangle[i]);j++ {
            var a,b int
            if i - 1 >= 0 {
                 a = dp[i - 1][j]
            }
            if j - 1 >= 0 {
                b = dp[i][j - 1]
            }
            
            dp[i][j] = min(a,b) + triangle[i][j]
        }
    }
    ret := 10001
    for i := 0;i < m;i++ {
        if dp[m-1][i] < ret {
            ret = dp[m-1][i]
        }
    }
    return ret
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

调试后的代码
func minimumTotal(triangle [][]int) int {
    m := len(triangle)
    dp := make([][]int,m)
    for i := 0;i < m;i++ {
        dp[i] = make([]int, m)
    }
    dp[0][0] = triangle[0][0]
    for i := 1;i < m;i++ {
        dp[i][0] = dp[i - 1][0] + triangle[i][0]
        for j := 1;j < len(triangle[i]);j++ { 
            if j == len(triangle[i]) - 1 {
                dp[i][j] = dp[i-1][j-1] + triangle[i][j]
            }else {
                 dp[i][j] = min(dp[i-1][j],dp[i-1][j-1]) + triangle[i][j]
            }  
        }
    }
    ret := 10001
    // for i := 0;i < m;i++ {
    //     for j := 0;j < m;j++ {
    //         fmt.Println(dp[i][j])
    //     }
    //     fmt.Println()
    // }
    for i := 0;i < m;i++ {
        if dp[m-1][i] < ret {
            ret = dp[m-1][i]
        }
    }
    return ret
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

真正难的还是细节：
知道思路以后，对面上面的调试，需要处理很多边界条件，

1. dp[i][0] 如何处理才能更好
2. dp[i][j] 里面是 i-1,j-1,自己一开始写成了 i-1
3. 计算 dp[i][0] = dp[i-1][0] + dp[i][0] 这里应该是 triangle[i][0] 的

所以知道思路真的不一定会做，因为里面还有很多细节需要处理和思考
自下而上怎么做呢？
从倒数第二行开始遍历

dp[i][j]是从下到上经过的最小的和
dp[i][j] = min(dp[i][j],dp[i+1,j+1) + triangle[i][j]
初始化
dp[m-1][0->m-1] = triangle[m][m]

func minimumTotal(triangle [][]int) int {
    m := len(triangle)
    dp := make([][]int, m)
    for i := 0;i < m;i++ {
        dp[i] = make([]int, m)
    }
    // 对最后一行初始化
    for i := 0;i < m;i++ {
        dp[m - 1][i] = triangle[m - 1][i]
    }
    for i := m - 2;i >= 0;i-- {
        for j := 0;j < len(triangle[i]);j++ {
            dp[i][j] = min(dp[i+1][j],dp[i+1][j + 1]) + triangle[i][j]
        }
    }
    return dp[0][0]
}

func min(x, y int)int {
    if x < y {
        return x
    }
    return y
}

调试卡比较久的地方在

for i := 0;i < m;i++ {
    dp[i] = make([]int, m)
}

2. 看题解：
不知道为啥，看题解的时候不喜欢去看节省空间的动态规划，感觉是没必要么，还是觉得难，估计都有的，呀呀呀，慢慢来吧

3. 复杂度分析：
时间复杂度：O(n*n) 遍历 i * j
空间复杂度: O(n*n) dp数组进行存储

4. 总结：
4.1： 思路知道不一定可以做出来的，因为里面有很多细节要处理解决的

4.2: 对比自上而下的动态规划，比较容易想出来，细节多一点，自下而上的方法，不容易想，细节没有那么多

