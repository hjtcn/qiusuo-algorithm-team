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
    思路：这道题踩坑在初始化了。
         1.dp数组的意义
         dp[i][j]单词0～i,0～j需要的最少操作数使得单词相同
         2.状态转移方程
         if(word1[i]==word2[j]){
             dp[i][j]=dp[i-1][j-1]
         }
         else{
             //插入(dp[i][j-1])，删除(dp[i-1][j])，替换(dp[i-1][j-1])三种操作
             dp[i][j]=Math.min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
         }
        最后看题解，明白虽然我一开始方程写对了。但是我没有确定三种操作对应dp数组的含义
         3.初始化
         一开始只把dp数组初始化为0，
         后来写完状态转移方程，并执行测试用例的时候，发现不对。
         开始怀疑状态转移方程是不是写错了。
         回到初始化，为什么要将边界值赋值为位置呢？
 (i,j)   0   1  2  3  4  5  
         1
         2
         3
         4
         5
         从dp数组上思考，dp[i][0]也就是word2的长度为空
         此时word1有多长，就要有多少操作数
         同理，dp[0][j]代表word1长度为空
         此时word2有多长，就要有多少操作数
         4.遍历顺序
         两个单词对比，无所谓顺序，不过题目是说word1转换成word2，先遍历word1，再2会更容易理解一些
         不过由于状态方程的依赖，i,j的遍历需要从小到大，从上到下
         5.举例

    



*/

var minDistance = function(word1, word2) {
    let len1=word1.length,len2=word2.length
    let dp=Array.from(Array(len1+1),()=>Array(len2+1).fill(0));
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
                dp[i][j]=Math.min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1

            }
        }
    }
    return dp[len1][len2]
};

/*
    时间复杂度:O(N^2)
    空间复杂度:O(N^2)
*/

/*
    思考：1.论枚举的重要性
         2.每一天就算面临同一道题，也有不同的疑惑，不同的坑，多经历经历会越来越好的
*/ 