//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 请必须使用时间复杂度为 O(log n) 的算法。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,3,5,6], target = 5
//输出: 2
//
//
// 示例 2:
//
//
//输入: nums = [1,3,5,6], target = 2
//输出: 1
//
//
// 示例 3:
//
//
//输入: nums = [1,3,5,6], target = 7
//输出: 4
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums 为 无重复元素 的 升序 排列数组
// -104 <= target <= 104
//
// Related Topics 数组 二分查找
// 👍 1580 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	mid := 0

	for low <= high {
		mid = low + (high - low) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			high = mid - 1
		}

		if nums[mid] < target {
			low = mid + 1
		}
	}

	if nums[mid] > target {
		return mid
	}
	if nums[mid] < target {
		return mid + 1
	}

	return 0
}

func main() {
	nums := []int{1, 3}
	target := 2
	res := searchInsert(nums, target)

	fmt.Println(res)
}

/**
总结：
本题属于二分的稍微变形，在标准二分搜索target的基础上，增加了未查询到target时，target应该插入的位置
解题思路：
	for循环中正常按照二分的逻辑，查找target
	若查找不到,则target应插入mid的左边或右边
	插入左边，相当于插入mid的位置，插入右边，就是插入mid+1的位置

复杂度分析：
	时间复杂度：时间复杂度还是正常二分的时间复杂度O(logN)
	空间复杂度：没有额外申请空间，所以空间复杂度为O(1)
 */
//leetcode submit region end(Prohibit modification and deletion)
