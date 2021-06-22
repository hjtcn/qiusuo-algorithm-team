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
 

提示：

1 <= text1.length, text2.length <= 1000
text1 和 text2 仅由小写英文字符组成。

*/

/*
    思路：动态规划。
        进步可能就是看到两个字符串的比较，以及要统计公共子序列长度，可以构建状态转移方程
        a b c d e
    a   1 1 1 1 1 
    c   1 1 2 2 2
    e   1 1 2 2 3
    如果text1[i]==text2[j]
    则dp[i][j]=dp[i-1][j-1]+1
    反之，
    dp[i][j]=Max(dp[i-1][j],dp[i][j-1])

    动态规划重点：1.防止数组溢出。索引从1开始遍历
                2.二维数组初始化

    题解：降维。根据昨天做的题可以更加确定
    dp[i-1][j]为更新前的dp[j]
    dp[i][j-1]为dp[j-1]
    因此，难点就在于dp[i-1][j-1]怎么获取,如何提前记录。
        a b c d e
    a   1 1 1 1 1 
    c   1 1 2 2 2
    e   1 1 2 2 3
    for(let i=0;i<len1;i++){
        let cur=0
        for(let j=0;j<len2;j++){
            //提前记录的过程
           let prev=cur
           cur=dp[j]
           //更新过程
        }
    }
    总结：dp[i-1][j-1]为dp[i-1][j]的前一位
    当前轮遍历：
         prev为未更新前dp[j]的前一位。

    

*/

var longestCommonSubsequence = function(text1, text2) {
    let len1=text1.length,len2=text2.length
    let dp=Array.from(Array(len1+1),()=>Array(len2+1).fill(0))
    for(let i=1;i<=len1;i++){
        for(let j=1;j<=len2;j++){
         
            if(text1[i-1]==text2[j-1]){
                dp[i][j]=dp[i-1][j-1]+1
            }
            else{
                dp[i][j]=Math.max(dp[i-1][j],dp[i][j-1])
            }
        }
    }
    return dp[len1][len2]
}
/*
    时间复杂度：O(M*N)
    空间复杂度：O(M*N)
*/


var longestCommonSubsequence = function(text1, text2) {
    let len1=text1.length,len2=text2.length
    let dp=Array(len2).fill(0)
    for(let i=1;i<=len1;i++){
        let cur=0
        for(let j=1;j<=len2;j++){
            let prev=cur
            cur=dp[j]
            if(text1[i-1]==text2[j-1]){
                dp[j]=prev+1
            }
            else{
                //这里dp[j]||0防止出现NaN
                dp[j]=Math.max(dp[j]||0,dp[j-1])
            }
        }
    }
    return dp[len2]
}

/*
    时间复杂度：O(M*N)
    空间复杂度：O(N)
*/

/*
    慢慢练习，总结自己的一套模版
*/