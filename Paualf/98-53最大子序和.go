给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：
输入：nums = [1]
输出：1
示例 3：
输入：nums = [0]
输出：0
示例 4：
输入：nums = [-1]
输出：-1
示例 5：
输入：nums = [-100000]
输出：-100000
 
提示：
1 <= nums.length <= 3 * 104
-105 <= nums[i] <= 105
 
进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。

1. Clearfication:
第一时间想到的是贪心算法：
定义 ret = 0,sum = 0
让后 for  i : [0,len(nums))
sum = sum + nums[i]
if sum > ret {
ret = sum
}
if sum < 0 {
sum = 0
}
return ret

2. Coding:
func maxSubArray(nums []int) int {
    ret := math.MinInt64
    sum := 0
    for i := 0;i < len(nums);i++ {
        sum = sum + nums[i]
        if sum > ret {
            ret = sum
        }
        if sum < 0 {
            sum = 0
        }
    }
    return ret
}
一开始想的是定义 ret = 0,后来 run case 的时候发现有负数的情况，所以ret需要定义为 math.MinInt64 =  -1 << 63

3. 看题解：
动态规划：
nums 数组的长度是n,下标从0到 n - 1.
定义： f(i) : 以第i个数结尾的连续子数组的最大和，求 max{f(i)} i : [0,n-1],如何求 f(i) 呢？nums[i] 单独拆成一段或者加入 f(i-1)对应的那一段，取决于 nums[i] 和 f(i - 1) + nums[i] 的大小，我们希望获得比较大的，可以写出动态规划转移方程：
f(i) = max(f(i - 1) + nums[i], nums[i])
func maxSubArray(nums []int) int {
    max := nums[0]
    
    for i := 1;i < len(nums);i++ {
        if nums[i] + nums[i - 1] > nums[i] {
            nums[i] += nums[i - 1]
        }
        
        if nums[i] > max {
            max = nums[i]
        }
    }
    
    return max
}
分治：
我们定义一个操作 get(a,l,r) 表示查询a序列 [l,r] 区间内的最大子段和，我们需要求 get(nums,0,nums.size() - 1)。如何分治实现这个操作呢？[l,r] 取 m = (l + r) / 2, 对区间 [l,m] 和 [m + 1, r] 分治求解。当递归逐层深入直到区间长度缩小为1的时候，递归开始回升。
分治的代码和思路没有看的完全明白，还涉及到线段树的相关概念

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1）

5. 总结：
5.1: 定义问题：f(i) : 以第i个数结尾的连续子数组的最大和
5.2: 分解子问题： f(i) = max(f(i - 1) + nums[i], nums[i])
5.3: 动态转移方程
