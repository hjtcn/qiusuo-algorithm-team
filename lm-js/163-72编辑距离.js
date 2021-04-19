/*
72. 编辑距离
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

*/

/*
    思路：怎么说呢，看完题解感觉自己有点太畏惧这道题了。三种操作不知道该从何解析起。然后我就先把最长公共子序列的模版先卡卡写上了，然后就想把word1的长度去减去最长公共子序列。一开始都正好走反了。
        这个不是从模型开始了，而是从题意分情况解析开始
        1.如果两个单词中字符相等，代表不用操作，不增不删
        if(word1[i]==word2[j])
        dp[i][j]=dp[i-1][j-1]
        2.如果两个单词中字符不相等，可以有三种操作
            1)word1插入一个字符。dp[i][j]=dp[i][j-1]+1
            2)删除一个字符.dp[i][j]=dp[i-1][j]+1
            3)替换一个字符.dp[i][j]=dp[i-1][j-1]+1
        总结if(word[i]!=word2[j]){
            dp[i][j]=min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
        }
    
        递归的过程蛮好理解。我卡在哪里了呢。。。
        初始化刚开始光看代码不太理解，为啥要以下标进行初始化呢。
        画个图感受感受就出来了，哈哈哈哈
                ''   r   o   s
            ''  0    1   2   3
            h   1    1   2   3
            o   2    2   1   2
            r   3    2   2   2
            s   4    3   3   2
            e   5    4   4   3

*/

var minDistance = function(word1, word2) {
    let m=word1.length,n=word2.length
    let dp=Array.from(Array(m+1),()=>Array(n+1).fill(0))
    for(let i=0;i<=m;i++){
        dp[i][0]=i
    }
    for(let j=0;j<=n;j++){
        dp[0][j]=j
    }
    for(let i=1;i<=m;i++){
        for(let j=1;j<=n;j++){
            let s1=word1[i-1],s2=word2[j-1]
            if(s1==s2){
                dp[i][j]=dp[i-1][j-1]
            }
            else{
                dp[i][j]=Math.min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
            }
        }
    }
    return dp[m][n]
};


/*
    时间复杂度：O(M*N)
    空间复杂度：O(M*N)
*/


/*
降维其实对比找规律也是满好写的，但说到底还是盲写。
尤其是当我们记录上一个prev值并且各种复用时，差点崩溃。
这个降维实在太难弄了，甚至比反转链表的各种记录都要难理解。加油吧，现在是生写，多来几次可能就不怕了。
    猜坑：1.dp初始化，word2字符串是目标字符串，要以它进行遍历进行初始化
        2.prev记录。之前直接复制为0就行，现在又要复制dp[0],且dp[0]要不断等于i~m,进行提前记录。
*/
var minDistance = function(word1, word2) {
    let m=word1.length,n=word2.length
    let dp=new Array(m+1)
    for(let i=0;i<=n;i++){
        dp[i]=i
    }
    for(let i=1;i<=m;i++){
        //太过灵活，搞不动啦。。。。。
        let tmp=dp[0]
        dp[0]=i
        for(let j=1;j<=n;j++){
            let prev=tmp
            tmp=dp[j]
            let s1=word1[i-1],s2=word2[j-1]
            if(s1==s2){
                dp[j]=prev
            }
            else{
                dp[j]=Math.min(dp[j-1],dp[j],prev)+1
            }

        }
    }
    return dp[n]
};

/*
    时间复杂度：O(M*N)
    空间复杂度：O(max(M,N))
*/


/*
    思考：1.不要畏惧，我觉得这个题我没有按我前两天的节奏稳步思考，今天看到三个操作的时候，根本就没多想。
         2.初始化，也是为了边界控制，多写写画画会更有感觉一点。
         3.虽然现在降维还比较模糊，但养成思维习惯总会好的。
*/