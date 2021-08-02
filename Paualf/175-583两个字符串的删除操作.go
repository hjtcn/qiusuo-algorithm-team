给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。
示例：
输入: "sea", "eat"
输出: 2
解释: 第一步将"sea"变为"ea"，第二步将"eat"变为"ea"
 
提示：
给定单词的长度不超过500。
给定单词中的字符只含有小写字母。

1.Clarfication:
          s e a
       0 1 2 3  
     e 1 2 1 2
     a 2 3 2 1  
     t  3 4 3 2  
      只能删除任意一个字符串中的一个字符
      
      怎么去用程序去表示删除一个字符呢？

2. 看题解：
看了代码随想录的题解：
1. 确定dp数组(dp table) 以及下标的含义:
dp[i][j]: 以i-1为结尾的字符串word1，和以j-1为结尾的字符串word2，所需要删除元素的最少次数

2. 确定递推公式
当word1[i-1] 与 word2[j-1] 相同的时候
当word1[i-1] 与 word2[j-1] 不相同的时候
当word1[i-1]与word2[j-1] 相同的时候，dp[i][j] = dp[i-1][j-1]
当word1[i-1]与word2[j-1] 不相同的时候，有三种情况:
1. 删word1[i-1],最少操作次数为 dp[i-1][j] + 1
2. 删word2[j-1],最少操作次数为dp[i][j-1] + 1
3. 同时删除word1[i-1]和word2[j-1],操作的最小次数为dp[i-1][j-1] + 2
取最小值

3. dp数组如何初始化
dp[i][0] 和 dp[0][j] 是一定要初始化的
dp[i][0]: word2为空字符串，以i-1为结尾的字符串word1要删除多少个元素，才和word2 相同
dp[i][0] = i
同理
dp[0][j] = j

4. 确定遍历顺序
从递推公式：dp[i][j] = min(dp[i-1][j-1] + 2,min(dp[i-1][j], dp[i][j-1])+1),和 dp[i][j] = dp[i-1][j-1] 可以看出dp[i][j]都是根据左上方、正上方、正左方推出来的。

5. case by case

func minDistance(word1 string, word2 string) int {
    m,n := len(word1),len(word2)
    dp := make([][]int, m + 1)
    for i := 0;i < m + 1;i++ {
        dp[i] = make([]int,n + 1)
    }
    for i := 0;i < m + 1;i++ {
        dp[i][0] = i
    }  
    for j := 0;j < n + 1;j++ {
        dp[0][j] = j
    }  
    for i := 1;i < m + 1;i++ {
        for j := 1;j < n + 1;j++ {
            if word1[i - 1] == word2[j - 1] {
                dp[i][j] = dp[i-1][j-1]
            }else {
                dp[i][j] = min(dp[i-1][j-1] + 2,min(dp[i-1][j],dp[i][j-1]) + 1)
            }
        }
    }
   // fmt.Println(dp)
    return dp[m][n] 
}

func min(x,y int)int {
    if x < y {
        return x
    }
    return y
}

3. 复杂度分析：
时间复杂度：O(m * n)
空间复杂度：O(m)

4. 总结：
4.1: 题解里面最长公共子序列也挺巧妙的（没有仔细看）

4.2: 不要贪心，也不要害怕，时间和精力是有限的，搞懂一道题是一道题，看懂一道题解是一道题解，看不懂看看其他的看懂了就行，不要给自己太大压力，平常心
