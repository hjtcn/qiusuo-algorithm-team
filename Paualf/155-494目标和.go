给你一个整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
示例 1：
输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
示例 2：
输入：nums = [1], target = 1
输出：1
 
提示：
1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000

1. Clearfication:
 01 背包：
        什么是01背包问题呢？
        放or不放么
    可以不可以用递归实现？
        如何分解子问题？  

coding:

func findTargetSumWays(nums []int, target int) int {
    ret := 0
    helper(nums,target,0,&ret,0)
    return ret
}

func helper(nums []int,target int,currentSum int,ret *int,index int) {
    // terminator
    if index == len(nums) && currentSum == target {
        *ret = *ret + 1
        return
    }
    if index >= len(nums) {
        return
    }
    
    // process current logic
    helper(nums,target,currentSum + nums[index],ret,index+ 1)
    helper(nums,target,currentSum - nums[index],ret,index+ 1)
    // drill down
}

2. 看题解：
题解里面的回溯是这样写的
func findTargetSumWays(nums []int,target int) (count int) {
    var backtrack func(int, int)
    backtrack = func(index, sum int) {
        if index == len(nums) {
            if sum == target {
                count++
            }
            return
        }
        backtrack(index + 1,sum + num[index])
        backtrack(index + 1,sum - num[index])
    }
    backtrack(0, 0)
    return
}

对比着改一波我的代码
func findTargetSumWays(nums []int, target int) int {
    ret := 0
    helper(nums,target,0,&ret,0)
    return ret
}

func helper(nums []int,target int,currentSum int,ret *int,index int) {
    // terminator
    if index == len(nums) {
        if currentSum == target {
            *ret = *ret + 1
            return
        }
        return
    }
    
    // process current logic
    helper(nums,target,currentSum + nums[index],ret,index+ 1)
    helper(nums,target,currentSum - nums[index],ret,index+ 1)
    // drill down
}

0，1背包问题：
公式：
     （sum-neg）-neg = sum - 2neg = target
neg = (sum - target) / 2
sum - target 为非偶数，若不符合该条件可直接返回0
问题转化成在数组nums中选取若干元素，使得这些元素之和等于neg, 计算选取元素的方案数
定义二维数组dp,其中dp[i][j] 表示在数组nums的前i个数中选取元素，使得这些元素之和等于j的方案数。数组nums的长度为n，则最终答案为dp[n][neg]
边界条件：
dp[0][j],j = 1 => 1
dp[0][j],j>=1, => 0
如果 j < num, 则不能选num，这个时候有 dp[i][j] = dp[i-1][j]
如果 j > num,则如果不选 nums,方案数是 dp[i -1][j], 如果选num,方案数是 dp[i-1][j-num]
这个时候有 dp[i][j] = dp[i-1][j] + dp[i-1][j-num]

func findTargetSumWays(nums []int, target int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    diff := sum - target
    if diff < 0 || diff%2 == 1 {
        return 0
    }
    n, neg := len(nums), diff/2
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, neg+1)
    }
    dp[0][0] = 1
    for i, num := range nums {
        for j := 0; j <= neg; j++ {
            dp[i+1][j] = dp[i][j]
            if j >= num {
                dp[i+1][j] += dp[i][j-num]
            }
        }
    }
    return dp[n][neg]
}

代码中为了处理边界条件将i都+1了，所以看起来和题解中的代码不是很一样

3. 复杂度分析：
时间复杂度：回溯:O(2^n),动态规划：O(n * neg)
间复杂度：O(n),.动态规划：O(n*n)

4. 总结：
这一次看懂了动态规划方程的意思，不过下次还是不一定能写出来，因为推导出动态规划方程转化为代码的时候，需要注意很多细节
