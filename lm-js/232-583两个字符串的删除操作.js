/*
583. 两个字符串的删除操作
给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。

 

示例：

输入: "sea", "eat"
输出: 2
解释: 第一步将"sea"变为"ea"，第二步将"eat"变为"ea"
 

提示：

给定单词的长度不超过500。
给定单词中的字符只含有小写字母。


*/

/*
    思路：考虑删除的过程没想到，灵机一动想到了，抽象为获取最长相同子序列k，最后让两个单词的长度和-2*k
    1.dp数组的含义
    dp[i][j]单词0～i,0~j的最长相同子序列
    2.状态转移方程
    if(word1[i]==word2[j]){
        dp[i][j]=dp[i-1][j-1]+1
    }
    else{
        dp[i][j]=Math.max(dp[i-1],dp[j-1])
    }
    3.dp初始化
    (1).dp=Array.from(Array(len1+1),()=>Array(len2+1))
    4.遍历顺序
    对比两个字符串，无所谓遍历顺序
    边界值：防止数组溢出
    for循环从1开始
    对比word1[i-1],word2[j-1]
    5.举例


    降维优化动规
    滚动降维
    现在也可以写的比较丝滑了，哈哈哈

题解：删除操作，比较最小值(学会列举可能性)
    1.相等
    不删除
    2.不想等（三种执行方案的最小值）
    1）删除第一个单词的字符
    2）删除第二个单词的字符
    3）都删

    1.dp数组的含义
    dp[i][j]从0～i,0~j的最小删除数
    2.状态转移方程
    if(word1[i]==word[j]){
        dp[i][j]=dp[i-1][j-1]
    }
    else{
        dp[i][j]=Math.min(dp[i-1][j-1]+2,dp[i-1][j]+1,dp[i][j-1]+1)
    }
    3.dp初始化
    如果都没有删除的话，那初始状态就应该是字符所在位置
    距离看起来：
    (i,j) 1 2 3 4 5 6
     1
     2
     3
     4
     5
     6
    for(let i=1;i<=len1;i++){
        dp[i][0]=i
    }
    for(let j=1;j<=len2;j++){
        dp[0][j]=j
    }
    4.遍历过程。(同上)
    5.举例

*/


var minDistance = function(word1, word2) {
    let len1=word1.length,len2=word2.length
    let dp=Array.from(Array(len1+1),()=>Array(len2+1).fill(0))
    for(let i=1;i<=len1;i++){
        for(let j=1;j<=len2;j++){
            if(word1[i-1]==word2[j-1]){
                dp[i][j]=dp[i-1][j-1]+1
            }
            else{
                dp[i][j]=Math.max(dp[i-1][j],dp[i][j-1])
            }
        }
    }
    return len1+len2-2*dp[len1][len2]
};
/*
    时间复杂度：O(N^2)
    空间复杂度：O(N^2)
*/


var minDistance = function(word1, word2) {
    let len1=word1.length,len2=word2.length
    let dp=Array(len2+1).fill(0)
    for(let i=1;i<=len1;i++){
        let cur=0
        for(let j=1;j<=len2;j++){
            let prev=cur
            cur=dp[j]
            if(word1[i-1]==word2[j-1]){
                dp[j]=prev+1
            }
            else{
                dp[j]=Math.max(dp[j],dp[j-1])
            }
        }
    }
    return len1+len2-2*dp[len2]
};
/*
    时间复杂度：O(N^2)
    空间复杂度：O(N)
*/

var minDistance = function(word1, word2) {
    let len1=word1.length,len2=word2.length
    let dp=Array.from(Array(len1+1),()=>Array(len2+1).fill(0))
    for(let i=1;i<=len1;i++){
        dp[i][0]=i
    }
    for(let j=1;j<=len2;j++){
        dp[0][j]=j
    }
    for(let i=1;i<=len1;i++){
        for(let j=1;j<=len2;j++){
            if(word1[i-1]==word2[j-1]){
                dp[i][j]=dp[i-1][j-1]
            }
            else{
                dp[i][j]=Math.min(dp[i-1][j-1]+2,dp[i-1][j]+1,dp[i][j-1]+1)
            }
        }
    }
    return dp[len1][len2]
};

/*
    时间复杂度：O(N^2)
    空间复杂度：O(N)
*/

/*
    思考:需要执行操作时，一定要去思考所有可能性，说不定可以列举呢

*/