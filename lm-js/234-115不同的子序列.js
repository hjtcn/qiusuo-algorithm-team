/*
115. 不同的子序列
给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。

字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）

题目数据保证答案符合 32 位带符号整数范围。

 

示例 1：

输入：s = "rabbbit", t = "rabbit"
输出：3
解释：
如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
rabbbit
rabbbit
rabbbit
示例 2：

输入：s = "babgbag", t = "bag"
输出：5
解释：
如下图所示, 有 5 种可以从 s 中得到 "bag" 的方案。 
babgbag
babgbag
babgbag
babgbag
babgbag
 

提示：

0 <= s.length, t.length <= 1000
s 和 t 由英文字母组成

*/

/*
    思路：这道题有点难度
        1.状态转移方程的递推
        2.初始化的思考

        先上五步曲
        1）dp数组的含义
        在没有很明显的可抽象的方案，就暴力假设dp数组的含义为问题指向。
        即dp[i][j]指0~i(s)，0～j(t)中,t在s的子序列中的出现次数
        2)状态转移方程
        if(s[i]==t[i]){
            dp[i][j]=dp[i-1][j]+dp[i-1][j-1]
            //这里少考虑了不使用当前s[i]对比的情况
            就像示例1一样：
            如果rabb,rab对比时,
            s[3]==t[2]，既要知道rab和ra的对比,也要知道rab和rab的对比
        }
        else{
            dp[i][j]=dp[i-1][j]
        }
        3.dp初始化
        dp[i][0],t为'',s出现空''的情况,只能是把所有给删除掉
        即dp[i][0]=1
        dp[0][j],s为空，j非空，即dp[0][j]=0
        dp[0][9],s为空，j为空，dp[0][0]=1
        4.遍历顺序
        从状态方程出发，从上到下，从左到右。
        5.举例
*/
var numDistinct = function(s, t) {
    let len1=s.length,len2=t.length
    if(len1<len2){
        return 0
    }
    let dp=Array.from(Array(len1+1),()=>Array(len2+1).fill(0))
    for(let i=0;i<=len1;i++){
        dp[i][0]=1
    }
    for(let j=1;j<=len2;j++){
        dp[0][j]=0
    }
    //dp[i][j]  0~i,0~j中t出现在s的次数
    for(let i=1;i<=len1;i++){
        for(let j=1;j<=len2;j++){
            if(s[i-1]==t[j-1]){
                dp[i][j]=dp[i-1][j-1]+dp[i-1][j]
            }
            else{
                dp[i][j]=dp[i-1][j]
            }
        }
    }
    return dp[len1][len2]
};

/*
    时间复杂度：O(N^2)
    空间复杂度：O(N^2)
*/

/*
    这个状态转移方程推一半我就想放弃，之前习惯了遍历过程中+1，当时我的状态转移方程一推出来，是真不到咋计算出现个数了。
        if(s[i-1]==t[j-1]){
            dp[i][j]=dp[i-1][j-1]
        }
        else{
            dp[i][j]=dp[i-1][j]
        }
    
    思考：最近这几道题，dp初始化的过程还是很重要的，要多想想，五步曲要走完。

*/