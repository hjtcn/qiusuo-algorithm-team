//给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limi
//t 。

// 如果不存在满足条件的子数组，则返回 0 。
//
//
//
// 示例 1：
//
// 输入：nums = [8,2,4,7], limit = 4
//输出：2
//解释：所有子数组如下：
//[8] 最大绝对差 |8-8| = 0 <= 4.
//[8,2] 最大绝对差 |8-2| = 6 > 4.
//[8,2,4] 最大绝对差 |8-2| = 6 > 4.
//[8,2,4,7] 最大绝对差 |8-2| = 6 > 4.
//[2] 最大绝对差 |2-2| = 0 <= 4.
//[2,4] 最大绝对差 |2-4| = 2 <= 4.
//[2,4,7] 最大绝对差 |2-7| = 5 > 4.
//[4] 最大绝对差 |4-4| = 0 <= 4.
//[4,7] 最大绝对差 |4-7| = 3 <= 4.
//[7] 最大绝对差 |7-7| = 0 <= 4.
//因此，满足题意的最长子数组的长度为 2 。
//
//
// 示例 2：
//
// 输入：nums = [10,1,2,4,7,2], limit = 5
//输出：4
//解释：满足题意的最长子数组是 [2,4,7,2]，其最大绝对差 |2-7| = 5 <= 5 。
//
//
// 示例 3：
//
// 输入：nums = [4,2,2,2,4,4,2,2], limit = 0
//输出：3
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10^5
// 1 <= nums[i] <= 10^9
// 0 <= limit <= 10^9
//
// Related Topics 数组 Sliding Window
// 👍 46 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func longestSubarray(nums []int, limit int) int {
	max_long := 0		//连续子数组的最长长度
	max_num := nums[0]	//可理解为滑动窗口中最大值
	min_num := nums[0]	//滑动窗口中最小值
	curr_long := 0		//滑动窗口的长度
	curr_first_key := 0	//滑动窗口的首个元素的key

	for key, val := range nums {
		//若当前元素大于最大值，则更新最大值
		if val > max_num {
			max_num = val
		}

		//若当前元素小于最小值，则更新最小值
		if val < min_num {
			min_num = val
		}

		if (max_num - min_num) <= limit {	//若最大最小值之差小于等于limit, 则增加当前滑动窗口的长度
			curr_long++
		} else {	//若最大值和最小值之差大于limit，则更新max_long、最大值、最小值、滑动窗口首元素的key
			if curr_long > max_long {
				max_long = curr_long
			}

			curr_first_key++
			min_num, max_num = getMinAndMax(nums[curr_first_key:key+1])
		}

	}

	//若当前滑动窗口的长度大于 max_long，则更新max_long
	if curr_long > max_long {
		max_long = curr_long
	}

	return max_long
}

/**
获取数组中最大值和最小值
*/
func getMinAndMax(nums []int)  (int,int) {
	min := nums[0]
	max := nums[0]

	for _, val := range nums {
		if val > max {
			max = val
		}

		if val < min {
			min = val
		}
	}

	return min, max
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums := []int{10,1,2,4,7,2}
	limit := 5
	res := longestSubarray(nums, limit)

	fmt.Println(res)
}

/**
题解：绝对差不超限制的最长连续子数组

耗时：112ms
内存：9M

维护滑动窗口的最大值和最小值
只要保证最大值和最小值之差不大于limit即可

遍历nums时不断更新最大值和最小值
并判断最大值和最小值之差
若之差 <= limit
	则更新当前当前滑动窗口的长度
若之差 > limit
	则更新最大值、最小值、最长连续子数组的长度

复杂度分析：
	时间复杂度：<O(n^2)
	空间复杂度：O(n)
 */
