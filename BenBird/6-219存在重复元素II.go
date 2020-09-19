//给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值
// 至多为 k。
//
//
//
// 示例 1:
//
// 输入: nums = [1,2,3,1], k = 3
//输出: true
//
// 示例 2:
//
// 输入: nums = [1,0,1,1], k = 1
//输出: true
//
// 示例 3:
//
// 输入: nums = [1,2,3,1,2,3], k = 2
//输出: false
// Related Topics 数组 哈希表
// 👍 203 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyDuplicate(nums []int, k int) bool {
	temp_map := map[int]int{}

	for key, val := range nums {
		if _, ok := temp_map[val]; ok {
			return true;
		}
		temp_map[val] = 1

		if len(temp_map) > k {
			delete(temp_map, nums[key - k])
		}
	}

	return false
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums := []int{1, 2, 3,1,2,3}
	k := 2
	res := containsNearbyDuplicate(nums, k)

	fmt.Println(res)
}

/**
题解：存在重复元素II

耗时：12ms
内存：6.7M

利用Map模拟集合，集合特点：集合中没有重复元素
初始化Map: temp_map
循环遍历nums
判断每个元素是否存于temp_map,
	若存在：直接返回true
	不存在：将该元素存入Map: temp_map
然后，判断temp_map的长度是否大于k
	若大于：释放temp_map的第一个元素
	不大于：继续遍历nums

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(min(n, k)
 */
