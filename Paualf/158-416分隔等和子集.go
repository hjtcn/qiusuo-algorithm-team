给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
 
示例 1：
输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。
示例 2：
输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。
 
提示：
• 1 <= nums.length <= 200
• 1 <= nums[i] <= 100

1. Clearfication:
     求和： 如果不是偶数则 return false
        如果是偶数，target = sum / 2
        我们需要找能不能有元素和等于 target 的
        0，1 背包问题，背包容量为 target, 看能不能有元素填满这个背包
        定义
        dp[i][j]:
            截止到数组下标i是否存在满足和为j的数，满足为true,不满足为false
            if nums[i] > j {
                dp[i][j] = dp[i - 1][j]
            }else {
                dp[i][j] = dp[i - 1][j] || dp[i -1][j-nums[i]]
            }
        return dp[n][target]

Coding:

func canPartition(nums []int) bool {
    n := len(nums)
    sum := 0
    for i := 0;i < n;i++ {
        sum += nums[i]
    }   
    if sum / 2 == 1 {
        return false
    }
    target := sum / 2
    dp := make([][]bool, n + 1)
    for i := 0;i < n ;i++ {
        dp[i] = make([]bool,target + 1)
    }
    dp[0][0] = true
    for i := 1;i < n;i++ {
        for j := 1;j < target + 1;j++ {
            if nums[i] > j {
                dp[i][j] = dp[i - 1][j]
            }else {
                dp[i][j] = dp[i - 1][j] || dp[i-1][j-nums[i]]
            }
        }
    }
    for i := 0;i < n;i++ {
        for j := 0;j < target + 1;j++ {
            fmt.Println(dp[i][j])
        }
    }
    return dp[n - 1][target]
}
这次是思路写出来了，调试没有通过。。。
func canPartition(nums []int) bool {
    n := len(nums)
    if n < 2 {
        return false
    }
    sum, max := 0, 0
    for _, v := range nums {
        sum += v
        if v > max {
            max = v
        }
    }
    if sum%2 != 0 {
        return false
    }
    target := sum / 2
    if max > target {
        return false
    }
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, target+1)
    }
    for i := 0; i < n; i++ {
        dp[i][0] = true
    }
    dp[0][nums[0]] = true
    for i := 1; i < n; i++ {
        v := nums[i]
        for j := 1; j <= target; j++ {
            if j >= v {
                dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }
    return dp[n-1][target]
}

看了下题解，漏判断，当一个元素 >= sum/2的时候应该return fasle，还有如果不是偶数 return false
初始化的时候也初始化错了

func canPartition(nums []int) bool {
    n := len(nums)
    if n < 2 {
        return false
    }
    sum, max := 0, 0
    for _, v := range nums {
        sum += v
        if v > max {
            max = v
        }
    }
    if sum%2 != 0 {
        return false
    }
    target := sum / 2
    if max > target {
        return false
    }
    dp := make([]bool, target+1)
    dp[0] = true
    for i := 0; i < n; i++ {
        v := nums[i]
        for j := target; j >= v; j-- {
            dp[j] = dp[j] || dp[j-v]
        }
    }
    return dp[target]
}

使用一维数组缩小使用空间的时候，要注意需要从大到小进行计算，如果从小到大更新dp值，那么在计算dp[j] 的值的时候，dp[j-nums[i]]已经是被更新过的状态，不再是上一行的dp值了

3. 复杂度分析：
时间复杂度：O(n*m)
空间复杂度：O(n*m)

4. 总结：
4.1: 需要注意判断数量是否是偶数，还需要判断数组中最大的值是否已经>= sum /2

4.2: dp初始化的时候要想清楚，这一环，是dp方程推导出来以后，题目是否能正常运行的比较重要的一部分
