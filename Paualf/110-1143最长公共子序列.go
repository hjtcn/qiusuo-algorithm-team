给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
示例 1：
输入：text1 = "abcde", text2 = "ace" 
输出：3  
解释：最长公共子序列是 "ace" ，它的长度为 3 。
示例 2：
输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc" ，它的长度为 3 。
示例 3：
输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0 。
提示：
1 <= text1.length, text2.length <= 1000
text1 和 text2 仅由小写英文字符组成。

1. Clearfication:
   1. 朴素法：
        将text1 的所有子序列放到字符串数组 a 中
        将text2 的所有子序列放到字符串数组 b 中
       
        可能出问题的地方在 怎么构建字符串数组呢？
ret := 0
for _,x := range a {
    for _,y range b {
         if y == x {
            if length(y) > ret {
                 ret = length(y)
             }
        }
    }
}
return ret
 
    2. 动态规划
        dp[i][j]的含义是：
            text2[0->i] 
            text1[0->j] 它们的最长公共子序列
            if text2[i] == text1[j] {
                dp[i][j] = max(max(dp[i - 1][j-1],dp[i - 1][j]),dp[i][j-1]) + 1
            }else {
                  dp[i][j] = max(max(dp[i - 1][j-1],dp[i - 1][j]),dp[i][j-1])
            }
            怎么初始化？
           
            初始化的时候没有想明白，dp[0][0 -> m] 和 dp[0 -> n][0] 是初始化成什么样的值呢？
           
            return dp[m - 1][n - 1]
         a b c d e
      a  1 1 1 1 1
      c  1 1 2 2 2  
      d  1 1 2 2 3  
2. Coding:

3. 看题解：
将字符串问题转化为典型的二维动态规划问题：
假设字符串 text1 和 text2 的长度分别为 m 和 n,创建 m + 1 行 n + 1列的二维数组dp,其中dp[i][j] 表示 text1[0:i] 和 text2[0:j] 的最长公共子序列的长度
text1[0:i] 表示text1的长度为i的前缀
text2[0:j] 表示text2的长度为j的前缀
func longestCommonSubsequence(text1,text2 string) int {
    m,n := len(text1),len(text2)
    
    dp := make([][]int, m + 1)
    
    for i := range dp {
        dp[i] = make([]int, n + 1)
    }
    
    for i,c1 := range text1 {
        for j,c2 := range text2 {
            if c1 == c2 {
                dp[i + 1][j + 1] = dp[i][j] + 1
            }else {
                dp[i + 1][j + 1] = max(dp[i][j+1], dp[i + 1][j])
            }
        }
    }
    
    return dp[m][n]
}
func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
func longestCommonSubsequence(text1 string,text 2 string) int {
    m,n := len(text1),len(text2)
    dp := make([][]int, m + 1)
    
    for i := 0;i <= m;i++ {
        dp[i] = make([]int, n + 1)
    }
    
    for i := 1;i <= m;i++ {
        for j := 1;j <= n;j++ {
            if text1[i-1] = = text2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            }else {
                dp[i][j] = max(dp[i-1][j],dp[i][j-1])
            }
        }
    }
    
    return dp[m][n]
}
func max(x,y int)int {
    if x > y {
        return x
    }
    return y
}

4. 复杂度分析：
时间复杂度：O(m * n)
空间复杂度：O(m * n)

5. 总结：
5.1:如果一个问题你描述不清楚，说明你没有真正理解

5.2: 细节还是蛮重要的，下标变换不对的话，题目就会出错的
