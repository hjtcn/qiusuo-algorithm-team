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
-1000 <= target <= 100
1. Clearfication:
递归计算它的和是否等于target

2. Coding:
// 递归计算它的和是否等于target
func findTargetSumWays(nums []int, target int) int {
    ret := 0
    helper(nums,0,len(nums),target,0,&ret)
    return ret
}
func helper(nums []int,level,max,target,currentSum int,ret *int) {
    // terminator
    if level >= max {
        if currentSum == target {
            *ret = *ret + 1
        }
        return 
    }
    // process current logic
    helper(nums,level+1,max,target,currentSum + nums[level],ret)
    helper(nums,level+1,max,target,currentSum - nums[level],ret)
    // drill down
    // reverse current status
}

3. 看题解：
题解看着有点费解。。。
背包的动态规划
// dp map
func findTargetSumWays(nums []int, S int) int {
    dp := make(map[int]int)
    dp[0 + nums[0]]++
    dp[0 - nums[0]]++
    for i := 1; i < len(nums); i++ {
        next := make(map[int]int)
        for sum := range dp {
            next[sum+nums[i]] += dp[sum]
            next[sum-nums[i]] += dp[sum]
        }
        dp = next
    }
    return dp[S]
}

还有一个转换求和的问题
func findTargetSumWays(nums []int, S int) int {
    sum:=0
    for e:=range nums{
        sum=sum+nums[e]
    }
    // sumAdd + sumCut = sum
    // sumAdd - sumCut = S
    // 2*sumAdd=sum+S
    // sumAdd=(sum+S)/2
    if S>sum{
        return 0
    }
    // 不是2的倍数, sumAdd不可能是小数，所以直接返回
    if (sum+S)&1!=0{
        return 0
    }
    sumAdd:=(sum+S)/2
    // 到这里问题就转换为给定一个数组，有多少种方法装满sumAdd这个数
    dp:=make([]int,sumAdd+1)
    dp[0]=1
    for _,num:=range nums{
        for i:=sumAdd;i>=num;i--{
            dp[i]+=dp[i-num]
        }
    }
    return dp[sumAdd]
}

4. 复杂度分析：
时间复杂度：递归（2^N）,动态规划 O(N*sum)
空间复杂度：递归O(N), 动态规划 O(N*sum)

5. 总结：
5.1: 背包问题的动态规划有点费解，可能也是最近有点忙，没有太多时间去看。。。
