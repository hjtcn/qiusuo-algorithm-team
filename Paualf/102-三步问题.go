三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。
示例1:
 输入：n = 3 
 输出：4
 说明: 有四种走法
示例2:
 输入：n = 5
 输出：13
提示:
n范围在[1, 1000000]之间

1. Clearfication:
 f(n) : 输入n,然后进行求解到达n有多少种走法
 f(n) = f(n - 1) + f(n - 2) + f(n - 3)
 n = 0 -> 1
 n = 1 -> 1
 n = 2 -> 2
感觉和爬楼梯的题目挺相似的

2. Coding:
func waysToStep(n int) int {
    if n <= 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    dp := make([]int, n + 1)
    dp[0] = 1
    dp[1] = 1
    dp[2] = 2
    for i := 3;i <= n;i++ {
        dp[i] = dp[i - 1] + dp[i - 2] + dp[i - 3] 
    }
    return dp[n] % 1000000007
}
当输入为76 的时候输出为 847277573 预期是 176653584
没有思路去想那里出问题了

3. 看题解:
func waysToStep(n int) int {
    if n <= 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    dp := make([]int, n + 1)
    dp[0] = 1
    dp[1] = 1
    dp[2] = 2
    for i := 3;i <= n;i++ {
        dp[i] = (dp[i - 1] + dp[i - 2] + dp[i - 3]) % 1000000007
    }
    return dp[n] 
}
想着是和昨天的题目是类似一样的，怎么通过不了呢，是自己去余的地方位置放错了
func waysToStep(n int) int {
    if n <= 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
 
     a := 1
     b := 1
     c := 2
    for i := 3;i <= n;i++ {
        a,b,c = b,c,(a+b+c) % 1000000007
    }
    return c
}

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) / O(1)

5. 总结：
5.1: 思路会了，将它写出来还是要仔细分析的，一不小心就会漏掉条件的
