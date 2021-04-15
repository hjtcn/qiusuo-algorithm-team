/*
1143. 最长公共子序列
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
*/

/*
    思路：其实一开始是没什么思路的，但拿出了回文子串的耐心，而且是两个字符串，那就先构建一个二维的状态转移方程吧。
        (i,j)   a   b   c  d  e
         a      1   1   1  1  1
         c      1   1   2  2  2    ----->if(s1[i]==s2[j]) dp[i][j]=dp[i-1][j-1]+1
         e      1   1   2  2  3           else    dp[i][j]=Math.max(dp[i-1][j],dp[i][j-1])
    
    感觉字符串相关的动规，每一步之间的关系太难想了。
    这个题，也遇到满多坑的：
    1.刚开始没想着dp[i-1][j-1]，只想到了向下和向右。
    2.内存溢出的控制，i=0或j=0
*/
var longestCommonSubsequence = function (text1, text2) {
    let m = text1.length, n = text2.length
    let dp = Array.from(Array(m), () => Array(n).fill(0))
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            let max = 0
            //下面一溜判断，只是为了防止溢出
            if (i == 0 && j == 0) {
                max = 0
            }
            else if (i == 0) {
                max = dp[i][j - 1]
            }
            else if (j == 0) {
                max = dp[i - 1][j]
            }
            else {
                max = Math.max(dp[i - 1][j], dp[i][j - 1])
            }
            //这个是状态转移的核心
            if (text1[i] == text2[j]) {
                if (i == 0 || j == 0) {
                    //这个位置也做了防止溢出的判断
                    dp[i][j] = 1
                }
                else {
                    dp[i][j] = dp[i - 1][j - 1] + 1
                }

            }
            else {
                dp[i][j] = max
            }
        }
    }
    return dp[m - 1][n - 1]

};

/*
    时间复杂度：O(N*M)
    空间复杂度：O(N*M)
*/


/*
    首先是看完题解的代码优化.节省溢出所做操作
    dp[i][j]由于会有[i-1][j-1],因此事i,j的遍历从1开始。
    此时字符串对比也需要前移一位。s1=text1[i-1],s2=text2[j-1]
    赞赞赞！要学会变通
*/

var longestCommonSubsequence = function (text1, text2) {
    let m = text1.length, n = text2.length
    let dp = Array.from(Array(m+1), () => Array(n+1).fill(0))
    for (let i = 1; i <= m; i++) {
        for (let j = 1; j <=n; j++) {
            let s1=text1[i-1],s2=text2[j-1]
            //这个是状态转移的核心
            if (s1 == s2) {
                dp[i][j] = dp[i - 1][j - 1] + 1
            }
            else {
                dp[i][j] = Math.max(dp[i-1][j],dp[i][j-1])
            }
        }
    }
    return dp[m][n]

};


/*
    降维也是看题解理解的。看着代码理解吧
    如Math.max(dp[j]||0,dp[j-1]).dp[j]代表的是上方的值，dp[j-1]代表左侧的值
     prev则用来记录左上方的值了。
    太巧妙了，dp降维主要就是复用，真的好难理解
    
    
*/
var longestCommonSubsequence = function (text1, text2) {
    let m = text1.length, n = text2.length
    let dp = new Array(m+1).fill(0)
    for (let i = 1; i <= m; i++) {
        let cur=0
        for (let j = 1; j <=n; j++) {
            let prev=cur
            cur=dp[j]
            let s1=text1[i-1],s2=text2[j-1]
            //这个是状态转移的核心
            if (s1 == s2) {
               dp[j]=prev+1
            }
            else {
                dp[j]=Math.max(dp[j]||0,dp[j-1])
            }
        }
    }
    return dp[n]
};
/*
    时间复杂度：O(N*M)
    空间复杂度：O(max(N,M))
*/