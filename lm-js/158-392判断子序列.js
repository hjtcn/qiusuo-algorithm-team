// 392. 判断子序列
// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

// 进阶：

// 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

// 致谢：

// 特别感谢 @pbrother 添加此问题并且创建所有测试用例。

 

// 示例 1：

// 输入：s = "abc", t = "ahbgdc"
// 输出：true
// 示例 2：

// 输入：s = "axc", t = "ahbgdc"
// 输出：false
 

// 提示：

// 0 <= s.length <= 100
// 0 <= t.length <= 10^4
// 两个字符串都只由小写字符组成。


/*
    思路：暴力对比。1.一瞬间脑子有在想，要不要变成数组再比对。后来想了想还是字符串吧，数组用的比较多了，字符串的截取拼接方法，返回值都不是很熟悉。
        回到正题。
        遍历主字符串。如果当前值t[i]等于s[i],此时i++,如果不等，将t[i]从主字符串中切掉。
        遇到的坑:
        1.果然出在slice()方法。去除当前字符。可以拼接t.slice(0,i)+t.slice(i+1)
        2.子串为空，直接返回true

*/

var isSubsequence = function(s, t) {
    let len=s.length
    if(!len){
        return true
    }
    for(let i=0;i<t.length;){
        if(s[i]==t[i]){
            i++;
            if(i==len){
                return true
            }
        }
        else{   
            t=t.slice(0,i)+t.slice(i+1)
        }
    }
    return false
};

/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/


/*

    我大概看了一眼题解中动态规划的代码，就开始在写与不写中摇摆不定了。时间空间复杂度也没有很优。有一种强行动规的感觉。
    好吧。为了培养动规习惯，，坚持
    1.如果s[0~i]是t[0,j-1]的子序列了，那么s[0,i]肯定是t[j]的子序列。(j代表t字符串指针的移动，i代表s字符串指针的移动)
    2.如果s[0~i-1]是t[0,j-1]的子序列切
*/


var isSubsequence = function(s, t) {
    let sLen=s.length,tLen=t.length
    if(!sLen){
        return true
    }
     //先把dp初始化.js初始化二维数组都不太好整
    let dp = Array.from(Array(sLen+1),()=>Array(tLen+1).fill(false))
    dp[0][0]=true
    dp[0].fill(true)
    for(let i=1;i<sLen+1;i++){
        for(let j=i;j<tLen+1;j++){
            if(dp[i][j-1]||(dp[i-1][j-1]&&s[i-1]==t[j-1])){
                dp[i][j]=true
            }
        }
    }
    return dp[sLen][tLen]
}

/*
    对于这个方案，有两个疑惑：
    1.为什么s[i-1]==t[j-1]而不是s[i]==t[j]
    2.为什么初始化二维数组，横纵长度都加了1。
    对于输入：s = "abc", t = "ahbgdc"，画画状态转移。
    0  1，2，3             (i,j)  0  1，2，3
0   1  1  1 1                0   1  1  1 1    
1   0  0  0 0                1   0  1  0 0  
2   0  0  0 0                2   0  1  0 0
3   0  0  0 0     ---->>>    3   0  1  1 0
4   0  0  0 0                4   0  1  1 1
5   0  0  0 0                5   0  1  1 1
6   0  0  0 0                6   0  1  1 1
*/

/*
    时间复杂度：O(N*M)
    空间复杂度：O(N*M)

*/