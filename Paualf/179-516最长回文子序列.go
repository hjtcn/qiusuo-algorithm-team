给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
示例 1：
输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。
示例 2：
输入：s = "cbbd"
输出：2
解释：一个可能的最长回文子序列为 "bb" 。
 
提示：
1 <= s.length <= 1000
s 仅由小写英文字母组成
    1. Clarification:
        和昨天的题目回文子串蛮像的
       动态规划5部曲：
        1. 确定dp数组（dp table) 以及下标的含义：
            布尔类型的dp[i][j]：表示区间范围[i,j] 注意是左闭右闭 的子串是否是回文串，如果是 dp[i][j] 为 true,否则为 false
        2. 确定递推公司
            整体上 s[i] 与 s[j] 是否相等
            s[i] 与 s[j] 不想等，dp[i][j] = false
            s[i] 与 s[j] 相等：分三种情况：
                1. 下标i与j相同，同一个字符如a
                2. 下标i与j相差1，例如aa
                3. i与j相差大于1的时候，如 cabac, s[i] 与 s[j] 已经相同了
         3. dp数组如何初始化
            dp[i][j] = false
         4. 遍历顺序
            因为用到 dp[i-1][j-1] 所以遍历顺序要从下到上
        5. case by case       
    
    我晕，运行了一下发现，不太行，因为子序列是可以中间删除元素的，子串是连续的所以昨天的代码不能用哈
func longestPalindromeSubseq(s string) int {
    if len(s) <= 0 {
        return 0
    }
    res := 1
    dp:=make([][]bool,len(s))
    for i:=0;i<len(s);i++{
        dp[i]=make([]bool,len(s))
    }
    for i:=len(s)-1;i>=0;i--{
        for j:=i;j<len(s);j++{
            if s[i]==s[j]{
                if j - i <= 1 {
                    dp[i][j]=true
                }else if dp[i+1][j-1]{
                    if j - i + 1 > res {
                        res = j - i + 1
                    }
                    dp[i][j]=true
                }
            }
        }
    }
    return res
}

上面的代码不能用哈
看着题解再分析一波，哈哈哈，说实话，还是有点害怕这种题目，没啥思路

动规五部曲

1. 确定dp数组(dp table) 以及下标的含义：
dp[i][j]：字符串s在[i:j] 范围内最长的回文子序列的长度为dp[i][j]

2. 确定递推公司
如果s[i][j] 相同，那么dp[i][j] = dp[i+1][j-1] + 2
如果 s[i] 与 s[j] 不相同，说明s[i] 和 s[j] 的同时加入并不能增加 [i,j] 区间回文子串的长度，那么分别加入s[i]、s[j] 看看哪一个可以组成最长的回文子序列
加入s[j] 的回文子序列的长度为dp[i+1][j]
加入s[i] 的回文子序列的长度为dp[i][j-1]
if s[i] == s[j] 
   dp[i][j] = dp[i+1][j-1] + 2
else 
   dp[i][j] = max(dp[i+1][j],dp[i][j-1])

3. dp数组如何初始化
初始化也挺不一样的，s[i] == s[j] 的时候
dp[i][j] = 1

4. 确定遍历顺序
dp[i][j] = dp[i+1][j-1] + 2 和 dp[i][j] = max(dp[i+1][j],dp[i][j-1])
 dp[i][j] 是依赖于dp[i+1][j-1] 和dp[i+1][j]
所以遍历i的时候一定要从下到上遍历，这样才能保证，下一行的数据是经过计算的

5. case by case
  c b b d
c 1 1  2 2
b    1 2 2
b       1 1
d          1

func longestPalindromeSubseq(s string) int {
    lenth:=len(s)
    dp:=make([][]int,lenth)
    
    for i:=0;i<lenth;i++{
        for j:=0;j<lenth;j++{
            if dp[i]==nil{
                dp[i]=make([]int,lenth)
            }
            if i==j{
                dp[i][j]=1
            }
        }
    }
    
    for i:=lenth-1;i>=0;i--{
        for j:=i+1;j<lenth;j++{
            if s[i]==s[j]{
                dp[i][j]=dp[i+1][j-1]+2
            }else {
                dp[i][j]=max(dp[i+1][j],dp[i][j-1])
            }
        }
    }
    return dp[0][lenth-1]
}

里面有很多细节：
1. dp数组的初始化
2. dp遍历的时候从下到上，并且边界都要处理好
3. 倒三角形的循环

3. 复杂度分析：
时间复杂度：O(m x n)
空间复杂度：O(m)

4. 总结：
4.1: 字符串在动态规划里面感觉还是比较难的，变化比较多，也是值得练习的，因为细节点比较多，因为真实的工作环境中，这些细节也是挺多的
