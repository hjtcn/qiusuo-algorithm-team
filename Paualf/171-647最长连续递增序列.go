给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。
示例 1：
输入：nums = [1,3,5,4,7]
输出：3
解释：最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。 
示例 2：
输入：nums = [2,2,2,2,2]
输出：1
解释：最长连续递增序列是 [2], 长度为1。
 
提示：
1 <= nums.length <= 104
-109 <= nums[i] <= 109

1. Clarification:
        每个点都有自己的子序列长度，返回最长的子序列长度
        动态规划5部曲：
            1. dp含义：
                dp[i]: 下标为i的子序列长度
            2. 动态规划方程：
                if nums[i] > nums[i-1] 
                    dp[i] = dp[i-1] + 1
                else 
                    dp[i] = 1
            3. 初始化
                dp[0] =  1
            4. 遍历顺序
                从低到高
            5.构造case
                nums: [1,3,5,4,7]
                dp :  [1,2,3,1,2]
                返回max(dp) = 3

coding:
func findLengthOfLCIS(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    
    dp[0] = 1
    for i := 1;i < n;i++ {
        if nums[i] > nums[i-1] {
            dp[i] = dp[i-1] + 1
        }else {
            dp[i] = 1
        }
    }
    ret := 0
    for i := 0;i < n;i++ {
        if dp[i] > ret {
            ret = dp[i]
        }
    }
    return ret
}

2. 看题解：

贪心:
func findLengthOfLCIS(nums []int) (ans int) {
    start := 0
    for i, v := range nums {
        if i > 0 && v <= nums[i-1] {
            start = i
        }
        ans = max(ans, i-start+1)
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
这样的贪心好像更简单一点

class Solution {
public:
    int findLengthOfLCIS(vector<int>& nums) {
        if (nums.size() == 0) return 0;
        int result = 1; // 连续子序列最少也是1
        int count = 1;
        for (int i = 0; i < nums.size() - 1; i++) {
            if (nums[i + 1] > nums[i]) { // 连续记录
                count++;
            } else { // 不连续，count从头开始
                count = 1;
            }
            if (count > result) result = count;
        }
        return result;
    }
};

3.复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) or O(1)

4. 总结：
4.1: 慢慢觉得，动态规划很考验对问题的分析和理解能力，对平常做需求也是会有帮助的
