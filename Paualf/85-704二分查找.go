/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
示例 1:
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
 
提示：
你可以假设 nums 中的所有元素是不重复的。
n 将在 [1, 10000]之间。
nums 的每个元素都将在 [-9999, 9999]之间。
*/

1. Clearfication:
升序的整型数组中，查找目标 target
下标是从数组索引  num := len(nums) - 1,索引是从 [0, num - 1] 开始
left := 0,right := len(nums) - 1,搜索区间 [left,right]
每一次执行，mid := left + (right - left) / 2 
for left <= right {
}
每一次求mid，搜索区间就会分为三段区间 [left,mid - 1], mid,[mid + 1,right]

我一直没有搞清楚什么时候写 for left <= right 什么时候写 left < right，我仔细分析了一下
在数组 [2,3] 里面查找目标元素 target
1. 如果初始化定义: left = 0,right = len(nums) - 1 : 1 的话这个时候要写 for left <= right，因为元素都是可以访问到的
2. 如果初始化定义：left = 0,right = len(nums)： 2 的话这个时候要写 for left < right，因为下标2是访问不到的

ex:
如果我们要查找的元素是5，数组元素[2,3]
第一种情况: left <= right 的情况：
left = 0,right = 1,target = 5,left <= right
mid := 0, nums[mid] < target,left++ 
change: left = 1, right = 1, target = 5,left <= right
mid := 1, nums[mid] < target,left++
change:left = 2,right =1, 不满足for循环条件退出 left > right

如果我们要查找的元素是1，数组元素[2,3]
left = 0,right = 1,target = 1,left <= right
mid := 0,nums[mid] > target,right --
change: left = 0,right = 0,target =1,left <= right
mid := 0,nums[mid] > target,right--
change:left = 0,right = -1, 不满足for循环条件退出, left > right
     这样定义是满足条件的,

第一种情况:left < right 的情况：
如果我们定义： [2,3], left = 0,right =1,target = 5,left < right
           mid := 0,nums[mid] < target,left++
change: left = 1,right = 1 不满足 left < right 的判断条件，这个时候会发现下标元素为1的地方没有访问到
如果数组是 [2,5] 的话我们就漏查找元素了
[2,3], left = 0,right = 1,target = 1,left < right
mid := 0,nums[mid] > target,right --
change: left = 0,right = 0, 不满足 left < right 的判断条件，这个时候会发现下标元素为0的地方没有访问到
如果数组是[1,2] 的话我们就漏查找元素了

第二种情况：
在数组 [2,3] 里面查找目标元素 target
 我们定义 left = 0,right = len(nums) : 2,for left < right,target = 5
left = 0,right = 2, left < right
mid = 1,nums[mid] < target ,left ++
change: left = 1,right = 2,left < right
mid = 1, nums[mid] < target,left ++
change: left = 2,right = 2，不满足 left < right 结束
[2,3] left = 0,right = 2,for left < right;target = 1
mid := 1,nums[mid] > target,right --
change,left = 0,right = 1,for left < right 
mid := 0,nums[mid] > target;right --
left = 0,right = 0不满足 left < right 结束
发现下标的定义和 left < right 或者 left <= right 有很大关系。。。

2. Coding：
left,right := 0,len(nums) - 1
for left <= right {
}
func search(nums []int, target int) int {
    left,right := 0,len(nums) - 1
    
    for left <= right {
        mid := left + (right - left) / 2
        
        if nums[mid] == target {
            return mid
        }else if nums[mid] > target {
            right = mid - 1
        }else if nums[mid] < target {
            left = mid + 1
        }
    }
    
    return -1
}

left,right := 0,len(nums)
for left < right {
}
func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    left,right := 0,len(nums)
    
    for left < right {
        mid := left + (right - left) / 2
        
        if nums[mid] == target {
            return mid
        }else if nums[mid] > target {
            right = mid 
        }else if nums[mid] < target {
            left = mid + 1
        }
    }
    
    return -1
}

上面分析了大半天，边界还是分析错了，如果 nums[mid] > target 的时候，执行的是 right = mid，不是 right = mid - 1 。。。。。。

3. 看题解：
func search(nums []int,target int) int {
    return bin_search(0, len(nums) - 1,target, nums)
}
func bin_search(low,high,target int,nums []int) int {
    if low > high {
        return -1
    }
    
    mid := low + (high - low) / 2
    if nums[mid] == target {
        return mid
    }else if nums[mid] > target {
        return bin_search(low,mid - 1,target,nums)
    }else  {
        return bin_search(mid + 1,high, target, nums)
    }
}

4. 复杂度分析：
时间复杂度：O（logN)：每一次分一半，像一颗二叉树一样，它的时间复杂度就是它的树高 O(logN)
空间复杂度：O（1）

5. 总结:
5.1: 二分真的不简单哈,边界条件比较多
