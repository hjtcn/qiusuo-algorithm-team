给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符
 
示例 1：
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：
输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
 
提示：
0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成
    1.Clarification:
      
      动态规划5部曲：
       1.dp定义:
         dp[i][j]: word1[0-i] 与 word2[0-j] 转换的最小步数

       2. 动态规划方程
           if word1[i] == word2[j] {
               dp[i][j] = dp[i-1][j-1]
           }else {
               dp[i][j] = min(dp[i][j-1],dp[i-1][j]) + 1
           }     

        3. 初始化
           dp[0][i] word1 是0的时候，删除掉word2 的长度，所以初始化 
           for i := 0;i < len(word1) + 1;i++ {
               dp[0][i] = i
           }
           同理：
           dp[j][0]
           for j := 0;j < len(word2) + 1;j++ {
               dp[j][0] = j
           }

        4. 遍历顺序
            从小到大

        5. case by case
      h o r s e
    0 1 2 3 4 5
  r 1 1 2 2 3 4
  o 2 2 1 2 3 4
  s 3 2 3 2 2 3

func minDistance(word1 string, word2 string) int {
    m,n := len(word1),len(word2)
    dp := make([][]int,n + 1)
    for i := 0;i < n + 1;i++ {
        dp[i] = make([]int, m + 1)
    }
    // 初始化
    for i := 0;i < m + 1;i++ {
        dp[0][i] = i
    }
    for j := 0;j < n + 1;j++ {
        dp[j][0] = j
    }
    for i := 1;i < n + 1;i++ {
        for j := 1;j < m + 1;j++ {
            if word2[i-1] == word1[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(min(dp[i][j-1],dp[i-1][j]),dp[i-1][j-1]) + 1
            }
        }
    }
    fmt.Println(dp)
    return dp[n][m]
}

func min(x,y int)int {
    if x < y {
        return x
    }
    return y
}

dp 方程分析错了,当 word2[i-1] 和 word1[j-1] 不想等的时候，dp[i][j] 可能从左边dp[i-1][j]，上边dp[i][j-1]，左上dp[i-1][j-1] 过来
还有m,n 的初始化自己也初始化的怪怪的

2. 看题解：
感觉还是抽象和分析子问题的能力
将增删改抽象成数据的变化
将一个大问题分解成若干个小问题，大问题由小问题解决转化而来

3. 复杂度分析：
时间复杂度：O(mxn)
空间复杂度：O(m)

4. 总结：
4.1: 一开始感觉还是懵懵的有点害怕，动起手来以后还好，所以遇到复杂的问题的时候不要害怕，将复杂的问题分解为简单的小问题，将小问题解决

4.2: 联想到工程里面，复杂的逻辑可以由不同的小块代码逻辑组合而来，我们把小块代码逻辑搞懂了，大问题也就慢慢解决了
