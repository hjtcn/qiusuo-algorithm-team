/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
进阶：
你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
 
示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：
输入：nums = [], target = 0
输出：[-1,-1]
 
提示：
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109

*/
1. Clearfication:
暴力法：
从左扫描target [0,len(nums)) 查找 target 的索引下标    
从右扫描target[len(nums)-1, 0] 查找 target 的下标
扫描两遍数组，时间复杂度为O(n)
使用二分查找:
找到target 的开始位置，不断搜索查找空间
查找范围：left,right := 0,len(nums),退出条件 left == right 注意下标越界，查找target 第一个开始的位置，收缩右边界
 for left < right 
mid := (left + right) / 2 
if nums[mid] >= target {
right = mid
}else if nums[mid] < target {
left = mid + 1
}
if left == len(nums) {
return -1
}
return left
查找target 右边第一个结束的位置，收缩左边界
left,right := 0,len(nums)
for left < right {
mid := (left + right) / 2
if nums[mid] <= target {
left = mid + 1
}else if nums[mid] > target {
right = mid
}
}
if left == len(nums) {
return -1
}
     return nums[left] == target ? left : left - 1

2. Coding:
func searchRange(nums []int, target int) []int {
    ret := make([]int, 2)
    ret[0] = -1
    ret[1] = -1
    if len(nums) == 0 {
        return ret
    }
    for i := 0;i < len(nums);i++ {
        if nums[i] == target {
            ret[0] = i
            break
        }
    }
    for i := len(nums) - 1;i >= 0;i-- {
        if nums[i] == target {
            ret[1] = i
            break
        }
    }
    return ret
}
自己尝试写了代码，有报错的
func searchRange(nums []int, target int) []int {
     ret := make([]int, 2)
     ret[0] = -1
     ret[1] = -1
     ret[0] = findLeft(nums,target)
     ret[1] = findRight(nums, target)
     return ret
}
func findLeft(nums []int, target int) int {
    left,right := 0,len(nums)
    for left < right {
        mid := left + (right - left) / 2
        
        if nums[mid] >= target {
            right = mid
        }else if nums[mid] < target {
            left = mid + 1
        }
    }
    if left == len(nums) {
        return -1
    }
    return left
}
func findRight(nums []int, target int) int {
    left,right := 0,len(nums)
    
    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] <= target {
            left = mid + 1
        }else if nums[mid] > target {
            right = mid
        }
    }
    if left == len(nums) {
        return -1
    }
    if nums[left] == target {
        return left
    }
    return left - 1
}

3. 看题解：
排序数组中，查找数组中第一个等于target 的位置：记为 leftIdx 和 第一个 大于等于 target 的位置减1，记为 rightIdx
二分查找中，寻找leftIdx即为在数组中寻找第一个大于等于target 的下标，寻找 rightIdx即为在数组中寻找第一个大于target 的下标，然后将下标减一。两者的判断条件不同
又重新看了一遍 labuladong 写的二分： https://labuladong.gitbook.io/algo/bi-du-wen-zhang/er-fen-cha-zhao-xiang-jie
印象中这是看了第三遍了，看了第三遍自己写的时候才有一丢丢感觉，感觉这道题目还是蛮有价值的
func searchRange(nums []int, target int) []int {
     ret := make([]int, 2)
     ret[0] = -1
     ret[1] = -1
     ret[0] = findLeft(nums,target)
     ret[1] = findRight(nums, target)
     return ret
}
func findLeft(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    left,right := 0,len(nums)
    for left < right {
        mid := left + (right - left) / 2
        
        if nums[mid] == target {
            right = mid
        }else if nums[mid] < target {
            left = mid + 1
        }else if nums[mid] > target {
            right = mid
        }
    }
    if left == len(nums) {
        return -1
    }
    if nums[left] == target {
        return left
    }
    return -1
}
func findRight(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    left,right := 0,len(nums)
    
    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] == target {
            left = mid + 1
        }else if nums[mid] > target {
            right = mid
        }else if nums[mid] < target {
            left = mid + 1
        }
    }
    if left == 0 {
        return -1
    }
    if nums[left - 1] == target {
        return left - 1
    }
    
    return -1
}

4. 复杂度分析：
时间复杂度：暴力搜索：O(n)，二分搜索:O(logN）
空间复杂度：O(1)

5. 总结：
5.1: 有些东西确实是可以量化的，有意识的做一件事情才有可能做成
