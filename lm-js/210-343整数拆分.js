/*
343. 整数拆分
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1:

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例 2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
说明: 你可以假设 n 不小于 2 且不大于 58。
*/

/*
    思路：递归超时。。只能算到27
    就是不断从1开始减，直到和小于等于1.
    题解：
    1.记忆化数组，首先和我写的递归不同的是，它是可以返回最大积的.
    for循环遍历 
    max=Math.max(res, i * (n - i), i * dfs(n - i))
    最后返回赋值：return memo[n]=max
    2.动态规划：
    for(i=2;i<=n;i++){
        //j<=i-j这个地方也满巧妙。
        for(j=1;j<=i-j;j++){
            dp[i]=Math.max(dp[i],dp[i-j]*j,(i-j)*j)
        }
    }
    
    只想到了dp[i]=dp[i-j]*j。
    (i-j)*j是完全没有意识到

*/
//超时递归
var integerBreak = function(n) {
    let max=0
    let dfs=(product,sum,arr)=>{
        let len=sum
        if(sum<=1){
            max=Math.max(product,max)
            return;
        }
        if(arr.length==0){
            len=sum-1
        }
        for(let i=1;i<=len;i++){
            dfs(product*i,sum-i,[...arr,i])
        }
    }
    dfs(1,n,[])
    return max
};

const integerBreak = (n) => {
    const memo =[];
    const dfs = (n) => {
      if (memo[n]) return memo[n];
      let max = 0;
      for (let i = 1; i <= n - 1; i++) {
        max = Math.max(max, i * (n - i), i * dfs(n - i));
      }
      return memo[n] = max;
    };
    return dfs(n);
};


//动态规划
var integerBreak = function(n) {
    let dp=[]
    dp[1]=1
    dp[2]=1
    for(let i=3;i<=n;i++){
       dp[i]=0
       for(let j=1;j<=i-j;j++){
        dp[i]=Math.max(dp[i],dp[i-j]*j,(i-j)*j)
       }
    }
    return dp[n]
};


/*
    时间复杂度:O(N^2)
    空间复杂度:O(N)
*/

/*
    耐心，慢慢切分小块
*/