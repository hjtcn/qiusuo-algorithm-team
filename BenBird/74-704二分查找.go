//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否
//则返回 -1。
//
//
//示例 1:
//
// 输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
//
//
// 示例 2:
//
// 输入: nums = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不存在 nums 中因此返回 -1
//
//
//
//
// 提示：
//
//
// 你可以假设 nums 中的所有元素是不重复的。
// n 将在 [1, 10000]之间。
// nums 的每个元素都将在 [-9999, 9999]之间。
//
// Related Topics 数组 二分查找
// 👍 840 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	first_key, last_key, mid_key := 0, len(nums) -1 , -1

	for first_key <= last_key {
		//mid_key := (first_key + last_key) / 2
		mid_key = first_key + (last_key - first_key) / 2 //防止内存溢出(first_key + last_key 大于size(int))，但本题可以不用考虑

		if nums[mid_key] == target {
			return mid_key
		}

		if nums[mid_key] > target {
			last_key = mid_key - 1
		}

		if nums[mid_key] < target {
			first_key = mid_key + 1
		}
	}

	return -1
}

func searchTwo(nums []int, target int) int {
	var first_key int = 0
	var last_key int = len(nums) - 1
	var mid_key = -1

	for first_key <= last_key {
		//mid_key := (first_key + last_key) / 2
		mid_key = first_key + (last_key - first_key) / 2 //防止内存溢出(first_key + last_key 大于size(int))，但本题可以不用考虑

		if nums[mid_key] == target {
			return mid_key
		}

		if nums[mid_key] > target {
			last_key = mid_key - 1
		}

		if nums[mid_key] < target {
			first_key = mid_key + 1
		}
	}

	return -1
}

func main() {
	nums := []int{-34, -23, -1, 0, 2, 4, 5, 7, 9, 10, 23, 45, 56}
	target := 6

	res := search(nums, target)
	println(res)
}

//leetcode submit region end(Prohibit modification and deletion)
