//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
//
//
// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
// Related Topics 数组 哈希表
// 👍 9015 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	new_nums := map[int]int{}
	res := []int{}

	for key, val := range nums {
		if v, ok := new_nums[val]; ok {
			return []int{v, key}
		}

		temp := target - val
		new_nums[temp] = key
	}

	return res
}

//测试用例代码
func main()  {
	nums := [] int {1, 2, 3, 4, 5}
	res := twoSum(nums, 3)
	fmt.Println(res)
}

/** 题解
两数之和：

一遍哈希表：
首先需要开辟额外空间new_nums，用来存储遍历过的元素a的key，和target - a的值b，以b为key1，元素a的下标key2为值，存入数组(此处为切片)中
可表示为 new_nums[target - a] = key2
并在每次遍历时，判断当前元素 是否存在于 new_nums， 如果存在，则返回 改元素在new_nums中的下标和该元素在nums的下标
若不存在，则将改元素需要匹配的另一个数，存入new_nums中，该元素为key，该元素在nums的下标为value，存入new_nums
这样只需要遍历一次，即可找出两个数相加的和为target

复杂度分析：
	时间复杂度：O(n)
	只遍历了包含n个元素的数组一次，并且在数组中查找元素只需O(1)的时间

	空间复杂度：O(n)
	所需额外空间依赖于nums中的元素数量，所以所需额外空间存储元素不超过n个
 */
