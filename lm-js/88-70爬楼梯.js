// 70. 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

// 注意：给定 n 是一个正整数。

// 示例 1：

// 输入： 2
// 输出： 2
// 解释： 有两种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶
// 2.  2 阶
// 示例 2：

// 输入： 3
// 输出： 3
// 解释： 有三种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶 + 1 阶
// 2.  1 阶 + 2 阶
// 3.  2 阶 + 1 阶


/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function(n) {
    let res=[]
    res[1]=1
    res[2]=2
    for(let i=3;i<=n;i++){
        res[i]=res[i-1]+res[i-2]
    }
    return res[n]
};

//这题用递归不行，会超时的，可以用递推，以小见大，记录曾经的值。
// 时间复杂度：O(N)
// 一层for循环
// 空间复杂度：O(N)
// 借用数组记录n-1和n-2的值。

//看题解，了解到动规的做法。
// 2=1+1
// 3=2+1
// 5=3+2
// 8=5+3
// 13=8+5
/*
想到 p=q;
    q=r;
    r=q+p;
*/
  
var climbStairs = function(n) {
    let p=0,q=0,r=1
    for(let i=1;i<=n;i++){
        p=q;
        q=r;
        r=p+q
    }
    return r
};

//优化了空间复杂度
// 时间复杂度：O(N)
// 一层for循环
// 空间复杂度：O(1)
// 仅声明了三个变量