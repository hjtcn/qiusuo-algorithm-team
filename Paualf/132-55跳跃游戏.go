给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。
 
示例 1：
输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：
输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
 
提示：
• 1 <= nums.length <= 3 * 104
• 0 <= nums[i] <= 105
1. Clearfication:
怎么去思考这类问题？
怎么分解和解决子问题
想到动态规划：
dp[n] 是i可以到达最远下标的位置
dp[n] = i + nums[i], 但是 n 是不是可以到达呢？
取决于dp[n-1] >= 1
如果 dp[n] == len(nums) -1 return true else return false
初始化dp[0] = nums[0]

2. Coding:
下面逻辑不是很正确感觉
func canJump(nums []int) bool {
    dp := make([]int, len(nums) + 1)
    dp[0] = nums[0]
    target := len(nums) - 1
    
    for i := 1;i < len(nums);i++ {
        if dp[i - 1] >= 1 {
            dp[i] =  i + nums[i]
            
            if dp[i] == target {
                return true
            }
        }
    }
    
    return false
}                

3. 看题解：
维护 rightmost 变量
func canJump(nums []int) bool {
    n := len(nums)
    
    rightmost := 0
    
    for i := 0;i < n;i++ {
        if i <= rightmost {
            rightmost = max(rightmost, i + nums[i])
            
            if rightmost >= n - 1 {
                return true
            }
        }
    }
    
    return false  
}
func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
自己少判断了一个循环
func canJump(nums []int) bool {
    if len(nums) <= 1 {
        return true
    }
    
    dp := make([]bool, len(nums))
    dp[0] = true
    
    for i := 1;i < len(nums);i++ {
        for j := i - 1;j >=0;j-- {
            if dp[j] && nums[j] + j >= i {
                dp[i] = true
                break
            }
        }
    }
    
    return dp[len(nums) - 1]
}
缩小目标，都是for循环，感觉不同的人写差好多
func canJump(nums []int)bool {
    end := len(nums) - 1
    
    for i := len(nums) - 2;i > -1;i-- {
        if nums[i] + i >= end {
            end = i
        }
    }
    
    if end == 0 {
        return true
    }
    
    return false
}

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

5. 总结：
5.1: 解决问题的方式和如何解决问题很重要
