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
    思路：上周自己疯狂想要使用贪心计算，各种排序，发现排序的可能性太多。
    题解：动态规划。这道题的dp意义也是第一次学习。
         找到最大子集，也就是包含子集最多。
         dp[j][k]中j,k分别代表0的个数，1的个数。
         首先需要做的，就是遍历数组，获得每个字符串所包含的0的个数和1的个数。
        之后以m,n为开始节点，倒着枚举。（原因：每个字符串只能使用一次，即为有限背包）
        状态转移方程：dp[j][k]=max(dp[j][k],dp[j-count0][k-count1]+1)

*/

var findMaxForm = function(strs, m, n) {
    let len=strs.length
    let getLen01=(str)=>{
        let count0=0,count1=0
        for(let i=0;i<str.length;i++){
            if(str[i]-0==0){
                count0++;
            }
            else{
                count1++;
            }
        }
        return {
            count0,count1
        }
    }
    let dp=Array.from(Array(m+1),()=>Array(n+1).fill(0))
    for(let i=0;i<len;i++){
        let {count0,count1}=getLen01(strs[i])
        for(let j=m;j>=count0;j--){
            for(let k=n;k>=count1;k--){
                dp[j][k]=Math.max(dp[j][k],dp[j-count0][k-count1]+1)
            }
        }
    }
    return dp[m][n]
    
};

/*
    时间复杂度：O(m*n*len)
    空间复杂度：O(m*n)
*/

/*
    这道题典型的01背包问题，只是物品的val需要用两个值代表(0的个数，1的个数)。
    详细题解：https://leetcode-cn.com/problems/ones-and-zeroes/solution/474-yi-he-ling-01bei-bao-xiang-jie-by-ca-s9vr/
    动规五部曲：
    1.确认dp数组及下标意义
    2.状态转移方程
    3.dp数组如何初始化
    4.确定遍历顺序
    5.举例推导dp数组.画出二维图
*/