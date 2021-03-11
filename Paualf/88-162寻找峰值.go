/*
峰值元素是指其值大于左右相邻值的元素。
给你一个输入数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
你可以假设 nums[-1] = nums[n] = -∞ 。
示例 1：
输入：nums = [1,2,3,1]
输出：2
解释：3 是峰值元素，你的函数应该返回其索引 2。
示例 2：
输入：nums = [1,2,1,3,5,6,4]
输出：1 或 5 
解释：你的函数可以返回索引 1，其峰值元素为 2；
      或者返回索引 5， 其峰值元素为 6。
提示：
1 <= nums.length <= 1000
-231 <= nums[i] <= 231 - 1
对于所有有效的 i 都有 nums[i] != nums[i + 1]
 
进阶：你可以实现时间复杂度为 O(logN) 的解决方案吗？
*/
1. Clearfication:
想到的就是暴力法，暴力法遍历数组节点，> 左边 && > 右边 返回 index
二分的思路没有想到

2. Coding:
里面的边界条件
 nums[-1] = -1
nums[n] = -无穷大，自己还是调了好一会，暴力方法都不是很熟练
 func findPeakElement(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 0
    } 
    for i := 1;i < len(nums);i++ {
        left := nums[i - 1]
        var right int
        if i + 1 < n {
            right = nums[i + 1]
        }else if i + 1 == n{
            right = math.MinInt64
        }
        if nums[i] > left && nums[i] > right {
            return i
        }
    }
    return 0
}

3. 看题解
这个题解自己还是有点没看懂的感觉
func findPeakElement(nums []int) int {
    for i := 0;i < len(nums) - 1;i++ {
        if nums[i] > nums[i + 1] {
            return i
        }
    }
    return len(nums) - 1
}
搜索空间：
func findPeakElement(nums []int) int {
    return search(nums, 0, len(nums) -1)    
}
func search(nums []int,l,r int) {
    if l == r {
        return l
    }
    mid := l + (r - l) / 2
    if nums[mid] > nums[mid + 1] {
        return search(nums,l,mid)
    }
    return search(nums, mid + 1,r)
    
}

4. 复杂度分析：
时间复杂度：线性搜索O(N), 二分查找O(log(N))
空间复杂度：O(1)

5. 总结：
5.1: 搜索空间，感觉写代码真的是一个分析的过程，不是上来就可以写的
