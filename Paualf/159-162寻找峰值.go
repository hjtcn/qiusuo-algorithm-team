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
1. Clearfication:
 寻找峰值
        找到到大于左右相邻值的元素
        
        for i := 1 -> len(nums) - 2
        ret := -1
        if nums[i] > nums[i - 1] && nums[i] > nums[i + 1] {
            ret = i
            break
        }
        // 处理最后一个元素的边界条件，最后一个元素大于前一个元素，最后一个元素后面没有元素了
        if ret == -1  && nums[n-1] > nums[n-2]{
            ret = n - 1
        }
       
        return ret

coding:
func findPeakElement(nums []int) int {
    n := len(nums)
    if n <= 0 {
        return -1
    }
    if n == 1 {
        return 0
    }
    if n >= 2 {
        if nums[0] > nums[1] {
            return 0
        }
    }
    ret := -1
    for i := 1;i <= n - 2;i++ {
        if nums[i] > nums[i - 1] && nums[i] > nums[i+1] {
            ret = i
            break
        }
    }
    if ret == -1 && nums[n - 1] > nums[n - 2] {
        ret = n - 1
    }
    return ret
}

写完代码提交的时候，一开始是漏了开头的处理的，漏了这个逻辑
 if n >= 2 {
        if nums[0] > nums[1] {
            return 0
        }
    }
虽然是朴素的解法，还是有一点细节要注意和解决的

2. 看题解
func findPeakElement(nums []int) int {
    for i := 0;i < len(nums) - 1;i++ {
        if nums[i] > nums[i + 1] {
            return i
        }
    }
    return len(nums) - 1
}

感觉是利用数据的特征将规律推出来的

func findPeakElement(nums []int) int {
    return search(nums,0,len(nums) - 1)
}

func search(nums []int,l,r int) int {
    if l == r {
        return l
    }
    mid := l + (r - l) / 2
    if nums[mid] > nums[mid + 1] {
        return search(nums, l ,mid)
    }
    
    return search(nums,mid + 1,r)
}

二分查找使用的还是蛮巧妙的，本质上还是利用了数据的特征

3. 复杂度分析：
时间复杂度：线性扫描O(n), 二分 O(logN)
空间复杂度：O(1) or 二分递归使用递归栈空间O(n)

4. 总结：
4.1:想要把一道题一件事情做好都是需要耐心，需要注意边界条件的
