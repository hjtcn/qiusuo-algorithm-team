//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]
//
// 示例 2：
//
//
//输入：nums = [-7,-3,2,3,11]
//输出：[4,9,9,49,121]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums 已按 非递减顺序 排序
//
//
//
//
// 进阶：
//
//
// 请你设计时间复杂度为 O(n) 的算法解决本问题
//
// Related Topics 数组 双指针 排序
// 👍 549 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func sortedSquaresTwo(nums []int) []int {
	new_arr := []int{}

	i, j := 0, len(nums) - 1

	for i <= j {
		if i == j {
			new_arr = append(new_arr, nums[i] * nums[i])
			break
		}

		left := nums[i] * nums[i]
		right := nums[j] * nums[j]

		if left > right {
			i++
			new_arr = append(new_arr, left)
		} else {
			j--
			new_arr = append(new_arr, right)
		}
	}

	low, high := 0, len(new_arr) - 1
	for low < high {
		new_arr[low], new_arr[high] = new_arr[high], new_arr[low]
		low++
		high--
	}

	return new_arr
}

func sortedSquaresThree(nums []int) []int {
	nums_len := len(nums)
	i, j, curr_point := 0, nums_len - 1, nums_len - 1

	new_arr := make([]int, nums_len)

	for i <= j {
		low := nums[i] * nums[i]
		high := nums[j] * nums[j]

		if low > high {
			i++
			new_arr[curr_point] = low
		} else {
			j--
			new_arr[curr_point] = high
		}

		curr_point--
	}

	return new_arr
}

func main() {
	nums := []int{-7,-3,2,3,11}
	res := sortedSquaresThree(nums)

	fmt.Println(res)
}

/**
总结
解题思路：
	双指针，因为是非递减序列，还包含负数，不断比较数组两端数字 平方的大小，大的直接塞入新定义的结果集的末尾元素，依次类推

复杂度分析：
	时间复杂度：因为总共对数组遍历了 数组长度 次，也就是n，所以时间复杂度是O(n)
	空间复杂度：除了结果集所需空间，没有开辟新的空间，也就是O(1)
 */
//leetcode submit region end(Prohibit modification and deletion)
