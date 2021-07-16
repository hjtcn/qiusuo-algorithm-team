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

1. Clarification:
         这道题感觉不一样的地方在于，我们首先要把完全平方数求出来，放到一个数组里面，然后接下来就是完全背包问题了
         如，< 12 的完全平方数有
         [1,4,9]: 如果确定这个数是不是完全平方数呢？
         var nums []int
         for i := 1;i <= n;i++ {
             if i * i <= n {
                 nums = append(nums, i * i)
             }
         }
        感觉剩下就和 coin change 零钱兑换一样了
         1. dp数组含义
            dp[j]: 组成和为j的完全平方数的个数
        2. 动态方程
            dp[j] = min(dp[j- nums[i]])
        3. 初始化
            dp[0] = 0
            dp[1-n] = n + 1            
        4. 遍历顺序
coding:

func numSquares(n int) int {
    var nums []int
    dp := make([]int, n + 1)
    for i := 1;i < n+1;i++ {
        dp[i] = n + 1
    }
    for i := 1;i <= n;i++ {
        if i * i <= n {
            nums = append(nums, i * i)
        }
    }
    //fmt.Println(nums)
    dp[0] = 0
    for i := 1;i <=n;i++ {
        for _,num := range nums {
            if num <= i {
                dp[i] = min(dp[i],dp[i - num] + 1)
            }
        }
    }
    if dp[n] > n {
        return -1
    }
    return dp[n]
}

func min(x,y int)int {
    if x < y {
        return x
    }
    return y
}

动态规划方程上面写错了
应该是 dp[j] = min(dp[j - num) + 1)

2. 看题解：
题解里面有很多都是直接算的，我觉得还是先把完全平方数算出来，然后去计算更加清晰一点

3. 复杂度分析：
时间复杂度：O(n * n)
空间复杂度：O(n)

4. 总结：
4.1: 感觉现在写这种题蛮快的，哈哈哈，过了阵痛期，练习还是有效果的
