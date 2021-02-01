// 1137. 第 N 个泰波那契数
// 泰波那契序列 Tn 定义如下： 

// T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2

// 给你整数 n，请返回第 n 个泰波那契数 Tn 的值。

 

// 示例 1：

// 输入：n = 4
// 输出：4
// 解释：
// T_3 = 0 + 1 + 1 = 2
// T_4 = 1 + 1 + 2 = 4
// 示例 2：

// 输入：n = 25
// 输出：1389537
 

// 提示：

// 0 <= n <= 37
// 答案保证是一个 32 位整数，即 answer <= 2^31 - 1。


/* 
思路:和斐波那契一样，Tn=T(n-1)+T(n-2)+T(n-3)
    如果简单使用递归，是会超时的，所以首先我想到的就是数组记录之前的数据，减少重复计算
*/


/*
 * @lc app=leetcode.cn id=1137 lang=javascript
 *
 * [1137] 第 N 个泰波那契数
 */

// @lc code=start
/**
 * @param {number} n
 * @return {number}
 */
//这是自己一开始写的代码
var tribonacci = function(n) {
    let arr=[]
    arr[0]=0,arr[1]=1,arr[2]=1
    for(let i=3;i<=n;i++){
        arr[i]=arr[i-1]+arr[i-2]+arr[i-3]
    }
    return arr[n]
};
// @lc code=end


//优化空间，这个是看题解了解到的思路，解构赋值可以构建出一种滑动窗口的感觉
var tribonacci = function(n) {
    if(n==0) return 0
    if(n==1||n==2) return 1
    let a=0,b=1,c=1
   for(let i=2;i<n;i++){
       [a,b,c]=[b,c,a+b+c]//不断更新数据，c为原三数之和
   }
   return c
};


//这个也是看题解了解到的尾递归，利用参数记录，
//第一个参数是目标数值，递归过程一直减减，直到有目标返回值。
//后面三个参数实现的从0到n递增的过程，最后作为目标值返回
//很巧妙
var tribonacci = function(n) {
    let dfs=(n,t1,t2,t3)=>{
        if(n===0) return t1
        if(n===1) return t2
        if(n===2) return t3
        return dfs(n-1,t2,t3,t1+t2+t3)
    }
    return dfs(n,0,1,1);
}