/*
474. 一和零
给你一个二进制字符串数组 strs 和两个整数 m 和 n 。

请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。

如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。

 

示例 1：

输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
输出：4
解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。
示例 2：

输入：strs = ["10", "0", "1"], m = 1, n = 1
输出：2
解释：最大的子集是 {"0", "1"} ，所以答案是 2 。
 

提示：

1 <= strs.length <= 600
1 <= strs[i].length <= 100
strs[i] 仅由 '0' 和 '1' 组成
1 <= m, n <= 100

*/

/*
    题解：01背包，将0和1分开
        对于字符串数组来说，每一个字符串都包含不同的01个数
        例如：0001  有三个0，一个1
        到了此时和目标和就有点联系了。
        1.确认dp数组以及下标的含义
        最多多少个字符串包含0的个数不超过m,包含1的个数不超过n，用dp[m][n]表示
        2.确认递推关系
        当前dp[i][j]由前一个字符串推出来
        前一个字符串是dp[i-count0][j-count1]
        因此dp[i][j]=dp[i-count0][j-count1]+1
        遍历过程中去最大值
        dp[i][j]=Max(dp[i][j],dp[i-count0][j-count1]+1)
        3.dp初始化
        let dp=Array.from(Array(m+1),()=>Array(n+1).fill(0))
        4.确认遍历顺序
        1)只要先遍历字符串数组中每个字符串包含的01个数---》{count0,count1}
        2)随后无论先遍历0，还是先遍历1都可以，都属于一个背包。
        5.举例确认dp数组

*/

var findMaxForm = function(strs, m, n) {
    let dfs=(str)=>{
        let count0=0,count1=0
        for(let i=0;i<str.length;i++){
            if(str[i]=='0'){
                count0++;
            }
            else{
                count1++;
            }
        }
        return {count0,count1}
    }
    let dp=Array.from(Array(m+1),()=>Array(n+1).fill(0))
    for(let i=0;i<strs.length;i++){
        let {count0,count1}=dfs(strs[i])
        for(let j=m;j>=count0;j--){
            for(let k=n;k>=count1;k--){
                dp[j][k]=Math.max(dp[j][k],dp[j-count0][k-count1]+1)
            }
        }
    }
    return dp[m][n]
};
/*
    时间复杂度：O(l*m*n)
    空间复杂度：O(m*n)
*/

/*
    思考：当初做的时候特别难，但是记忆也并不是特别深刻。模仿了题解，但并没有总结出模版
*/