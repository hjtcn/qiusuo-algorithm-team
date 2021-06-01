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
1. Cleatfiaction
  partition 操作怎么分析和解决呢？
    联想到快速排序
    sort 以后 双指针进行遍历
    [1,5,5,11]
     i = 0,j = n - 1
     leftSum = nums[0]
     rightSum = nums[n - 1]
       
        for i < j {
            if leftSum < rightSum && i + 1 < j {
                leftSum += nums[i + 1]
                i++
            } else {
                if j - 1 > i {
                    rightSum += nums[j-1]
                    j--
                }
            }
        }
    return leftSum == rightSum

2. Coding:
func canPartition(nums []int) bool {
    sort.Slice(nums,func(i ,j int)bool{
        return nums[i] < nums[j]   
    })
    n := len(nums)
      i,j := 0,n - 1
     leftSum := nums[0]
     rightSum := nums[n - 1]
        
        for i < j {
            if leftSum == rightSum {
                 i++
                 continue    
            }
            if leftSum < rightSum && i + 1 < j {
                leftSum += nums[i + 1]
                i++
            } else {
                if j - 1 > i {
                    rightSum += nums[j-1]
                }
                 j--
            }
        }
    return leftSum == rightSum
}
看到[2,2,1,1] 这个case以后发现自己的这个思路是错的

3. 看题解：
func canPartition(nums []int) bool {
    n := len(nums)
    
    if n < 2 {
        return false
    }
    
    sum,max := 0,0
    for _,v := range nums {
        sum += v 
        if v > max {
            max = v
        }
    }
    
    if sum % 2 != 0 {
        return false
    }
    target := sum / 2
    if max > target {
        return false
    }
    
    dp := make([][]bool, n)
    
    for i := range dp {
        dp[i] = make([]bool, target + 1)
    }
    
    for i := 0;i < n;i++ {
        dp[i][0] = true
    }
    
    dp[0][nums[0]] = true
    for i := 1;i < n;i++ {
        v := nums[i]
        for j := 1;j <= target;j++ {
            if j >= v {
                dp[i][j] = dp[i - 1][j] || dp[i - 1][j - v]
            }else {
                dp[i][j] = dp[i - 1][j]
            }
        }
    }
    
    return dp[n-1][target]
}
感觉难的地方是转换问题的地方
dp[i][j]: 表示从数组的[0,i]下标范围选取若干个正整数（可以是0个），是否存在一种选取方案使得被选取的正整数的和等于j。
看到题解，如果找一半的话我们是不是就可以用dfs搜索去解决了，如果超时的话，就可以用memo数组去减少时间复杂度，再优化就是动态规划了

4. 复杂度分析：
时间复杂度：O(n*target)
空间复杂度：O(n)

5. 总结：
5.1: 问题的转化和分析是比较关键的
