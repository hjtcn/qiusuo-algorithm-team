/*
279. 完全平方数
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

 

示例 1：

输入：n = 12
输出：3 
解释：12 = 4 + 4 + 4
示例 2：

输入：n = 13
输出：2
解释：13 = 4 + 9
 
提示：

1 <= n <= 104
*/


/*
    思路：刚做完零钱兑换的题，趁热打铁。
        要求返回为n的完全平方数的最小数量
        首先对n去平方根，并且取整为len。
        len*len就是最大的完全平方数
        1.dp数组的意义：
        dp[j]和为j的最少的完全平方个数
        2.状态转移方程
        dp[j]=Math.min(dp[j],dp[j-i*i]+1)
        i是从1到len的有可能的组成完全平方数的根。
        3.dp初始化
        求最小值。数组初始化赋值为Infinity
        同时dp[0]=0,dp[1]=1
        4.遍历顺序
        先和后根或(先根后和)都可
        5.举例推导
        题解：四平方和定理:
        当n=(4^k)*(8m+7),n只能被表示为四个正整数的平方和.
        同时任何一个数都不会被超过4个的平方和组成
    
        因此此题的解决方案：
        1.如果sqrt(n)*sqrt(n)==n
        n即为平方和。。。
        一个判断为平方和isPerfectSquare的函数
        2.n=(4^k)*(8m+7),n为4个数的平方和
        3.n为2个数的平方和。
        for(i=1;i*i<=n;i++){
            j=n-i*i
            if(isPerfectSquare(j)){
                return 2
            }
        }

        4.如果前面三个都不是，就可以直接return 3.


*/

var numSquares = function(n) {
    let len=parseInt(Math.sqrt(n))
    let dp=Array(n+1).fill(Infinity)
    dp[1]=1,dp[0]=0
    for(let i=1;i<=len;i++){
        for(let j=2;j<=n;j++){
            if(j-i*i>=0)
            {
                dp[j]=Math.min(dp[j],dp[j-i*i]+1)
            }
        }
    }
    return dp[n]
};
/*
    时间复杂度：O(N*sqrt(N))
    空间复杂度：O(N)
*/


var numSquares = function(n) {
    
    let isPerfectSquare=(n)=>{
        let k=parseInt(Math.sqrt(n))
        return k*k===n
    }
    let checkoutAnswer4=(n)=>{
        //这个地方的代码也是有点内容的
        //如何判断n=(4^k)*(8m+7)
        while(n%4==0){
            n/=4
        }
        return n%8==7
    }
    if(isPerfectSquare(n)){
        return 1
    }
    if(checkoutAnswer4(n)){
        return 4
    }
    for(let i=1;i<=n;i++){
        let j=n-i*i
        if(isPerfectSquare(j)){
            return 2
        }
    }
    return 3

}
/*
    时间复杂度：O(sqrt(N))
    空间复杂度：O(N)
*/

/*
    思考：跟平方有关的，就是要敢拆，才有无数种可能。
*/