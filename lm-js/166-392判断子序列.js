/*
392. 判断子序列
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

进阶：

如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

致谢：

特别感谢 @pbrother 添加此问题并且创建所有测试用例。

 

示例 1：

输入：s = "abc", t = "ahbgdc"
输出：true
示例 2：

输入：s = "axc", t = "ahbgdc"
输出：false
 

提示：

0 <= s.length <= 100
0 <= t.length <= 10^4
两个字符串都只由小写字符组成。

*/

/*
思路：方案一：字符串比对，修改父字符串。
    这个还是一开始就能想到的，还是暴力比对。
    如果字符相等，i++,s,t同时向后遍历
    如果不想等，t去除当前字符。
    一旦i==s.length,则返回true
    反之，跳出遍历，则返回false
    
    方案二：(错误思路)原来觉得动态规划的方案并不是最优的，就没看，但是今天自己尝试写一下的时候懵逼了。
    脑子里同时想的递归和动规，导致，dp赋值的时候我想的就是true，false.最后把自己陷入到死循环中。边界的赋值也不知道是该赋值true还是false.
        正确思路：dp[i][j]代表s[0~i],t[0,j]的字符串匹配个数。
        如果s[i]==t[j],则dp[i][j]=dp[i-1][j-1]+1.
        不想等，则dp[i][j]=dp[i][j-1]
    
*/

//方案二的错误代码

var isSubsequence = function(s, t) {
    let len1=s.length,len2=t.length
    //这里来回修改
    let dp=Array.from(Array(len1+1),()=>Array(len2+1).fill(true))
    for(let i=1;i<=len1;i++){
        for(let j=i;j<=len2;j++){
            if(s[i-1]==t[j-1]){
                dp[i][j]=dp[i-1][j-1]
            } 
            else{
                dp[i][j]=dp[i][j-1]
            }
        }
    }
    return dp[len1][len2]
}


//正确代码
var isSubsequence = function(s, t) {
    let len1=s.length
    if(!len1) return true
    for(let i=0;i<t.length;){
        if(s[i]==t[i]){
            i++
        }
        else{
            t=t.slice(0,i)+t.slice(i+1)
        }
        if(i==len1){
            return true
        }
    }
    return false
}
/*
时间复杂度：O(N)
空间复杂度：O(1)
*/

var isSubsequence = function(s, t) {
    let len1=s.length,len2=t.length
    let dp=Array.from(Array(len1+1),()=>Array(len2+1).fill(0))
    for(let i=1;i<=len1;i++){
        for(let j=i;j<=len2;j++){
            if(s[i-1]==t[j-1]){
                dp[i][j]=dp[i-1][j-1]+1
            }
            else{
                dp[i][j]=dp[i][j-1]
            }
        }
    }
    return dp[len1][len2]==len1
}

/*
时间复杂度：O(N*M)
空间复杂度：O(N*M)
*/