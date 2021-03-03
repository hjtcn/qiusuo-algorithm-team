/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] 。
请找出其中最小的元素。
示例 1：
输入：nums = [3,4,5,1,2]
输出：1
示例 2：
输入：nums = [4,5,6,7,0,1,2]
输出：0
示例 3：
输入：nums = [1]
输出：1
 
提示：
1 <= nums.length <= 5000
-5000 <= nums[i] <= 5000
nums 中的所有整数都是 唯一 的
nums 原来是一个升序排序的数组，但在预先未知的某个点上进行了旋转
*/

1. Clearfication:
分析数据特征：原有数组是升序的，在某个点上进行了旋转，我们找到中间的 mid 节点，可以判断 [left, mid] 和 [mid+1, right] 那一段是递增的，初始化min := math.MaxInt 为最大值，然后比较递增的那部分序列的最小值是否小于 min,如果小于 min 的话，则更新min的值，然后更新搜索区间
left,right := 0,len(nums) -1 
minNum := math.MaxInt
for left <= right {
mid := left + (right - left) / 2
if  nums[mid] >= nums[left] {
if nums[left] < minNum {
minNum = nums[left]
}
left = mid + 1
}else {
if nums[mid] < minNum {
minNum = nums[mid]
}
right = mid - 1
}
} 
return minNum
没有重复元素的情况下这样处理是可以的，
如果有重复元素的情况下，你可以分析清楚么？
如 case [10,1,10,10,10], 在这种情况下，上面的代码就不适用了，那么这种情况下应该怎么分析和处理呢？我目前没有想到哈

2. Coding:
没有重复元素的以后收缩查找空间
func findMin(nums []int) int {
    left,right := 0,len(nums) - 1
    
    minNum := math.MaxInt32
    
    for left <= right {
        mid := left + (right - left) / 2
        
        if nums[mid] >= nums[left] {
            if nums[left] < minNum {
                minNum = nums[left]
            }
            left = mid + 1
        }else {
            if nums[mid] < minNum {
                minNum = nums[mid]
            }
            right = mid -1
        }
    }
    
    return minNum
}

有重复元素的情况下，自己在上面的思路下面没有想到好的处理方法
然后就试了一下二分，发现试了一下，过了...
func findMin(nums []int) int {
    minNum := math.MaxInt32
    for i := 0;i < len(nums);i++ {
        if nums[i] < minNum {
            minNum = nums[i]
        }
    }
    return minNum
}
3. 看题解:
func finMin(nums []int) int {
    left,right := 0,len(nums) - 1 
    
    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] < nums[right] {
            right = mid
        }else if nums[mid] > nums[right] {
            left = mid + 1
        }else {
            right--
        }
    }
    
    return nums[left]
}

4. 复杂度分析:
时间复杂度：O(logN) 二分查找，每次收缩一半空间
空间复杂度：O(1)

5. 总结：
5.1 细节是魔鬼， right 怎么定义,循环判断怎么退出，判断条件怎么收缩空间，最后返回值如何处理，一个不小心就凉凉
