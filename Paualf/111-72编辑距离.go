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

1. Clearfication:
如何分析这种题目呢？脑子里面没有思路

2. Coding:

3. 看题解：
  定义dp[i][j]
dp[i][j] 代表 word1 中前i个字符，变换到word2中前j个字符串，最短需要操作的次数
需要考虑 word1 或 word2 一个字母都没有，即全增加/删除的情况，所以预留dp[0][j] 和 dp[i][0]
状态转移方程
增：dp[i][j] = dp[i][j-1] + 1
删: dp[i][j] = dp[i-1][j] + 1
改：dp[i][j] = dp[i-1][j-1] + 1
func minDistance(word1 string,word2 string) int {
    n1,n2 := len(word1),len(word2)
    
    dp := make([][]int, n1 + 1)
    for i := range dp {
        dp[i] = make([]int,n2 + 1)
    }
    
    for i := 0;i <= n1;i++ {
        dp[i][0] = i
    }
    
    for i := 0;i <= n2;i++ {
        dp[0][i] = i
    }
    
    for i := 1;i <= n1;i++ {
        for j := 1;j <= n2;j++ {
            if word1[i - 1] == word2[j - 1] {
                dp[i][j] = dp[i - 1][j - 1]
            }else {
                dp[i][j] = Min(dp[i-1][j-1] + 1,dp[i-1][j] + 1,dp[i][j-1]+1)
            }
        }
    }
        
    return dp[n1][n2]
}

func Min(args ...int) int {
    min := args[0]
    
    for _,item := range args {
        if item < min {
            min = item
        }
    }
    
    return min
}

memo 保存记录记过的值，从上而下
var _s1, _s2 string
func minDistanceDp(i, j int,memo [][]int) int {
    if i == -1 {
        return j + 1
    }
    if j == -1 {
        return i + 1
    }
    
    // 哈哈哈，超时是少了判断
    if memo[i][j] > 0 {
        return memo[i][j]
    }
    
    if _s1[i] == _s2[j] {
        memo[i][j] = minDistanceDp(i-1, j-1,memo)
    } else {
        memo[i][j] = Min(minDistanceDp(i-1, j-1,memo)+1,
            minDistanceDp(i-1, j,memo)+1,
            minDistanceDp(i, j-1,memo)+1)
    }
    return memo[i][j]
}

func minDistance(s1, s2 string) int {
    _s1 = s1
    _s2 = s2
    m := len(s1)
    n := len(s2)
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
    }
    return minDistanceDp(m - 1, n - 1,memo)
}
func Min(args ...int) int {
    min := args[0]
    for _, item := range args {
        if item < min {
            min = item
        }
    }
    return min
}
使用记忆化数组，提交的时候依然是超时,后来看了下面这个清晰的代码，发现是自己少了判断，哈哈哈
func minDistance(word1 string, word2 string) int {
    l1 := len(word1)
    l2 := len(word2)
    memo := make([][]int, l1)
    for i := range memo {
        memo[i] = make([]int, l2)
    }
    return subMinDist(word1, 0, word2, 0, memo)
}

func subMinDist(word1 string, idx1 int, word2 string, idx2 int, memo [][]int) int {
    if idx1 >= len(word1) {
        return len(word2) - idx2
    }
    
    if idx2 >= len(word2) {
        return len(word1) - idx1
    }
    if memo[idx1][idx2] > 0 {
        return memo[idx1][idx2]
    }
    if word1[idx1] == word2[idx2] {
        r := subMinDist(word1, idx1+1, word2, idx2+1, memo)
        memo[idx1][idx2] = r
        return r
    }
    replace := subMinDist(word1, idx1+1, word2, idx2+1, memo) + 1
    add := subMinDist(word1, idx1, word2, idx2+1, memo) + 1
    del := subMinDist(word1, idx1+1, word2, idx2, memo) + 1
    r := min(replace, min(add, del))
    memo[idx1][idx2] = r
    return r
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
这个题解整体是很舒服的，里面有些判断自己没有看懂

4. 复杂度分析：
时间复杂度：O(mn)
空间复杂度：O(mn)

5. 总结：
5.1:dp[i][j] 的意思还是有点不是很好理解的，如果定义没有定义准确，下面的推导都有可能是错误的
