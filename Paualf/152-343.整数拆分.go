1. Clearfication:
        先找和，然后和的乘积最大
        分解子问题：
            10 = 3 + 3 + 4
            10 = 1 + 9
            10 = 2 + 8
            10 = 3 + 7
            10 = 4 + 6
            10 = 5 + 5
        有多少种组合呢？
        可以把组合列出来么？
        这也是可以出两道题的。。。
代码想了递归，感觉思路不是太对

2. 看题解：
dp[i] 表示将正整数拆分成至少两个正整数和之后，这些正整数的最大乘积。
0，不是正整数，1是最小的正整数，0和1都不能拆分，所以dp[0]=dp[1] = 0
当 i>=2 时，假设对正整数i拆分出的第一个正整数是j(  j >=1 && j < i),则有以下两种方案
将 i 拆分成 j 和 i - j 的 和，且 i - j 不再拆分成多个正整数，此时的乘积是 j x (i - j)
将 i 拆分成j 和 i - j 的和，且 i - j 继续拆分成多个正整数，此时的乘积是 j x dp[i-j]
因此，当j固定时，有dp[i] = max(j x (i - j),j x dp[i-j]),由于j的取值范围是1到i-1,需要遍历所有的j得到dp[i]的最大值，因此可以得到状态转移方程如下：
dp[i] = max(max(j x (i - j), j x dp[i-j])) j >= 1 && j < i

func integerBreak(n int) int {
    dp := make([]int,n + 1)
    
    for i := 2;i <=n;i++ {
        curMax := 0
        for j := 1;j < i;j++ {
            curMax = max(curMax, max(j * (i - j),j * dp[i - j])) 
        }
        dp[i] = curMax
    }
    return dp[n]
}
func max(x,y int) int {
    if x > y {
        return x
    }
    return y
}

3. 复杂度分析：
时间复杂度：O(n*n) 两层for循环
空间复杂度：O(1)

4. 总结：
4.1： 算法是具体的，是清晰的，是每个步骤都是知道它意思的，所以尽可能的把思路写清楚是最好的第一步，第一步不是上去就写代码
