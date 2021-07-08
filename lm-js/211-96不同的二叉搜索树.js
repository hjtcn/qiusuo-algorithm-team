/*
96. 不同的二叉搜索树
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。

 

示例 1：


输入：n = 3
输出：5
示例 2：

输入：n = 1
输出：1
 

提示：

1 <= n <= 19

*/

/*
    思路：想复杂了，模拟了一遍全排列，然后在清除非二叉搜索树的时候又不知道准确的条件了。

    题解：看到题解，代码量好小啊。好吧，重要的是思路，而且莫非递归比动规用习惯了？下意识就想递归。
    二叉搜索树不可获取的就是找规律。
    [1,2,3,4,5] 以1为根：则有[2,3,4,5]在树的右边
    [2,3,4,5]以2为根：则有[3,4,5]在树的右边
            若以3为根：则有[2]在树的左边，[4,5]在树的右边
    不断切分
    直到左右各剩一个节点

    目标和不断递加左子树种类*右子树种类
    for(let i=2;i<=n;i++){
        //i-1代表去除根节点
        for(let j=0;j<=i-1;j++){
            dp[i]+=dp[j]*dp[i-1-j]
        }
    }

*/
//自模拟错误代码
var numTrees = function(n) {
    let count=0,map=[]
    let dfs=(arr,sum)=>{
        let len=arr.length
        //此处想寻找非二叉搜索树的规律
        if(len>2&&((arr[len-1]<arr[len-3]&&arr[len-2]>arr[len-3]))){
            return;
        }
        if(sum==0){
            console.log(arr)
            count++;
            return;
        }
         for(let i=1;i<=n;i++){
            if(!map[i]){
                map[i]=1
                dfs([...arr,i],sum-1)
                map[i]=0
            }
         }
    }
    dfs([],n)
    return count
};



var numTrees = function(n) {
    const dp = new Array(n + 1).fill(0)
    dp[0]=1
    dp[1]=1
    // 第一层for循环表示1...n的dp集合
    for(let i=2;i<=n;i++){
      // i-1表示减去根节点
      for(let j=0;j<=i-1;j++){
        dp[i] += dp[j]*dp[i-1-j]
      }
    }
    return dp[n]
  };

/*
    时间复杂度：O(N^2)
    空间复杂度：O(N)
*/

/*
    果然做的题多了，反而胆子大了，想的太多了
    有点巧妙，吓住了
*/