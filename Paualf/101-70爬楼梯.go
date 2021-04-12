假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。
示例 1：
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

1. Clearfication:
分解子问题：
定义：爬到n阶台阶的方法定义为: f(n)
子问题：想要爬到第n阶要从n-1 和n - 2 才能到，所以 f(n) = f(n - 1) + f(n - 2)
初始化：n= 1 => return 1, n = 2 return 2

2. Coding:
func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    return climbStairs(n - 1) + climbStairs(n - 2)
}
超时,里面有重复就算，如 climbStairs(5) : climbStairs(4) + climbStairs(3), climbStairs(6) : climbStairs(5) + climbStairs(4)，里面climbStairs(4) 是重复进行计算的
使用记忆化数组进行存储
func climbStairs(n int) int {
    memo := make([]int, n + 1)
    return helper(n,memo)
}
func helper(n int,memo []int) int {
    if memo[n] > 0 {
        return memo[n]
    }
    
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    
    memo[n] = helper(n - 1,memo) + helper(n - 2,memo)
    
    return memo[n]
}
动态规划：
定义一个数组进行存储
func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    
    dp := make([]int,n + 1)
    dp[1] = 1
    dp[2] = 2
    for i := 3;i <= n;i++ {
        dp[i] = dp[i - 1] + dp[i - 2]
    }
    return dp[n]
}
使用两个变量进行存储：
func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    
    a := 1
    b := 2
    for i := 3;i <= n;i++ {
        tmp := b
        b = a + b
        a = tmp
    }
    return b
}

3. 看题解：
func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    
    a := 1
    b := 2
    for i := 3;i <= n;i++ {
        a,b = b,a+b
    }
    return b
}

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) / O（1）

5. 总结：
5.1：问题定义，分解子问题，初始化，动态方程
