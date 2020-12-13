/*
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
*/

Clearfication:
n == 1 return 1
n == 2 return 2
F(n) = F(n-1) + F(n-2)
F(3) 爬到第3个台阶上，你可能从第1个台阶爬2步，从第2个台阶爬1步，然后大问题就分解为到达第1阶的步数 + 到达第2阶的步数的和
数学归纳法推导可得 F(n) = F(n-1) + F(n-2)

递推公式写出来以后，我们写代码就ok了
func climbStairs(n int) int {
    if n == 1 || n == 2 {
        return n
    }
    return climbStairs(n-1) + climbStairs(n-2)
}

发现提交的时候当 n = 44 的时候，提示超出时间限制
我们以 n = 5 分析一下
n = 5
1. climbStairs(4) + climbStairs(3)
2. climbStaris(4) = climbStaris(3) + climbStairs(2)
3. ClimbStaris(3) = climbStaris(2) + climbStairs(1)

通过这三步分析，我们可以发现里面有重复计算
我们可以使用一个记忆化数组将重复计算的数组缓存下来，如果需要进行重复计算的话，直接读取就可以了,加缓存的话就有写入缓存和读取缓存的步骤
这里是不对的代码，是记忆话数组的时候出错了，传参数的时候没有思考清楚，想要实现的是指针指向的数组元素，这样想有点复杂了，直接传入数组就可以，因为最后结果是想要返回的是 memory的最后一个元素，不是需要在数组中更新元素，同时发现了，自己在Go里面通过指针改变数组元素的操作不是很熟悉

func climbStairs(n int) int {
    ret := make([]int, n)
    helper(n, &ret)
    return ret[n-1]
}
func helper(n int,ret *[]int) {
    if *ret> 0 {
        return ret[n - 1]
    }
    if n == 1 || n == 2 {
        return n
    }
    *ret[n - 1] = helper(n - 1) + helper(n - 2)
    return *ret[n - 1]
}

加了记忆化数组以后：
func climbStairs(n int) int {
    if n <= 3 {
        return n
    }
    memory := make([]int, n + 1)
    return helper(n,memory)
}
func helper(n int, memory []int) int {
    if n <= 3 {
        return n
    }
    if memory[n] > 0 {  
        return memory[n]
    }
    
    memory[n] = helper(n - 1, memory) + helper(n-2, memory)
    return memory[n]
}

动态规划：
dp[n] = dp[n-1]+dp[n-2]
dp[1] = 1
dp[2] = 2
func climbStairs(n int) int {
    if n <= 3 {
        return n 
    }
    dp := make([]int, n + 1)
    dp[1] = 1
    dp[2] = 2
    for i := 3;i <= n;i++ {
        dp[i] = dp[i - 1] + dp[i - 2]
    }
    return dp[n]
}

对动态规划优化，使用两个变量进行循环赋值
func climbStairs(n int) int {
    if n <= 3 {
        return n 
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

总结：
1. 普通递归有重复计算，所以递归的时间复杂度是 2^n

2. 我们可以使用记忆化数组对代码进行优化，也可以使用动态规划对程序进行优化
