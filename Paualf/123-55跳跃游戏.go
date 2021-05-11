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
怎么保证连续的跳跃判断呢？

2. Coding:

3. 看题解：
去更新可以跳到的最长距离
func canJump(nums []int) bool {
    n := len(nums)
    
    rightmost := 0
    
    for i := 0;i < n;i++ {
        if i <= rightmost {
            rightmost = max(rightmost,i + nums[i])
            if rightmost >= n - 1 {
                return true
            }
        }
    }
    
    return false 
}

func max(x, y int)int {
    if x > y {
        return x
    }
    return y
}

动态规划：还是感觉蛮巧妙的
https://leetcode-cn.com/problems/jump-game/solution/55-tiao-yue-you-xi-golang-dong-tai-gui-hua-shi-xia/
func canJump(nums []int) bool {
    if len(nums) <= 1 {
        return true
    }
    
    dp := make([]bool, len(nums))
    dp[0] = true
    
    for i := 1;i < len(nums);i++ {
        for j := i - 1;j >= 0;j-- {
            if dp[j] && nums[j] >= i - j {
                dp[i] = true
                break
            }
        }
    }
    
    return dp[len(nums) - 1 ]
}

4. 复杂度分析：
时间复杂度：贪心：O(n), 动态规划：O(n*n)
空间复杂度：贪心:O(1),动态规划：O(n)

5. 总结：
5.1： 贪心的有些思路还是蛮巧妙的，想不到，可能是自己写的代码太少了
