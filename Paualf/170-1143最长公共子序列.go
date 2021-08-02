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

    1. Clarification:
         a b c d e
       0 0 0 0 0 0           
     a 0 1 1 1 1 1 
     c 0   
     e 0

        动态规划五部曲：
            1. dp含义
             dp[i][j] 代表 text1[0...i] text2[0...j]最长公共子序列

            2.动态规划方程
               if text1[i] == text1[j] {
                   dp[i][j] = dp[i-1][j-1] + 1
               }else {
                   dp[i][j] = max(dp[i-1][j],dp[i][j-1])
               }

           3. 初始化
             num1,num2 := len(text1),len(text2)
             dp := make([][]int, num1 + 1)
            for i := 0;i < num1 + 1;i ++ {
                dp[i] = make([]int, num2 + 1)
            }   

            4. 遍历顺序
                从小到大即可

            5. 构造case

coding:
func longestCommonSubsequence(text1 string, text2 string) int {
    num1,num2 := len(text1),len(text2)
    if num1 == 0 || num2 == 0 {
        return 0
    }
    dp := make([][]int, num1 + 1)
    for i := 0;i < num1 + 1;i++ {
        dp[i] = make([]int, num2 + 1)
    }
    for i := 1;i < num1 + 1;i++ {
        for j := 1;j < num2 + 1;j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i - 1][j - 1] + 1
            }else {
                dp[i][j] = max(dp[i-1][j],dp[i][j-1])
            }
        }
    }
    
    return dp[num1][num2]
}

func max(x,y int)int {
    if x > y {
        return x
    }
    return y
}

细节点在于初始化和for循环的地方
i := 1;i < num1 + 1;i++
     j := 1;j < num2 + 1;j++
比较的地方在于 text1[i-1] == text2[j-1]

2. 看题解：

3. 复杂度分析：
时间复杂度：O(m x n)
空间复杂度：O(n）

4. 总结：
4.1：还是要分析清楚的，慢慢的明白了提前把题目分析清楚的意义，提前把问题分析清楚，在你解决问题和分析问题的时候会给你很大的帮助，特别是有很多细节case的地方
