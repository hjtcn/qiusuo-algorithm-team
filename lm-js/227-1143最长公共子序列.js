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
    思路：坑点在于边界问题。
        两个字符串比较，还是比较容易构建网格模型的。
        1.dp数组的含义
        dp[i][j]  0~i,0~j两个字符串的最长公共子序列长度
        2.状态转移方程
        if(text1[i]==text2[j])
        {
            dp[i][j]=dp[i-1][j-1]+1
        }
        else{
            dp[i][j]=Math.max(dp[i-1],dp[j-1])
        }
        3.初始化
        dp二维数组为0
        4.遍历顺序，两个公共子序列不管先后，进行遍历
        5.举例

        自己写的时候，没有想出来遍历时从1～len,对比时text1[i-1]==text2[j-1],可以减少很多，边界控制

    题解：降维失败了，记录prev的过程很巧妙，需要开拓思维
        每一行的滚动复制
    (i,j)  0    1   2  3
     0    prev  j  
     1    j-1  dp[j]
     2
     if(text1[i-1]==text2[j-1]){
        dp[j]=prev+1
     }
     else{
        dp[j]=Math.max(dp[j],dp[j-1])
     }

难点：1.首先如何让prev作为i-1,j-1的位置
    在内层j的遍历过程中，
    prev=cur  
    cur=dp[j]
    到外层i开始++，cur更新为0
    从新开始记录prev,其实也就是没更新前的dp[j]
*/

var longestCommonSubsequence = function(text1, text2) {
    let len1=text1.length,len2=text2.length
    let dp=Array.from(Array(len1),()=>Array(len2).fill(0))
    for(let i=0;i<len1;i++){
        for(let j=0;j<len2;j++){
            if(text1[i]==text2[j]){
                if(i!=0&&j!=0){
                     dp[i][j]=dp[i-1][j-1]+1
                }
                else{
                    dp[i][j]=1
                }
            }
            else{
                if(i==0&&j==0){
                dp[i][j]=0
                } 
                else if(i==0){
                    dp[i][j]=dp[i][j-1]
                }
                else if(j==0){
                    dp[i][j]=dp[i-1][j]
                }
                else{ 
                     dp[i][j]=Math.max(dp[i-1][j],dp[i][j-1])
                }
            }
        }
    }
    return dp[len1-1][len2-1]
};

var longestCommonSubsequence = function(text1, text2) {
    let len1=text1.length,len2=text2.length
    let dp=Array.from(Array(len1+1),()=>Array(len2+1).fill(0))
    for(let i=1;i<=len1;i++){
        for(let j=1;j<=len2;j++){
            //要灵活比较
            if(text1[i-1]==text2[j-1]){
                     dp[i][j]=dp[i-1][j-1]+1   
            }
            else{
                     dp[i][j]=Math.max(dp[i-1][j],dp[i][j-1])
            }
        }
    }
    return dp[len1][len2]
};
/*
    时间复杂度：O(len1*len2)
    空间复杂度：O(len1*len2)
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
                dp[j]=Math.max(dp[j]||0,dp[j-1])
            }
        }
    }
    return dp[len2]
}

/*
    时间复杂度：O(len1*len2)
    空间复杂度：O(len2)
*/

/*
    多开拓开拓思维
    1.不会的过程可以构建模版
    2.会的时候优化代码思想。
*/